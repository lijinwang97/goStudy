package service

// log统计任务
import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"git.inke.cn/tpc/multimedia/server/utils/quick_tool/falcon"
	"github.com/hpcloud/tail"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Data struct {
	BitrateSum int
	BitrateAvg int
	Count      int
	BitrateMax int
	BitrateMin int
}

var (
	line        *tail.Line
	ok          bool
	DataLog     = make(map[string]*Data)
	LockMonitor sync.Mutex
)

func InitData(logPath string) {
	d := DataLog[logPath]
	d.BitrateSum = 0
	d.BitrateAvg = 0
	d.Count = 0
	d.BitrateMax = 0
	d.BitrateMin = math.MaxInt32
}

func InitLog(logPath, liveId string) {
	DataLog[logPath] = &Data{
		BitrateSum: 0,
		BitrateAvg: 0,
		Count:      0,
		BitrateMax: 0,
		BitrateMin: math.MaxInt32,
	}

	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(logPath, config)
	if err != nil {
		log.Error("tail file failed, err:", err)
		return
	}

	//每60s上传falcon一次。
	go func(logPath string) {
		for {
			<-time.After(time.Second * 10)
			host, _ := os.Hostname()
			LockMonitor.Lock()

			dataLog := DataLog[logPath]
			falcon.WriteFalcon("ExeMonitor-SumBitrate", "hostname:"+host+liveId, dataLog.BitrateSum, 60) // 加atomic
			falcon.WriteFalcon("ExeMonitor-AvgBitrate", "hostname:"+host+liveId, dataLog.BitrateAvg, 60)
			falcon.WriteFalcon("ExeMonitor-MaxBitrate", "hostname:"+host+liveId, dataLog.BitrateMax, 60)
			falcon.WriteFalcon("ExeMonitor-MinBitrate", "hostname:"+host+liveId, dataLog.BitrateMin, 60)
			log.Info("上传falcon的信息：sum:", dataLog.BitrateSum, " avg:", dataLog.BitrateAvg, " max:", dataLog.BitrateMax, " min:", dataLog.BitrateMin, " count:", dataLog.Count)
			InitData(logPath)
			log.Info("重新初始化后的  sum:", dataLog.BitrateSum, "    avg:", dataLog.BitrateAvg, "    max:", dataLog.BitrateMax, "   min:", dataLog.BitrateMin, "   count:", dataLog.Count)
			LockMonitor.Unlock()

		}
	}(logPath)

	for {
		line, ok = <-tails.Lines
		if !ok {
			log.Error("tail file close reopen, filename:", tails.Filename)
			return
		}
		/*
			log
			2020/07/24 17:33:19.430828  INFO [RTMPPUSH1] < R Acknowledgement. ignore. sequence number=1287192617. - client_session.go:210    split(" ")切分后的长度后13
			2020/07/24 17:33:19.658183 DEBUG bitrate=1627.272kbit/s - pushrtmp.go:69		split(" ")切分后的长度后6
		*/
		if !strings.Contains(line.Text, "bitrate") {
			continue
		}
		// 选出带有bitrate日志，获取其中bitrate
		bitrate := getBitrateFromLog(line.Text)
		if err == nil {
			LockMonitor.Lock()
			datalog := DataLog[logPath]
			datalog.BitrateMax = getMax(datalog.BitrateMax, bitrate)
			datalog.BitrateMin = getMin(datalog.BitrateMin, bitrate)
			datalog.BitrateSum = sumData(datalog.BitrateSum, bitrate)
			datalog.Count++
			datalog.BitrateAvg = datalog.BitrateSum / datalog.Count
			//log.Info("普通的  sum:", datalog.BitrateSum, "    avg:", datalog.BitrateAvg, "    max:", datalog.BitrateMax, "   min:", datalog.BitrateMin, "   count:", datalog.Count)
			LockMonitor.Unlock()
		}
	}
}

//获取log中的bitrate数值
func getBitrateFromLog(logString string) int {
	reg := regexp.MustCompile(`=[0-9]+.`)

	bitReg := reg.FindAllStringSubmatch(logString, -1)[0][0]
	bitInt := bitReg[1 : len(bitReg)-1]

	bitResult, err := strconv.Atoi(bitInt)
	if err == nil {
		return bitResult
	}
	return 0
}

func getMin(min int, bitrate int) int {
	if min <= bitrate {
		return min
	}
	return bitrate
}

func getMax(max int, bitrate int) int {
	if max >= bitrate {
		return max
	}
	return bitrate
}

func sumData(sum int, bitrate int) int {
	return sum + bitrate
}
