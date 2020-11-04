package service

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/multimedia/server/utils/media.stream.push_continuous/conf"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var (
	cmd      *exec.Cmd
	timeInfo TimeInfo
)

type TimeInfo struct {
	timeNow        int64
	afterMonthTime int64
	monthTimeStamp int64
}

func InitAndStartExe(cfg *conf.Config) {
	StartUserExe(cfg)
}

func KillAll(threadNmae string) error {
	cmd := exec.Command("killall", threadNmae)
	return cmd.Run()
}

//获取线程名，如：./bin/pushrtmp_macos ，获取 pushrtmp_macos
func getThreadName(pushrtmp string) string {
	splitThreadNmae := strings.Split(pushrtmp, "/")
	return splitThreadNmae[len(splitThreadNmae)-1]
}

// 启动配置里的用户的exe
func StartUserExe(cfg *conf.Config) {
	for _, userList := range cfg.NodeInfos.UserList {
		// Init
		threadNmae := getThreadName(userList.Pushrtmp)
		// 如果此线程有残留进程，则杀死
		err := KillAll(threadNmae)
		if err != nil {
			log.Info(err)
		}
		logAddress := "./logs/info-" + userList.LiveId + ".log"
		go StartExe(userList.Pushrtmp, userList.Video, userList.UserAddress, userList.LiveId, logAddress,cfg.MonitorConfig.RtmpBegin,cfg.MonitorConfig.RtmpLast)
		go InitLog(logAddress, userList.LiveId)
	}
}

// 启动exe
func StartExe(pushrtmp, video, userAddress, liveId, logAddress,rtmpBegin,rtmpLast string) error {
	liveAddress := rtmpBegin + userAddress + rtmpLast + liveId
	pushInfo := fmt.Sprintf("-i %s -o %s -r -l %s", video, liveAddress, logAddress)
	pushInfoSplit := strings.Split(pushInfo, " ")
	cmd := exec.Command(pushrtmp, pushInfoSplit...)
	if err := cmd.Start(); err != nil {
		log.Error("exe start err:", err)
		return err
	}
	log.Info("start exe is success! order: ", pushrtmp, " ", pushInfo)
	pid := cmd.Process.Pid
	go func() { // 一个月重启一次
		<-time.After(time.Hour * time.Duration(24*30))
		restartExe(pid)
	}()

	if err := cmd.Wait(); err != nil {
		log.Error("exe wait failed: ", err)
		//中断退出的话，会有残留进程，需要关闭
		if strings.Contains(err.Error(), "interrupt") {
			syscall.Kill(pid, syscall.SIGKILL)
			return err
		}
		StartExe(pushrtmp, video, userAddress, liveId, logAddress,rtmpBegin,rtmpLast)
		return err
	}
	return nil
}

//  重启
func restartExe(pid int) {
	if err := syscall.Kill(pid, 0); err == nil {
		syscall.Kill(pid, syscall.SIGKILL)
	}
}
