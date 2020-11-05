package server

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/service"

)



var (
	queryInfos []model.QueryInfo
)

func StartServer() {
	////添加条件信息到queryInfos
	service.AddQueryEsInfos(&queryInfos)

	//查询实时数据
	service.StartQueryES(queryInfos)

	//插入数据库
	InsertMySQL()
}


func InsertMySQL() {

}

func RangeQueryES1() {



	//postBody := ExeQueryEs("isp","cv","open_success_t","avg",1)
	//postBody := service.ExeQueryEs("app", "cv", "lostA_avg", "avg", 1)



}


