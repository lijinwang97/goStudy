package server

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/service"
)



var (
	untreatedConfs []model.UntreatedConf
	handleConfs    []model.HandleConf
	originalConf    model.OriginalConf

)

func StartServer() {
	// 通过输入条件来查询ES未处理数据，导入MySQL
	StartQueryUntreatedES()

	//查询ES处理后的数据，导入MySQL
	//StartQueryOriginalES()

	//查询ES处理后的数据，继续处理，导入MySQL
	//StartQueryHandleES()
}

func StartQueryOriginalES() {
	conf.AddOriginalConfs(&originalConf)
	service.QueryOriginalResToMysql(originalConf)
}



func StartQueryHandleES() {
	conf.AddHandleConfs(&handleConfs)
	service.StartHandleES(handleConfs)
}

func StartQueryUntreatedES() {
	//添加条件信息到untreatedInfos
	conf.AddUntreatedConfs(&untreatedConfs)
	//查询实时数据
	service.StartUntreatedES(untreatedConfs)
}


