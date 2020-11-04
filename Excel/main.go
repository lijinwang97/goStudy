package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Luxurioust/excelize"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//zego_microphoneNum(zego连麦主播数),cdn_play（cdn点播）,cdn_playBusiness（cdn点播业务量）
var clsSlice, livegoHuiyuan, zego_audioSlice, zego_cdnSlice, livego_centerSlice, cdn_collectSlice, livego_edgeSlice, cls_livego_cdnSlice, livego_huiyuanSlice, livego_TotalSlice,zego_microphoneNum,cdn_playNum,cdn_playBusiness []float64
var ali_cdnSlice, ws_cdnSlice, qq_cdnSlice, csy_cdnSlice, cls_livego_edgeSlice, cls_livego_centerSlice, cdnTotalSlice []float64
var origAli, origTx, origTotal []float64

//var add_result = add{}

func main() {

	falcon_postFuncInit()

	/*strMachine := General_Html_Machine("./machine.xlsx")
	totalWB := General_Html_BandWidth()
	datailWB := Details_Html_BandWidth()
	s := XHML_Body(strMachine, totalWB, datailWB)
	mailTo := []string{
		"guoyu@inke.cn",
		"jiangqingshan@inke.cn",
		//"wangsong@inke.cn",
		"lijinwang@inke.cn",
	}
	//邮件主题为"Hello"
	subject := "多媒体带宽统计"
	// 邮件正文  //
	body := "<html><body><table border='1' style='font-size:22px' cellspacing='0' cellpadding='0' >" + s + "</table></body></html>"
	err := Send_Mail(mailTo, subject, body)
	if err != nil {
		fmt.Printf("send gomail fail error:%v", err)
		return
	}*/

}

func falcon_postFuncInit() {
	for i := 1; i < 9; i++ { //这里1-9

		start_time, end_time := initData(i)
		getToken()
		//cls 总量
		cls := post_body{cls_host, cls_counters}
		origin_data := getdata(start_time, end_time, cls)
		all := deal_data(origin_data) / 1024 / 1024 / 1024
		fmt.Println(all)
		clsSlice = append(clsSlice, all)

		/*//ali源站
		orgAli := post_body{origAli_host, origAli_counters}
		origin_data = getdata(start_time, end_time, orgAli)
		oali := deal_data(origin_data) / 1024 / 1024 / 1024
		origAli = append(origAli, oali)

		//tx源站
		orgTx := post_body{origTx_host, origTX_counters}
		origin_data = getdata(start_time, end_time, orgTx)
		otx := deal_data(origin_data) / 1024 / 1024 / 1024
		origTx = append(origTx, otx)

		origTotal = append(origTotal, oali+otx)

		//livego 回源
		livego_host := post_body{livego_host, livego_counters}
		origin_data = getdata(start_time, end_time, livego_host)
		livegoc := deal_data(origin_data) * 3 / 100
		livegoHuiyuan = append(livegoHuiyuan, livegoc)
		fmt.Println(livegoHuiyuan)
		//ali cdn
		ali_cdn := post_body{ali_cdn_host, ali_cdn_counters} ///1024/1024/1024
		origin_data = getdata(start_time, end_time, ali_cdn)
		cdn1 := deal_data(origin_data)
		ali_cdnSlice = append(ali_cdnSlice, cdn1/1024/1024/1024)
		//ws cdn
		ws_cdn := post_body{ws_collect_host, ws_cdn_counters}
		origin_data = getdata(start_time, end_time, ws_cdn)
		cdn4 := deal_data(origin_data)
		ws_cdnSlice = append(ws_cdnSlice, cdn4/1024/1024/1024)
		//tx cdn
		qq_cdn := post_body{qq_cdn_host, qq_cdn_counters}
		origin_data = getdata(start_time, end_time, qq_cdn)
		cdn2 := deal_data(origin_data)
		qq_cdnSlice = append(qq_cdnSlice, cdn2/1024/1024/1024)
		//csy cdn
		csy_cdn := post_body{csy_cdn_host, csy_cdn_counters} ///1024/1024/1024
		origin_data = getdata(start_time, end_time, csy_cdn)
		cdn3 := deal_data(origin_data)
		csy_cdnSlice = append(csy_cdnSlice, cdn3/1024/1024/1024)

		cdnTotalSlice = append(cdnTotalSlice, (cdn1+cdn2+cdn2+cdn4)/1024/1024/1024)

		//livego edge
		cls_livego_edge := post_body{cls_livego_edge_host, cls_livego_edge_counters}
		origin_data = getdata(start_time, end_time, cls_livego_edge)
		livego1 := deal_data(origin_data) / 100
		cls_livego_edgeSlice = append(cls_livego_edgeSlice, livego1)

		//livego 中心
		cls_livego_center := post_body{cls_livego_centre_host, cls_livego_centre_counters}
		origin_data = getdata(start_time, end_time, cls_livego_center)
		livego2 := deal_data(origin_data) * 3 / 100
		cls_livego_centerSlice = append(cls_livego_centerSlice, livego1)

		livego_TotalSlice = append(livego_TotalSlice, livego2+livego1)

		//zego 连麦
		zego_audio := post_body{zego_audio_host, zego_audio_counters} ///1024/1024/1024
		origin_data = getdata(start_time, end_time, zego_audio)
		all = deal_data(origin_data)
		zego_audioSlice = append(zego_audioSlice, all/1024/1024/1024)

		//
		zego_cdn := post_body{zego_cdn_host, zego_cdn_counters} ///1024/1024/1024
		origin_data = getdata(start_time, end_time, zego_cdn)
		all = deal_data(origin_data)
		zego_cdnSlice = append(zego_cdnSlice, all/1024/1024/1024)*/

		// zego连麦主播数   todo 明天给小瑜姐检查
		zego_microphone := post_body{zego_microphone_host, zego_microphone_counters}
		origin_data = getdata(start_time, end_time, zego_microphone)
		all = deal_data(origin_data)
		zego_microphoneNum = append(zego_microphoneNum, all)

		// cdn点播
		cdn_play := post_body{cdn_play_host, cdn_play_counters}
		origin_data = getdata(start_time, end_time, cdn_play)
		all = deal_data(origin_data)
		cdn_playNum = append(cdn_playNum, all)


		// cdn点播业务量
	}
}

//带宽总量概括
func General_Html_BandWidth() string {
	var dayy []string

	for i := 1; i < 9; i++ { //这里1-9
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		//day1 := time.Now().AddDate(0, 0, -i).Format("20060102")
		if i < 8 { //这里8
			dayy = append(dayy, day)
		}
	}

	general_html := fmt.Sprintf("<tr bgcolor='Cyan'><td colspan='8' align='center' valian='middle'>%s</td></tr>", "带宽总量概览(Gbps)") +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "项目", RowPrintHtml1(dayy)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "cls项目", RowPrintHtml(clsSlice)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "源站", RowPrintHtml(origTotal)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "cdn带宽", RowPrintHtml(cdnTotalSlice)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "livego带宽", RowPrintHtml(livego_TotalSlice)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "zego连麦带宽", RowPrintHtml(zego_audioSlice)) +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "zego cdn带宽", RowPrintHtml(zego_cdnSlice))+
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "zego 连麦主播数", RowPrintHtml(zego_microphoneNum))+   //添加三个内容
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "cdn点播", RowPrintHtml(cdn_playNum))+
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td>%s</tr>", "cdn 点播业务量(点播业务量大于等于静态域名数量)", RowPrintHtml(cdn_playBusiness))
	return general_html

}

func Details_Html_BandWidth() string {
	var dayy []string
	var clsTotalDisp, clsOut, origsidecdn, origsideali, origsidetx, origsideacsy, origsidews []float64

	for i := 1; i < 9; i++ { //这里1-9
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		//day1 := time.Now().AddDate(0, 0, -i).Format("20060102")
		day1 := time.Now().AddDate(0, 0, -i).Format("20060102")

		//这里是 hive
		//cls 云联网
		sum := deal_data_2rows_total("clsdisppush" + day1)
		sum1 := deal_data_2rows_total("clsdispplay" + day1)
		clsTotalDisp = append(clsTotalDisp, sum1+sum)

		//cls 外网带宽
		sum = deal_data_2rows_total("clsout" + day1) // 后续需要改下
		clsOut = append(clsOut, sum)

		//回源ali
		cdn1 := deal_data_2rows_total("originMinuteWBali" + day1)
		origsideali = append(origsideali, cdn1)
		//回源 tx
		cdn2 := deal_data_2rows_total("originMinuteWBtx" + day1)
		origsidews = append(origsidews, cdn2)
		//回源创世云
		cdn3 := deal_data_2rows_total("originMinuteWBcsy" + day1)
		origsidetx = append(origsidetx, cdn3)
		//回源网宿
		cdn4 := deal_data_2rows_total("originMinuteWBws" + day1)
		origsideacsy = append(origsideacsy, cdn4)
		origsidecdn = append(origsidecdn, cdn1+cdn2+cdn3+cdn4)

		if i < 8 { //这里8
			dayy = append(dayy, day)
		}
	}

	details_html := fmt.Sprintf("<tr bgcolor='Cyan'><td colspan='10' align='center' valian='middle'>%s</td></tr>", "带宽总量详情") +
		fmt.Sprintf("<tr bgcolor='AliceBlue'><td align='center' valian='middle'>%s</td><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "项目", "带宽(Gbps)", RowPrintHtml1(dayy)) +
		fmt.Sprintf("<tr bgcolor='Azure'><td rowspan='2' align='center' valian='middle'>%s</td><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "cls", "云联网带宽", RowPrintHtml(clsTotalDisp)) +
		fmt.Sprintf("<tr bgcolor='Azure'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "外网带宽(Mbps)", RowPrintHtml(clsOut)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td rowspan='9' align='center' valian='middle'>%s</td><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "源站带宽", "外网带宽", RowPrintHtml(origTotal)) +
		//fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "外网带宽(自己)",RowPrintHtml(originTotal))+
		fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "阿里源站带宽", RowPrintHtml(origAli)) +
		//fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "阿里源站带宽(自己)",RowPrintHtml(originAl))+
		fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "腾讯源站带宽", RowPrintHtml(origTx)) +
		//fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "腾讯源站带宽(自己)",RowPrintHtml(originTx))+
		fmt.Sprintf("<tr bgcolor='Bisque'><td rowspan='5' align='center' valian='middle'>%s</td><td align='center' valian='middle'>%s</td>%s</tr>", "cdn回源带宽", "总量", RowPrintHtml(origsidecdn)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td align='center' valian='middle'>%s</td>%s</tr>", "阿里", RowPrintHtml(origsideali)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td align='center' valian='middle'>%s</td>%s</tr>", "网宿", RowPrintHtml(origsidews)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td align='center' valian='middle'>%s</td>%s</tr>", "腾讯", RowPrintHtml(origsidetx)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td align='center' valian='middle'>%s</td>%s</tr>", "创世云", RowPrintHtml(origsideacsy)) +
		fmt.Sprintf("<tr bgcolor='Bisque'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "livego回源带宽", RowPrintHtml(livegoHuiyuan)) +
		fmt.Sprintf("<tr bgcolor='LightCyan'><td rowspan='5' align='center' valian='middle'>%s</td><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "cdn带宽", "总量", RowPrintHtml(cdnTotalSlice)) +
		//fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "总量(自己)",RowPrintHtml(cdnTotal))+
		fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "阿里", RowPrintHtml(ali_cdnSlice)) +
		//fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "阿里(自己)",RowPrintHtml(cdnAl))+
		fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "创世云", RowPrintHtml(csy_cdnSlice)) +
		//fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "创世云(自己)",RowPrintHtml(cdnCsy))+
		fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "腾讯", RowPrintHtml(qq_cdnSlice)) +
		//fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "腾讯(自己)",RowPrintHtml(cdnTx))+
		fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "网宿", RowPrintHtml(ws_cdnSlice)) +
		//fmt.Sprintf("<tr bgcolor='LightCyan'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "网宿(自己)",RowPrintHtml(cdnWs))+
		fmt.Sprintf("<tr bgcolor='Linen'><td align='center' valian='middle'>%s</td><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "livego带宽", "总量", RowPrintHtml(livego_TotalSlice))
	//fmt.Sprintf("<tr bgcolor='Linen'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "livego代理带宽(业务)",RowPrintHtml(cls_livego_centerSlice))
	//fmt.Sprintf("<tr bgcolor='Linen'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "总量(自己)",RowPrintHtml(livego_TotalSlice))+
	//fmt.Sprintf("<tr bgcolor='Linen'><td colspan='2' align='center' valian='middle'>%s</td>%s</tr>", "livego代理带宽(自己)",RowPrintHtml(livego_Slice))
	fmt.Println("1", cls_livego_cdnSlice)

	return details_html
}

//机器成本
func General_Html_Machine(excelPath string) string {
	results := readExcel(excelPath)

	var general_html string
	general_html = fmt.Sprintf("<tr bgcolor='Cyan'><td colspan='3' align='center' valian='middle'>%s</td></tr>", "机器成本") +
		fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td><td  align='center' valian='middle'>%s</td><td align='center' valian='middle'>%s</td></tr>", "项目", "机器数", "合计(元)")
	for _, row := range results {
		general_html = general_html + fmt.Sprintf("<tr><td align='center' valian='middle'>%s</td><td  align='center' valian='middle'>%s</td><td align='center' valian='middle'>%s</td></tr>", row[0], row[1], row[2])

	}
	return general_html

}

func XHML_Body(general_html string, totalWB string, detailWB string) string {
	str := "<html><body><br/><br/><br/><table border='1' style='font-size:13px' cellspacing='0' cellpadding='0' width='1100' height='350'>" + totalWB + "</table>" +
		"<br/><br/><br/><table border='1' style='font-size:13px' cellspacing='0' cellpadding='0' width='1300' height='1200'>" + detailWB + "</table>" +
		"<br/><br/><br/><table border='1' style='font-size:13px' cellspacing='0' cellpadding='0' width='800' height='800'>" + general_html + "</table></body></html>"

	return str
}

func deal_data_2rows_total(file string) float64 {

	data_value := total{}

	path := file + ".txt"

	//println(path)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		println(err)
	}
	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		data := strings.Split(line, "\t")

		data_value.value, _ = strconv.ParseFloat(data[1], 64)
		break
	}
	return data_value.value
}

func readExcel(excelPath string) [][]string {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))
	return rows
}

func Send_Mail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "15140504381@163.com",
		"pass": "yumei,7231",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "TPC Media Server "+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                                      //发送给多个用户
	m.SetHeader("Subject", subject)                                   //设置邮件主题
	m.SetBody("text/html", body)                                      //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}

func RowPrintHtml1(list []string) string {
	return fmt.Sprintf("<td align='center' valian='middle'>%s</td>", strings.Join(list, "</td><td align='center' valian='middle'>"))
}
func RowPrintHtml(list []float64) string {
	var str string
	var stringlen string
	flag := "black"
	if len(list) == 0 {
		return stringlen
	}
	in1 := list[0]
	for i := 1; i <= len(list)-1; i++ {
		in2 := list[i]
		if in1-in2 < 0 {
			flag = "red"
		}
		if (in1 - in2) == 0.0 {

			str = fmt.Sprintf("--")
		} else {
			str = fmt.Sprintf("%.2f", in1-in2)
		}
		//strr := strconv.FormatFloat(list[i-1],'f',-1,64)
		strr := fmt.Sprintf("%.2f", list[i-1])
		stringlen += fmt.Sprintf("<td align='center' valian='middle'><font color=%s>%s</font></td>", flag, strr+"("+str+")")
		in1 = in2
		flag = "black"
	}

	//stringlen+=fmt.Sprintf("<td><font color=%s>%s</font></td>",flag,list[len(list)-1])
	return stringlen
	//fmt.Sprintf("<td>%s</td>", strings.Join(list, "</td><td>"))
}

type post_body struct {
	host     []string
	counters []string
}

//var slice = []Slices{}

type data_return []struct {
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
	Dstype   string `json:"dstype"`
	Step     int    `json:"step"`
	Values   []struct {
		Timestamp int     `json:"timestamp"`
		Value     float64 `json:"value"`
	} `json:"Values"`
	sum float64
}

type Slices struct {
	values   []value
	endpoint string
}

type tokens struct {
	Sig   string `json:"sig"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

type headers struct {
	Name string `json:"name"`
	Sig  string `json:"sig"`
}

type payloads struct {
	Step      int      `json:"step"`
	StartTime int      `json:"start_time"`
	HostNames []string `json:"hostnames"`
	EndTime   int      `json:"end_time"`
	Counters  []string `json:"counters"`
	ConsolFun string   `json:"consol_fun"`
}

var token = tokens{}
var header = headers{}

func getToken() {
	reqURL := "http://10.100.8.47:8080/api/v1/user/login"
	res, _ := http.Post(reqURL,
		"application/x-www-form-urlencoded",
		strings.NewReader("password=api123!&name=api"))
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &token)
	if err != nil {
		println(err)
	}
	header.Name = token.Name
	header.Sig = token.Sig
	//println(token.Sig)
}

func initData(i int) (string, string) {
	//time
	now := time.Now()
	d, _ := time.ParseDuration(fmt.Sprintf("%d%s", -24*i, "h"))
	yesterday := now.Add(d)
	year := strconv.Itoa(yesterday.Year())
	month := strconv.Itoa(int(yesterday.Month()))
	if int(yesterday.Month()) < 10 {
		month = "0" + month
	}
	day := strconv.Itoa(yesterday.Day())
	if int(yesterday.Day()) < 10 {
		day = "0" + day
	}
	start_time := year + "-" + month + "-" + day + " 00:00:00"
	end_time := year + "-" + month + "-" + day + " 23:59:59"

	return start_time, end_time
}

func getdata(start_time string, end_time string, bodys post_body) data_return {
	jsonStu, err := json.Marshal(header)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	start, _ := time.ParseInLocation("2006-01-02 15:04:05", start_time, loc)
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", end_time, loc)
	startstamp := int(start.Unix())
	endstamp := int(end.Unix())
	//println(startstamp)

	payload := payloads{60 * 5, startstamp, bodys.host, endstamp, bodys.counters, "AVERAGE"}
	jsonStr, _ := json.Marshal(payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://10.100.8.47:8080/api/v1/influx/history", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apitoken", string(jsonStu))
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var data = data_return{}
	_ = json.Unmarshal(body, &data)
	//fmt.Println(string(body))
	return data
}

type total struct {
	host  string
	value float64
}

var origAli_host = []string{
	"tx3-pe-cdn-monitor01.bj",
}
var origAli_counters = []string{
	"ali-liveorig-cbwp/cbwpid=cbwp-2zer4uvybg9gfu2p22d0a",
	"ali-liveorig-cbwp/cbwpid=cbwp-2zewf78xqchd5ocbdfebs",
}

var origTx_host = []string{
	"tx3-pe-cdn-monitor01.bj",
}
var origTX_counters = []string{
	" tx-slb-orig/slb=111.13.240.196",
	"tx-slb-orig/slb=111.13.240.209",
	"tx-slb-orig/slb=123.123.190.164",
	"tx-slb-orig/slb=123.123.190.234",
	"tx-slb-orig/slb=140.143.213.178",
	"tx-slb-orig/slb=49.7.70.142",
	"tx-slb-orig/slb=49.7.70.171",
}

var cls_host = []string{
	"ali-e-media-testqa03.bj",
	"tx-k8s-media-rtc-cls-config02.bj",
	"tx1-media-cls01.de01",
	"tx1-media-cls01.hk01",
	"tx1-media-cls01.jp01",
	"tx1-media-cls01.us01",
	"tx1-media-cls02.de01",
	"tx1-media-cls02.hk01",
	"tx1-media-cls02.jp01",
	"tx1-media-cls02.us01",
	"tx3-media-cls-01.gz",
	"tx3-media-cls-01.sh",
	"tx3-media-cls-02.gz",
	"tx3-media-cls-02.sh",
	"tx3-media-cls-03.gz",
	"tx3-media-cls-03.sh",
	"tx3-media-cls-04.gz",
	"tx3-media-cls-04.sh",
	"tx3-media-cls-05.gz",
	"tx3-media-cls-06.gz",
	"tx3-media-cls-07.gz",
	"tx3-media-cls-08.gz",
	"tx3-media-cls-09.gz",
	"tx3-media-cls-10.gz",
	"tx3-media-cls-config01.bj",
	"tx3-media-cls-disp01.bj",
	"tx3-media-cls-disp02.bj",
	"tx3-media-cls01.bj",
	"tx3-media-cls02.bj",
	"tx3-media-cls03.bj",
	"tx3-media-cls04.bj",
	"tx3-media-cls05.bj",
	"tx3-media-cls28.bj",
	"tx3-media-cls29.bj",
	"tx3-media-cls30.bj",
	"tx3-media-cls31.bj",
	"tx3-media-cls32.bj",
	"tx3-media-cls33.bj",
	"tx3-media-cls34.bj",
	"tx3-media-cls35.bj",
	"tx3-media-test01.bj",
	"tx4-media-cls-11.sh",
	"tx4-media-cls-12.sh",
	"tx4-media-cls-13.sh",
	"tx4-media-cls-14.sh",
	"tx4-media-cls-15.sh",
	"tx4-media-cls-16.sh",
	"tx4-media-cls-disp03.bj",
	"tx4-media-cls-disp04.bj",
	"tx4-media-cls18.bj",
	"tx4-media-cls19.bj",
	"tx4-media-cls20.bj",
	"tx4-media-cls21.bj",
	"tx4-media-cls22.bj",
	"tx4-media-rtmp2cls01.bj",
	"tx5-media-cls-disp05.bj",
	"tx5-media-cls-disp06.bj",
	"tx5-media-cls23.bj",
	"tx5-media-cls24.bj",
	"tx5-media-cls25.bj",
	"tx5-media-cls26.bj",
	"tx5-media-cls27.bj",
	"tx5-media-rtmp2cls02.bj",
}
var cls_counters = []string{
	//"event.code.count/code=0,event=ClsClientPublisherNumber.End,project=media.cls.agent.tx",
	"WanOuttraffic",
}

var livego_host = []string{
	"tx3-media-cls-disp01.bj",
}
var livego_counters = []string{
	"livegoquic-bw_send/hostname=代理腾讯01三线49.7.70.239",
	"livegoquic-bw_send/hostname=代理腾讯02三线49.7.70.68",
}

var ali_cdn_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var ali_cdn_counters = []string{
	"ali-cdn-bps/domain=alsource.rtc.gmugmu.com",
	"ali-cdn-bps/domain=alsource.rtc.inke.cn",
	"ali-cdn-bps/domain=alsource.test.inke.cn",
	"ali-cdn-bps/domain=alwebhls.inke.cn",
	"ali-cdn-bps/domain=audio-pull-al.inke.cn",
	"ali-cdn-bps/domain=audit-pull-al.inke.cn",
	"ali-cdn-bps/domain=extaudio-pull-al.inke.cn",
	"ali-cdn-bps/domain=extaudio-pull-ws.inke.cn",
	"ali-cdn-bps/domain=extlive-pull-al.inke.cn",
	"ali-cdn-bps/domain=foreign-audio-pull-al.inke.cn",
	"ali-cdn-bps/domain=foreign-live-pull-al.inke.cn",
	"ali-cdn-bps/domain=mini-pull-al.inke.cn",
	"ali-cdn-bps/domain=msgpush-live-pull-al.inke.cn",
}

var ws_collect_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var ws_cdn_counters = []string{
	"ws-cdn-bps/domain=act.inke.cn",
	"ws-cdn-bps/domain=activity-web.cdn.eomapp.cn",
	"ws-cdn-bps/domain=al.orig.inke.cn",
	"ws-cdn-bps/domain=aladdin-download.gmugmu.com",
	"ws-cdn-bps/domain=ali-bj-bgp.push.inke.cn",
	"ws-cdn-bps/domain=app.caratsvip.com",
	"ws-cdn-bps/domain=app.inke.cn",
	"ws-cdn-bps/domain=app.inke.com",
	"ws-cdn-bps/domain=app.yingtaorelian.com",
	"ws-cdn-bps/domain=audio-pull-ws.inke.cn",
	"ws-cdn-bps/domain=audit-pull-ws.inke.cn",
	"ws-cdn-bps/domain=avatar-dev.cdn.gmugmu.com",
	"ws-cdn-bps/domain=avatar.cdn.gmugmu.com",
	"ws-cdn-bps/domain=boss-images-gray.gmugmu.com",
	"ws-cdn-bps/domain=bujiu.img.ikstatic.cn",
	"ws-cdn-bps/domain=buylivehk.8686c.com",
	"ws-cdn-bps/domain=complain-dev.cdn.gmugmu.com",
	"ws-cdn-bps/domain=complain.cdn.gmugmu.com",
	"ws-cdn-bps/domain=cover-dev1.cdn.gmugmu.com",
	"ws-cdn-bps/domain=dl.meelive.cn",
	"ws-cdn-bps/domain=download.gmugmu.com",
	"ws-cdn-bps/domain=download.xiangshengclub.com",
	"ws-cdn-bps/domain=extaudio-pull-ws.inke.cn",
	"ws-cdn-bps/domain=extlive-pull-ws.inke.cn",
	"ws-cdn-bps/domain=foreign-audio-pull-ws.inke.cn",
	"ws-cdn-bps/domain=foreign-live-pull-ws.inke.cn",
	"ws-cdn-bps/domain=fx.img.imilive.cn",
	"ws-cdn-bps/domain=gaia.imilive.cn",
	"ws-cdn-bps/domain=gmu-gos.gmugmu.com",
	"ws-cdn-bps/domain=gmu-live-video.gmugmu.com",
	"ws-cdn-bps/domain=gmutest.gmugmu.com",
	"ws-cdn-bps/domain=h5.anyueclub.com",
	"ws-cdn-bps/domain=h5.inke.cn",
	"ws-cdn-bps/domain=ik-feed.inke.cn",
	"ws-cdn-bps/domain=im-images-gray.gmugmu.com",
	"ws-cdn-bps/domain=im-media-img.gmugmu.com",
	"ws-cdn-bps/domain=im-media-test.gmugmu.com",
	"ws-cdn-bps/domain=im-media-voc.gmugmu.com",
	"ws-cdn-bps/domain=image.scale.inke.cn",
	"ws-cdn-bps/domain=image.scale.inke.com",
	"ws-cdn-bps/domain=image.yaamari.com",
	"ws-cdn-bps/domain=images-gray.gmugmu.com",
	"ws-cdn-bps/domain=imagescale.9zhenge.com",
	"ws-cdn-bps/domain=imagescale.anyueclub.com",
	"ws-cdn-bps/domain=imagescale.boluohuyu.cn",
	"ws-cdn-bps/domain=imagescale.caratsvip.com",
	"ws-cdn-bps/domain=imagescale.fengyuhn.cn",
	"ws-cdn-bps/domain=imagescale.gmugmu.com",
	"ws-cdn-bps/domain=imagescale.hnyapu.cn",
	"ws-cdn-bps/domain=imagescale.hnyplive.cn",
	"ws-cdn-bps/domain=imagescale.hnypvoice.cn",
	"ws-cdn-bps/domain=imagescale.ichaoren.com",
	"ws-cdn-bps/domain=imagescale.imilive.cn",
	"ws-cdn-bps/domain=imagescale.inke.cn",
	"ws-cdn-bps/domain=imagescale.meetcplive.com",
	"ws-cdn-bps/domain=imagescale.starstarlive.com",
	"ws-cdn-bps/domain=imagescale.yaamari.com",
	"ws-cdn-bps/domain=imagescaledyik.zhishichaoren.cn",
	"ws-cdn-bps/domain=imagescaleik.meetstarlive.com",
	"ws-cdn-bps/domain=imagescalejyik.gzhuda.com",
	"ws-cdn-bps/domain=imagescaleky.hnyapu.cn",
	"ws-cdn-bps/domain=imagescalekyik.gzsnxxkj8.com",
	"ws-cdn-bps/domain=imagescalekyik.hnyapu.cn",
	"ws-cdn-bps/domain=imagescalekyik.xiangshengclub.com",
	"ws-cdn-bps/domain=img-buylive.ikbase.cn",
	"ws-cdn-bps/domain=img.1yimi.cn",
	"ws-cdn-bps/domain=img.9zhenge.com",
	"ws-cdn-bps/domain=img.anyueclub.com",
	"ws-cdn-bps/domain=img.boluohuyu.cn",
	"ws-cdn-bps/domain=img.caratsvip.com",
	"ws-cdn-bps/domain=img.fengyuhn.cn",
	"ws-cdn-bps/domain=img.gzhuda.com",
	"ws-cdn-bps/domain=img.gzlywlkj8.com",
	"ws-cdn-bps/domain=img.hnyapu.cn",
	"ws-cdn-bps/domain=img.hnyplive.cn",
	"ws-cdn-bps/domain=img.hnypvoice.cn",
	"ws-cdn-bps/domain=img.inke.cn",
	"ws-cdn-bps/domain=img.meelive.cn",
	"ws-cdn-bps/domain=img.meetcplive.com",
	"ws-cdn-bps/domain=img.starstarlive.com",
	"ws-cdn-bps/domain=img.thlrs.com",
	"ws-cdn-bps/domain=img.xiangshengclub.com",
	"ws-cdn-bps/domain=img.yingtaorelian.com",
	"ws-cdn-bps/domain=img.zhishichaoren.cn",
	"ws-cdn-bps/domain=img2.inke.cn",
	"ws-cdn-bps/domain=img2ik.meetcplive.com",
	"ws-cdn-bps/domain=img2ik.meetstarlive.com",
	"ws-cdn-bps/domain=img2jyik.gzhuda.com",
	"ws-cdn-bps/domain=img2kyik.xiangshengclub.com",
	"ws-cdn-bps/domain=imgky.hnyapu.cn",
	"ws-cdn-bps/domain=imgtx.ikstatic.cn",
	"ws-cdn-bps/domain=imgyimi.bjdyzg.com",
	"ws-cdn-bps/domain=istream.8686c.com",
	"ws-cdn-bps/domain=istream.inke.cn",
	"ws-cdn-bps/domain=istream2.8686c.com",
	"ws-cdn-bps/domain=istream2.inke.cn",
	"ws-cdn-bps/domain=istream3.8686c.com",
	"ws-cdn-bps/domain=istream4.8686c.com",
	"ws-cdn-bps/domain=istream5.8686c.com",
	"ws-cdn-bps/domain=istream6.8686c.com",
	"ws-cdn-bps/domain=istream6.inke.cn",
	"ws-cdn-bps/domain=istream7.8686c.com",
	"ws-cdn-bps/domain=istream99.8686c.cn",
	"ws-cdn-bps/domain=jsrecord.inke.cn",
	"ws-cdn-bps/domain=live-hls-ws.buylivehk.com",
	"ws-cdn-bps/domain=live-pull-ws.buylivehk.com",
	"ws-cdn-bps/domain=live-push-ws.buylivehk.com",
	"ws-cdn-bps/domain=live-rtcpull-ws.buylivehk.com",
	"ws-cdn-bps/domain=livews.imilive.cn",
	"ws-cdn-bps/domain=lyrtx.ikstatic.cn",
	"ws-cdn-bps/domain=m4a.9zhenge.com",
	"ws-cdn-bps/domain=m4a.boluohuyu.cn",
	"ws-cdn-bps/domain=m4a.fengyuhn.cn",
	"ws-cdn-bps/domain=m4a.hnyapu.cn",
	"ws-cdn-bps/domain=m4a.hnyplive.cn",
	"ws-cdn-bps/domain=m4a.hnypvoice.cn",
	"ws-cdn-bps/domain=m4a.xiangshengclub.com",
	"ws-cdn-bps/domain=m4a.yingtaorelian.com",
	"ws-cdn-bps/domain=m4aik.meetcplive.com",
	"ws-cdn-bps/domain=m4aik.meetstarlive.com",
	"ws-cdn-bps/domain=m4aik.starstarlive.com",
	"ws-cdn-bps/domain=m4aky.hnyapu.cn",
	"ws-cdn-bps/domain=m4akyik.hnyapu.cn",
	"ws-cdn-bps/domain=m4akyik.xiangshengclub.com",
	"ws-cdn-bps/domain=media.hnyapu.cn",
	"ws-cdn-bps/domain=media.meetstarlive.com",
	"ws-cdn-bps/domain=media.xiangshengclub.com",
	"ws-cdn-bps/domain=media.zhishichaoren.cn",
	"ws-cdn-bps/domain=medtx.ikstatic.cn",
	"ws-cdn-bps/domain=meelivepicture.inke.cn",
	"ws-cdn-bps/domain=meelivepicture2.inke.cn",
	"ws-cdn-bps/domain=meelivepicture6.inke.cn",
	"ws-cdn-bps/domain=mini-pull-ws.inke.cn",
	"ws-cdn-bps/domain=mlive.inke.cn",
	"ws-cdn-bps/domain=mlive10.inke.cn",
	"ws-cdn-bps/domain=mlive11.inke.cn",
	"ws-cdn-bps/domain=mlive12.inke.cn",
	"ws-cdn-bps/domain=mlive13.inke.cn",
	"ws-cdn-bps/domain=mlive14.inke.cn",
	"ws-cdn-bps/domain=mlive15.inke.cn",
	"ws-cdn-bps/domain=mlive16.inke.cn",
	"ws-cdn-bps/domain=mlive17.inke.cn",
	"ws-cdn-bps/domain=mlive18.inke.cn",
	"ws-cdn-bps/domain=mlive19.inke.cn",
	"ws-cdn-bps/domain=mlive2.inke.cn",
	"ws-cdn-bps/domain=mlive20.inke.cn",
	"ws-cdn-bps/domain=mlive21.inke.cn",
	"ws-cdn-bps/domain=mlive22.inke.cn",
	"ws-cdn-bps/domain=mlive23.inke.cn",
	"ws-cdn-bps/domain=mlive24.inke.cn",
	"ws-cdn-bps/domain=mlive25.inke.cn",
	"ws-cdn-bps/domain=mlive26.inke.cn",
	"ws-cdn-bps/domain=mlive27.inke.cn",
	"ws-cdn-bps/domain=mlive3.inke.cn",
	"ws-cdn-bps/domain=mlive4.inke.cn",
	"ws-cdn-bps/domain=mlive5.inke.cn",
	"ws-cdn-bps/domain=mlive6.inke.cn",
	"ws-cdn-bps/domain=mlive7.inke.cn",
	"ws-cdn-bps/domain=mlive8.inke.cn",
	"ws-cdn-bps/domain=mlive9.inke.cn",
	"ws-cdn-bps/domain=oper.yingtaorelian.com",
	"ws-cdn-bps/domain=publishws.8686c.com",
	"ws-cdn-bps/domain=publishws.imilive.cn",
	"ws-cdn-bps/domain=pull.inke.cn",
	"ws-cdn-bps/domain=pull.starstarlive.com",
	"ws-cdn-bps/domain=pull2.inke.cn",
	"ws-cdn-bps/domain=pull6.inke.cn",
	"ws-cdn-bps/domain=pullhls.inke.cn",
	"ws-cdn-bps/domain=pullhls.starstarlive.com",
	"ws-cdn-bps/domain=pullhls2.inke.cn",
	"ws-cdn-bps/domain=pullhls6.inke.cn",
	"ws-cdn-bps/domain=push.starstarlive.com",
	"ws-cdn-bps/domain=qq-bj-bgp.push.inke.cn",
	"ws-cdn-bps/domain=qq.orig.inke.cn",
	"ws-cdn-bps/domain=qq.video.imilive.cn",
	"ws-cdn-bps/domain=qqfeed.inke.cn",
	"ws-cdn-bps/domain=record.boluohuyu.cn",
	"ws-cdn-bps/domain=record.caratsvip.com",
	"ws-cdn-bps/domain=record.hnyapu.cn",
	"ws-cdn-bps/domain=record.inke.cn",
	"ws-cdn-bps/domain=record.xiangshengclub.com",
	"ws-cdn-bps/domain=record.yingtaorelian.com",
	"ws-cdn-bps/domain=record2.xiangshengclub.com",
	"ws-cdn-bps/domain=recordky.hnyapu.cn",
	"ws-cdn-bps/domain=resource.tvhoo.com",
	"ws-cdn-bps/domain=srv.faceshop.inke.cn",
	"ws-cdn-bps/domain=starstarlive.8686c.com",
	"ws-cdn-bps/domain=static.boluohuyu.cn",
	"ws-cdn-bps/domain=static.gzlywlkj8.com",
	"ws-cdn-bps/domain=static.hnyapu.cn",
	"ws-cdn-bps/domain=static.hnyplive.cn",
	"ws-cdn-bps/domain=static.hnypvoice.cn",
	"ws-cdn-bps/domain=static.ikbase.cn",
	"ws-cdn-bps/domain=static.inke.cn",
	"ws-cdn-bps/domain=static.inke.com",
	"ws-cdn-bps/domain=static.starstarlive.com",
	"ws-cdn-bps/domain=testpull.inke.cn",
	"ws-cdn-bps/domain=testpush.inke.cn",
	"ws-cdn-bps/domain=video.anyueclub.com",
	"ws-cdn-bps/domain=video.inke.8686c.com",
	"ws-cdn-bps/domain=videostatic.imilive.cn",
	"ws-cdn-bps/domain=voc.anyueclub.com",
	"ws-cdn-bps/domain=voc.hnyapu.cn",
	"ws-cdn-bps/domain=voc.hnyplive.cn",
	"ws-cdn-bps/domain=voc.hnypvoice.cn",
	"ws-cdn-bps/domain=voc.xiangshengclub.com",
	"ws-cdn-bps/domain=voc.zhishichaoren.cn",
	"ws-cdn-bps/domain=vocik.hnyapu.cn",
	"ws-cdn-bps/domain=vocik.hnyplive.cn",
	"ws-cdn-bps/domain=vocik.hnypvoice.cn",
	"ws-cdn-bps/domain=vocik.meetcplive.com",
	"ws-cdn-bps/domain=vocjyik.gzhuda.com",
	"ws-cdn-bps/domain=vocky.hnyapu.cn",
	"ws-cdn-bps/domain=vockyik.hnyapu.cn",
	"ws-cdn-bps/domain=vockyik.meetstarlive.com",
	"ws-cdn-bps/domain=vockyik.xiangshengclub.com",
	"ws-cdn-bps/domain=voctx.ikstatic.cn",
	"ws-cdn-bps/domain=voice.yaamari.com",
	"ws-cdn-bps/domain=voicetime.anyueclub.com",
	"ws-cdn-bps/domain=webcdn.ikbase.cn",
	"ws-cdn-bps/domain=webcdn.inke.cn",
	"ws-cdn-bps/domain=webcdn.starstarlive.com",
	"ws-cdn-bps/domain=webstatic.ikbase.cn",
	"ws-cdn-bps/domain=webstatic.imilive.cn",
	"ws-cdn-bps/domain=webstatic.inke.cn",
	"ws-cdn-bps/domain=ws.img.cheesesuperman.com",
	"ws-cdn-bps/domain=ws.img.imilive.cn",
	"ws-cdn-bps/domain=ws.video.imilive.cn",
	"ws-cdn-bps/domain=wsaudio.hls.inke.cn",
	"ws-cdn-bps/domain=wsaudio.pull.inke.cn",
	"ws-cdn-bps/domain=wsaudio.rtc.inke.cn",
	"ws-cdn-bps/domain=wsaudio.rtc.meetstarlive.com",
	"ws-cdn-bps/domain=wschorus.pull.inke.cn",
	"ws-cdn-bps/domain=wsfeed.inke.cn",
	"ws-cdn-bps/domain=wsimg.cheesesuperman.com",
	"ws-cdn-bps/domain=wsimg.imilive.cn",
	"ws-cdn-bps/domain=wsm4a.inke.cn",
	"ws-cdn-bps/domain=wsrecord.inke.cn",
	"ws-cdn-bps/domain=wssource.audio.hnyapu.cn",
	"ws-cdn-bps/domain=wssource.audio.inke.cn",
	"ws-cdn-bps/domain=wssource.audio.xiangshengclub.com",
	"ws-cdn-bps/domain=wssource.audioky.hnyapu.cn",
	"ws-cdn-bps/domain=wssource.hls.inke.cn",
	"ws-cdn-bps/domain=wssource.pull.ichaoren.com",
	"ws-cdn-bps/domain=wssource.pull.inke.cn",
	"ws-cdn-bps/domain=wssource.rtc.9zhenge.com",
	"ws-cdn-bps/domain=wssource.rtc.anyueclub.com",
	"ws-cdn-bps/domain=wssource.rtc.boomim.cn",
	"ws-cdn-bps/domain=wssource.rtc.caratsvip.com",
	"ws-cdn-bps/domain=wssource.rtc.fengyuhn.cn",
	"ws-cdn-bps/domain=wssource.rtc.gmugmu.com",
	"ws-cdn-bps/domain=wssource.rtc.inke.cn",
	"ws-cdn-bps/domain=wssource.rtc.meelove.cn",
	"ws-cdn-bps/domain=wssource.temp.inke.cn",
	"ws-cdn-bps/domain=wssource.test.inke.cn",
	"ws-cdn-bps/domain=wssource.test2.inke.cn",
	"ws-cdn-bps/domain=wssource.vs.inke.cn",
	"ws-cdn-bps/domain=wswebhls.inke.cn",
	"ws-cdn-bps/domain=wswebpull.inke.cn",
	"ws-cdn-bps/domain=wswebpull1.inke.cn",
	"ws-cdn-bps/domain=wswebpull2.inke.cn",
	"ws-cdn-bps/domain=wswebpull3.inke.cn",
	"ws-cdn-bps/domain=wswebpull4.inke.cn",
	"ws-cdn-bps/domain=wswebpull5.inke.cn",
	"ws-cdn-bps/domain=wswebpull6.inke.cn",
	"ws-cdn-bps/domain=wswebpull7.inke.cn",
	"ws-cdn-bps/domain=wswebpull8.inke.cn",
	"ws-cdn-bps/domain=wswebpull9.inke.cn",
}

var qq_cdn_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var qq_cdn_counters = []string{
	"tx-cdn-bps/domain=5331.livepush.myqcloud.com",
	"tx-cdn-bps/domain=ali.img.imilive.cn",
	"tx-cdn-bps/domain=ali.video.imilive.cn",
	"tx-cdn-bps/domain=app.anyueclub.com",
	"tx-cdn-bps/domain=app.boomim.cn",
	"tx-cdn-bps/domain=app.duiyuan520.com",
	"tx-cdn-bps/domain=app.inke.cn",
	"tx-cdn-bps/domain=avatar.cdn.gmugmu.com",
	"tx-cdn-bps/domain=bujiu.img.ikstatic.cn",
	"tx-cdn-bps/domain=chatvoc.inke.cn",
	"tx-cdn-bps/domain=client.resource.inke.cn",
	"tx-cdn-bps/domain=download.boluohuyu.cn",
	"tx-cdn-bps/domain=download.gmugmu.com",
	"tx-cdn-bps/domain=fx.img.imilive.cn",
	"tx-cdn-bps/domain=gaia-video-1252926420.file.myqcloud.com",
	"tx-cdn-bps/domain=gaia.imilive.cn",
	"tx-cdn-bps/domain=h5.anyueclub.com",
	"tx-cdn-bps/domain=h5.inke.cn",
	"tx-cdn-bps/domain=ik-feed.inke.cn",
	"tx-cdn-bps/domain=im-media-img.gmugmu.com",
	"tx-cdn-bps/domain=im-media-voc.gmugmu.com",
	"tx-cdn-bps/domain=image.scale.inke.cn",
	"tx-cdn-bps/domain=image.scale.inke.com",
	"tx-cdn-bps/domain=imagescale.9zhenge.com",
	"tx-cdn-bps/domain=imagescale.hnyplive.cn",
	"tx-cdn-bps/domain=imagescale.hnypvoice.cn",
	"tx-cdn-bps/domain=imagescale.imilive.cn",
	"tx-cdn-bps/domain=imagescale.inke.cn",
	"tx-cdn-bps/domain=imagescale.meetcplive.com",
	"tx-cdn-bps/domain=imagescale.starstarlive.com",
	"tx-cdn-bps/domain=imagescale.yingtaorelian.com",
	"tx-cdn-bps/domain=imagescaleik.meetstarlive.com",
	"tx-cdn-bps/domain=img-buylive.ikbase.cn",
	"tx-cdn-bps/domain=img.1yimi.cn",
	"tx-cdn-bps/domain=img.9zhenge.com",
	"tx-cdn-bps/domain=img.anyueclub.com",
	"tx-cdn-bps/domain=img.boluohuyu.cn",
	"tx-cdn-bps/domain=img.duiyuan520.com",
	"tx-cdn-bps/domain=img.gzhuda.com",
	"tx-cdn-bps/domain=img.hngaojia.cn",
	"tx-cdn-bps/domain=img.hnyapu.cn",
	"tx-cdn-bps/domain=img.hnyplive.cn",
	"tx-cdn-bps/domain=img.hnypvoice.cn",
	"tx-cdn-bps/domain=img.ikstatic.cn",
	"tx-cdn-bps/domain=img.ingkee.cn",
	"tx-cdn-bps/domain=img.inke.cn",
	"tx-cdn-bps/domain=img.iyyhd.com",
	"tx-cdn-bps/domain=img.lhzhiying.com",
	"tx-cdn-bps/domain=img.meelove.cn",
	"tx-cdn-bps/domain=img.meetcplive.com",
	"tx-cdn-bps/domain=img.thlrs.com",
	"tx-cdn-bps/domain=img.xiangshengclub.com",
	"tx-cdn-bps/domain=img.yingtaorelian.com",
	"tx-cdn-bps/domain=img.zhishichaoren.cn",
	"tx-cdn-bps/domain=img2.inke.cn",
	"tx-cdn-bps/domain=img2ik.boluohuyu.cn",
	"tx-cdn-bps/domain=img2ik.meetcplive.com",
	"tx-cdn-bps/domain=img2ik.meetstarlive.com",
	"tx-cdn-bps/domain=img2kyik.gzsnxxkj8.com",
	"tx-cdn-bps/domain=img2kyik.hnyapu.cn",
	"tx-cdn-bps/domain=imgky.hnyapu.cn",
	"tx-cdn-bps/domain=imgstatickyik.hnyapu.cn",
	"tx-cdn-bps/domain=imgstatisticik.boluohuyu.cn",
	"tx-cdn-bps/domain=imgstatisticik.meetcplive.com",
	"tx-cdn-bps/domain=imgtx.ikstatic.cn",
	"tx-cdn-bps/domain=imgyimi.bjdyzg.com",
	"tx-cdn-bps/domain=jsrecord.inke.cn",
	"tx-cdn-bps/domain=lipstickh5.inke.cn",
	"tx-cdn-bps/domain=live-hls-qq.buylivehk.com",
	"tx-cdn-bps/domain=live-pull-hls.buylivehk.com",
	"tx-cdn-bps/domain=live-pull-qq.buylivehk.com",
	"tx-cdn-bps/domain=live-push-qq.buylivehk.com",
	"tx-cdn-bps/domain=lrc.inke.cn",
	"tx-cdn-bps/domain=lyrtx.ikstatic.cn",
	"tx-cdn-bps/domain=m4a.9zhenge.com",
	"tx-cdn-bps/domain=m4a.boluohuyu.cn",
	"tx-cdn-bps/domain=m4a.duiyuan520.com",
	"tx-cdn-bps/domain=m4a.hngaojia.cn",
	"tx-cdn-bps/domain=m4a.hnyapu.cn",
	"tx-cdn-bps/domain=m4a.hnyplive.cn",
	"tx-cdn-bps/domain=m4a.hnypvoice.cn",
	"tx-cdn-bps/domain=m4a.inke.cn",
	"tx-cdn-bps/domain=m4a.iyyhd.com",
	"tx-cdn-bps/domain=m4a.xiangshengclub.com",
	"tx-cdn-bps/domain=m4a.yingtaorelian.com",
	"tx-cdn-bps/domain=m4a2.inke.cn",
	"tx-cdn-bps/domain=m4aik.boluohuyu.cn",
	"tx-cdn-bps/domain=m4aik.meetcplive.com",
	"tx-cdn-bps/domain=m4aik.meetstarlive.com",
	"tx-cdn-bps/domain=m4akyik.gzsnxxkj8.com",
	"tx-cdn-bps/domain=media.hnyapu.cn",
	"tx-cdn-bps/domain=media.meetstarlive.com",
	"tx-cdn-bps/domain=media.xiangshengclub.com",
	"tx-cdn-bps/domain=media.zhishichaoren.cn",
	"tx-cdn-bps/domain=media01.inke.cn",
	"tx-cdn-bps/domain=media01.inke.com",
	"tx-cdn-bps/domain=media02.inke.cn",
	"tx-cdn-bps/domain=media02.inke.com",
	"tx-cdn-bps/domain=medtx.ikstatic.cn",
	"tx-cdn-bps/domain=miniimg.busi.inke.cn",
	"tx-cdn-bps/domain=minim4a.busi.inke.cn",
	"tx-cdn-bps/domain=music.inke.cn",
	"tx-cdn-bps/domain=oper.duiyuan520.com",
	"tx-cdn-bps/domain=oper.yingtaorelian.com",
	"tx-cdn-bps/domain=qq.chrousimg.ichaoren.com",
	"tx-cdn-bps/domain=qq.video.imilive.cn",
	"tx-cdn-bps/domain=qqaudio.hls.inke.cn",
	"tx-cdn-bps/domain=qqaudio.pull.inke.cn",
	"tx-cdn-bps/domain=qqaudio.push.inke.cn",
	"tx-cdn-bps/domain=qqfeed.inke.cn",
	"tx-cdn-bps/domain=qqimg.ichaoren.com",
	"tx-cdn-bps/domain=qqpicture.inke.cn",
	"tx-cdn-bps/domain=qqpull.inke.cn",
	"tx-cdn-bps/domain=qqpull2.inke.cn",
	"tx-cdn-bps/domain=qqpullhls.inke.cn",
	"tx-cdn-bps/domain=qqpullhls2.inke.cn",
	"tx-cdn-bps/domain=qqpush.inke.cn",
	"tx-cdn-bps/domain=qqpush2.inke.cn",
	"tx-cdn-bps/domain=qqshortvideo201705311745-1252926420.file.myqcloud.com",
	"tx-cdn-bps/domain=qqsource.pull.inke.cn",
	"tx-cdn-bps/domain=qqsource.rtc.inke.cn",
	"tx-cdn-bps/domain=qqsource.tmp.inke.cn",
	"tx-cdn-bps/domain=qqwebhls.inke.cn",
	"tx-cdn-bps/domain=record.boluohuyu.cn",
	"tx-cdn-bps/domain=record.duiyuan520.com",
	"tx-cdn-bps/domain=record.fengyuhn.cn",
	"tx-cdn-bps/domain=record.ichaoren.com",
	"tx-cdn-bps/domain=record.inke.cn",
	"tx-cdn-bps/domain=record.iyyhd.com",
	"tx-cdn-bps/domain=record.yingtaorelian.com",
	"tx-cdn-bps/domain=record2.inke.cn",
	"tx-cdn-bps/domain=recordqa.inke.cn",
	"tx-cdn-bps/domain=sdk-test-file-1252926420.file.myqcloud.com",
	"tx-cdn-bps/domain=srv.faceshop.inke.cn",
	"tx-cdn-bps/domain=static.boluohuyu.cn",
	"tx-cdn-bps/domain=static.hnyapu.cn",
	"tx-cdn-bps/domain=static.hnyplive.cn",
	"tx-cdn-bps/domain=static.hnypvoice.cn",
	"tx-cdn-bps/domain=static.ikbase.cn",
	"tx-cdn-bps/domain=static.inke.cn",
	"tx-cdn-bps/domain=static.inke.com",
	"tx-cdn-bps/domain=tx.img.imilive.cn",
	"tx-cdn-bps/domain=tx.video.imilive.cn",
	"tx-cdn-bps/domain=tzkhjh5aa.inke.cn",
	"tx-cdn-bps/domain=video.anyueclub.com",
	"tx-cdn-bps/domain=voc.hnyapu.cn",
	"tx-cdn-bps/domain=voc.hnyplive.cn",
	"tx-cdn-bps/domain=voc.hnypvoice.cn",
	"tx-cdn-bps/domain=voc.inke.cn",
	"tx-cdn-bps/domain=voc.xiangshengclub.com",
	"tx-cdn-bps/domain=voc.zhishichaoren.cn",
	"tx-cdn-bps/domain=vocik.boluohuyu.cn",
	"tx-cdn-bps/domain=vocik.hnyplive.cn",
	"tx-cdn-bps/domain=vocik.hnypvoice.cn",
	"tx-cdn-bps/domain=vocik.meetcplive.com",
	"tx-cdn-bps/domain=vockyik.gzsnxxkj8.com",
	"tx-cdn-bps/domain=vockyik.meetstarlive.com",
	"tx-cdn-bps/domain=voctx.ikstatic.cn",
	"tx-cdn-bps/domain=wb.record.inke.cn",
	"tx-cdn-bps/domain=webcdn.ikbase.cn",
	"tx-cdn-bps/domain=webcdn.inke.cn",
	"tx-cdn-bps/domain=webstatic.ikbase.cn",
	"tx-cdn-bps/domain=webstatic.imilive.cn",
	"tx-cdn-bps/domain=webstatic.inke.cn",
	"tx-cdn-bps/domain=ws.img.cheesesuperman.com",
	"tx-cdn-bps/domain=ws.img.imilive.cn",
	"tx-cdn-bps/domain=wsfeed.inke.cn",
	"tx-cdn-bps/domain=wsimg.cheesesuperman.com",
	"tx-cdn-bps/domain=wsimg.imilive.cn",
	"tx-cdn-bps/domain=wsrecord.inke.cn",
}

var csy_cdn_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var csy_cdn_counters = []string{
	"csydwall_csysource.rtc.inke.cn",
}

var cls_livego_edge_host = []string{
	"tx3-media-cls-disp01.bj",
}
var cls_livego_edge_counters = []string{
	"livegoedge-bw_send/hostname=北方边缘吉林移动222.34.54.142",
	"livegoedge-bw_send/hostname=北方边缘吉林联通119.52.228.92",
	"livegoedge-bw_send/hostname=北方边缘山东移动120.221.66.209",
	"livegoedge-bw_send/hostname=北方边缘山东联通124.134.126.154",
	"livegoedge-bw_send/hostname=北方边缘河南移动111.6.188.56",
	"livegoedge-bw_send/hostname=北方边缘河南联通219.157.114.217",
	"livegoedge-bw_send/hostname=北方边缘辽宁移动221.180.218.231",
	"livegoedge-bw_send/hostname=北方边缘辽宁联通124.95.149.244",
	"livegoedge-bw_send/hostname=北方边缘黑龙江移动111.42.180.18",
	"livegoedge-bw_send/hostname=北方边缘黑龙江联通218.7.198.208",
	"livegoedge-bw_send/hostname=南方边缘广东电信(02)121.32.236.55",
	"livegoedge-bw_send/hostname=南方边缘广东电信121.32.236.40",
	"livegoedge-bw_send/hostname=南方边缘广东移动120.233.1.66",
	"livegoedge-bw_send/hostname=南方边缘江苏电信0258.220.44.210",
	"livegoedge-bw_send/hostname=南方边缘江苏电信58.220.44.236",
	"livegoedge-bw_send/hostname=南方边缘江苏移动223.111.192.123",
	"livegoedge-bw_send/hostname=南方边缘浙江电信183.134.11.141",
	"livegoedge-bw_send/hostname=南方边缘浙江移动117.149.154.222",
	"livegoedge-bw_send/hostname=南方边缘湖北电信116.211.79.36",
	"livegoedge-bw_send/hostname=南方边缘湖北移动111.47.221.158",
	"livegoedge-bw_send/hostname=南方边缘福建电信125.77.133.204",
	"livegoedge-bw_send/hostname=南方边缘福建移动112.5.63.177",
}

var cls_livego_centre_host = []string{
	"tx3-media-cls-disp01.bj",
}
var cls_livego_centre_counters = []string{
	"livegocenter-bw_send/hostname=北方中心河北三线(主1)111.225.214.118",
	"livegocenter-bw_send/hostname=北方中心河北三线(主2)111.225.214.121",
	"livegocenter-bw_send/hostname=北方中心河北三线(主3)111.225.214.119",
	"livegocenter-bw_send/hostname=北方中心河北三线(主4)111.225.214.122",
	"livegocenter-bw_send/hostname=南方中心常州三线(主1)118.184.169.114",
	"livegocenter-bw_send/hostname=南方中心常州三线(主2)118.184.169.113",
	"livegocenter-bw_send/hostname=南方中心常州三线(主3)118.184.169.115",
	"livegocenter-bw_send/hostname=南方中心常州三线(主4)118.184.169.126",
}

var cls_livego_cdn_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var cls_livego_cdn_counters = []string{
	"ws_vmp_out/hostname=inke-changzhou2_0001",
	"ws_vmp_out/hostname=inke-changzhou3_0001",
	"ws_vmp_out/hostname=inke-changzhou4_0001",
	"ws_vmp_out/hostname=inke-changzhou_0001",
	"ws_vmp_out/hostname=inke-haerbin-yd01_0001",
	"ws_vmp_out/hostname=inke-huizhou-yd0001_0001",
	"ws_vmp_out/hostname=inke-langfang3_0001",
	"ws_vmp_out/hostname=inke-langfang4_0001",
	"ws_vmp_out/hostname=inke-xiangyang-yd01_0001",
	"ws_vmp_out/hostname=inke_changchun_0001",
	"ws_vmp_out/hostname=inke_jiangmen2_0001",
	"ws_vmp_out/hostname=inke_jiangmen_0001",
	"ws_vmp_out/hostname=inke_jiaxing_0001",
	"ws_vmp_out/hostname=inke_langfang1_0001",
	"ws_vmp_out/hostname=inke_langfang2_0001",
	"ws_vmp_out/hostname=inke_lianyungang_0001",
	"ws_vmp_out/hostname=inke_luohe_0001",
	"ws_vmp_out/hostname=inke_ningbo_0001",
	"ws_vmp_out/hostname=inke_qingdao_yd1_0001",
	"ws_vmp_out/hostname=inke_quanzhou_0001",
	"ws_vmp_out/hostname=inke_quanzhoudx_0001",
	"ws_vmp_out/hostname=inke_shenyang_0001",
	"ws_vmp_out/hostname=inke_shenyang_wt02_0001",
	"ws_vmp_out/hostname=inke_suihua_0001",
	"ws_vmp_out/hostname=inke_tonghua_wt01_0001",
	"ws_vmp_out/hostname=inke_xiaogan_0001",
	"ws_vmp_out/hostname=inke_yangzhou2_0001",
	"ws_vmp_out/hostname=inke_yangzhou_0001",
	"ws_vmp_out/hostname=inke_zhengzhou_0001",
	"ws_vmp_out/hostname=inke_zibo_0001",
	"ws_vmp_out/hostname=ws_langfang3_lvs1_hb_0001",
	"ws_vmp_out/hostname=ws_langfang3_lvs2_hb",
}

var zego_audio_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var zego_audio_counters = []string{
	"cdn-zego/class=bandwidth,desc=77,domain=hdl-ws.meetstarlive.com,type=cdn",
	"cdn-zego/class=bandwidth,desc=77,domain=publish-ws.meetstarlive.com,type=cdn",
	"cdn-zego/class=bandwidth,desc=77,domain=rtmp-ws.meetstarlive.com,type=cdn",
	"cdn-zego/class=bandwidth,desc=bujiu,domain=play-aliyun.anyueclub.com,type=cdn",
	"cdn-zego/class=bandwidth,desc=duiyuan,domain=play-aliyun.yingtaorelian.com,type=cdn",
	"cdn-zego/class=bandwidth,desc=duiyuan,domain=publish-aliyun.boluohuyu.cn,type=cdn",
	"cdn-zego/class=bandwidth,desc=vins,domain=hdl-ws.vins.live,type=cdn",
	"cdn-zego/class=bandwidth,desc=vins,domain=publish-ws.vins.live,type=cdn",
	"cdn-zego/class=bandwidth,desc=vins,domain=rtmp-ws.vins.live,type=cdn",
	"cdn-zego/class=bandwidth,desc=yingke,domain=pic.ws.inchat.yitu.me,type=cdn",
	"cdn-zego/class=bandwidth,desc=yingke,domain=zg-ws.hdl.inke.cn,type=cdn",
	"cdn-zego/class=bandwidth,desc=yingke,domain=zg-ws.push.inke.cn,type=cdn",
	"cdn-zego/class=bandwidth,desc=yingke,domain=zg-ws.rtmp.inke.cn,type=cdn",
}

var zego_cdn_host = []string{
	"ali-a-ops-lmmonitor01.bj",
	"ali-a-pe-cdn-monitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}
var zego_cdn_counters = []string{
	"cdn-zego/class=bandwidth,desc=yingke,type=lm",
	"cdn-zego/class=bandwidth,desc=duiyuan,type=lm",
	"cdn-zego/class=bandwidth,desc=77,type=lm",
	"cdn-zego/class=bandwidth,desc=vins,type=lm",
	"cdn-zego/class=bandwidth,desc=jimu,type=lm",
	"cdn-zego/class=bandwidth,desc=shengquan,type=lm",
	"cdn-zego/class=bandwidth,desc=bujiu,type=lm",
	"cdn-zego/class=bandwidth,desc=starstar,type=lm",
}

var zego_microphone_host = []string{
	"tx-k8s-media-zego-task-stat01.bj",
}

var zego_microphone_counters = []string{
	/*"zego__master_users_four",
	"zego__master_users_nine",
	"zego__master_users_six",
	"zego__softvoice_users_nine_sv",*/

	tx-k8s-media-zego-task-stat02.bj zego__master_users_four
	tx-k8s-media-zego-task-stat02.bj zego__master_users_nine
	tx-k8s-media-zego-task-stat01.bj zego__master_users_four
	tx-k8s-media-zego-task-stat01.bj zego__master_users_nine
	tx-k8s-media-zego-task-stat01.bj zego__master_users_six
	tx-k8s-media-zego-task-stat01.bj zego__softvoice_users_nine_sv
	tx-k8s-media-zego-task-stat02.bj zego__master_users_six
	tx-k8s-media-zego-task-stat02.bj zego__softvoice_users_nine_sv
	tx-k8s-media-zego-task-stat03.bj zego__master_users_four
	tx-k8s-media-zego-task-stat03.bj zego__master_users_nine
	tx-k8s-media-zego-task-stat03.bj zego__master_users_six
	tx-k8s-media-zego-task-stat03.bj zego__softvoice_users_nine_sv
}

var cdn_play_host = []string{
	"ali-a-pe-cdn-monitor01.bj",
	"ali-a-ops-lmmonitor01.bj",
	"tx3-pe-cdn-monitor01.bj",
}

var cdn_play_counters = []string{
	"metric=ali-cdn-bps",  // 阿里
	"metric=ws-cdn-bps",   // ws
	"metric=tx-cdn-bps",   // tx
}



func deal_data(origin data_return) float64 {
	var slice []Slices
	var add_result add
	var all float64 = 0
	var max float64 = 0
	len := 0
	//sum_temp := 0
	for _, row1 := range origin {
		var value_temp = value{}
		var slices_temp = Slices{}
		var sum float64 = 0
		var temp float64 = 0
		slices_temp.endpoint = row1.Endpoint
		//println(row1.Endpoint)
		len = 0
		for _, row2 := range row1.Values {
			sum += row2.Value
			temp += row2.Value

			value_temp.timestamp = row2.Timestamp
			value_temp.value = temp

			slices_temp.values = append(slices_temp.values, value_temp)
			temp = 0
			len++
		}
		slice = append(slice, slices_temp)
		//fmt.Println(slices_temp)
		row1.sum = sum
		all += sum
		sum = 0
	}
	for i := 0; i < len; i++ {
		var sum float64 = 0
		for _, row := range slice {
			if cap(row.values) > i && cap(row.values) != 0 {
				sum += row.values[i].value
				add_result.timestamp = row.values[i].timestamp
			}
		}
		fmt.Println("sum:", sum)
		if max < sum {
			max = sum
		}
		fmt.Println("max:", max)
		add_result.value = append(add_result.value, sum)
		//println(sum)
		sum = 0
	}

	//for _, row1 := range slice {
	//  fmt.Println(row1.endpoint)
	//  for _,row2 := range row1.values{
	//    println(int(row2.value))
	//  }
	//}
	return max
}

type add struct {
	value     []float64
	timestamp int
}

type value struct {
	value     float64
	timestamp int
}
