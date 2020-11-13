package service

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/dao"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/manager"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"context"
	"fmt"
)

type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

func New(c *conf.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
	if s.mgr != nil {
		s.mgr.Close()
	}
}

/*

适用于两个过滤条件+求avg，max等条件情况
例如：
queryInfo := HostName:conf.Conf.EsHost.CdnPullOpenHost,
			Group1: "isp",
			Group2: "cv",
			Target: "open_success_t",
			Method: "avg",
			Minute: 1,
ExeUntreatedQueryEs(queryInfo model.UntreatedConf) = SQL: select avg(open_success_t) from tbl between 前(3+1)分钟 and 前3分钟 group by app,cv;
*/
func ExeUntreatedQueryEs(queryInfo model.UntreatedConf) {

	switch queryInfo.GroupFilter {
	case model.GF20:
		GF20QueryEsAndInsertMySql(queryInfo)
	case model.GF30:
		GF30QueryEsAndInsertMySql(queryInfo)
	case model.GF40:
		GF40QueryEsAndInsertMySql(queryInfo)
	case model.GF50:
		GF50QueryEsAndInsertMySql(queryInfo)
	default:
		log.Errorf("ExeUntreatedQueryEs  is GroupFilter error! GroupFilter:%s", queryInfo.GroupFilter)
	}
}

func GF50QueryEsAndInsertMySql(utConf model.UntreatedConf) {
	//GF40QueryEsAndInsertMySql(utConf)

	fmt.Println("GF50QueryEsAndInsertMySql")
	GF50PostBody := PostJoint(utConf,model.Level5)

	var GF50MysqlInfos []model.MysqUntreatedInfo
	var EsResultGF50 model.UntreatedJsonRes

	SendEsPost(GF50PostBody, &EsResultGF50, utConf.HostName)

	dao.CollectEsRes2Mysql(&GF50MysqlInfos, utConf, &EsResultGF50, model.Level5)
	//插入数据库
	printInfos(GF50MysqlInfos)
	//dao.InsertMysql(&GF50MysqlInfos, utConf.TblName)
}

func GF40QueryEsAndInsertMySql(utConf model.UntreatedConf) {
	GF30QueryEsAndInsertMySql(utConf)

	fmt.Println("GF40QueryEsAndInsertMySql")
	GF40PostBody := PostJoint(utConf,model.Level4)

	var GF40MysqlInfos []model.MysqUntreatedInfo
	var EsResultGF40 model.UntreatedJsonRes

	SendEsPost(GF40PostBody, &EsResultGF40, utConf.HostName)

	dao.CollectEsRes2Mysql(&GF40MysqlInfos, utConf, &EsResultGF40, model.Level4)
	//插入数据库
	printInfos(GF40MysqlInfos)
	//dao.InsertMysql(&GF40MysqlInfos, utConf.TblName)
}

func GF30QueryEsAndInsertMySql(utConf model.UntreatedConf) {

	GF20QueryEsAndInsertMySql(utConf)

	fmt.Println("GF30QueryEsAndInsertMySql")
	GF30PostBody := PostJoint(utConf,model.Level3)

	var GF30MysqlInfos []model.MysqUntreatedInfo
	var EsResultGF30 model.UntreatedJsonRes

	SendEsPost(GF30PostBody, &EsResultGF30, utConf.HostName)

	dao.CollectEsRes2Mysql(&GF30MysqlInfos, utConf, &EsResultGF30, model.Level3)
	//插入数据库
	printInfos(GF30MysqlInfos)
	//dao.InsertMysql(&GF30MysqlInfos, utConf.TblName)

}

func GF20QueryEsAndInsertMySql(utConf model.UntreatedConf) {
	GF10QueryEsAndInsertMySql(utConf)
	fmt.Println("GF20QueryEsAndInsertMySql")

	GF20PostBody := PostJoint(utConf,model.Level2)

	var GF20MysqlInfos []model.MysqUntreatedInfo

	var EsResultGF20 model.UntreatedJsonRes

	SendEsPost(GF20PostBody, &EsResultGF20, utConf.HostName)

	dao.CollectEsRes2Mysql(&GF20MysqlInfos, utConf, &EsResultGF20, model.Level2)
	//插入数据库
	printInfos(GF20MysqlInfos)
	//dao.InsertMysql(&GF20MysqlInfos, utConf.TblName)
}

func GF10QueryEsAndInsertMySql(utConf model.UntreatedConf) {

	fmt.Println("GF10QueryEsAndInsertMySql")
	GF10PostBody := PostJoint(utConf,model.Level1)
	var EsResultGF10 model.UntreatedJsonRes
	var GF10MysqlInfos []model.MysqUntreatedInfo

	SendEsPost(GF10PostBody, &EsResultGF10, utConf.HostName)

	dao.CollectEsRes2Mysql(&GF10MysqlInfos, utConf, &EsResultGF10, model.Level1)
	printInfos(GF10MysqlInfos)
	//dao.InsertMysql(&GF10MysqlInfos, utConf.TblName)
}

func printInfos(infos []model.MysqUntreatedInfo) {
	for _, info := range infos {
		fmt.Printf("Ymd:%s\t All:%f\t G1:%s\t G1N:%d\t G2:%s\t G2N:%d\t G3:%s\t G3N:%d\t G4:%s\t G4N:%d\t G5:%s\t G5N:%d\t G6:%s\t G6N:%d\t Value1:%f\t Value2:%f\t Value3:%f\t Value4:%f\t Value5:%f\t Value6:%f\t Event_time:%d\n",
			info.Ymd, info.All, info.Group1, info.Group1_Num, info.Group2, info.Group2_Num, info.Group3, info.Group3_Num,
			info.Group4, info.Group4_Num, info.Group5, info.Group5_Num, info.Value1, info.Value2,info.Value3, info.Value4,info.Value5, info.Value6, info.Event_Time)
	}
}

//遍历queryInfos，对每个queryInfo进行查询ES操作
func StartUntreatedES(queryInfos []model.UntreatedConf) {
	for _, queryInfo := range queryInfos {
		ExeUntreatedQueryEs(queryInfo)
	}
}

//queryHandleInfos，对每个queryHandleInfo进行查询ES操作
func StartHandleES(queryHandleInfos []model.HandleConf) {
	for _, handleInfo := range queryHandleInfos {
		ExeHandleQueryEs(handleInfo)
	}
}

func ExeHandleQueryEs(queryHandleInfo model.HandleConf) {
	switch queryHandleInfo.GroupFilter {
	case model.GF10:
		GF10QueryHandleResToMysql(queryHandleInfo)
	default:
		log.Errorf("ExeUntreatedQueryEs  is GroupFilter error! GroupFilter:%s", queryHandleInfo.GroupFilter)
	}
}

func QueryOriginalResToMysql(originalConf model.OriginalConf) {
	postBody := GetBodyByOriginal(originalConf)

	var MysqlInfos []model.MysqOriginalInfo

	var EsResult model.OriginalJsonRes

	SendEsPost(postBody, &EsResult, originalConf.HostName)

	dao.CollectEsOriginalRes(&MysqlInfos, originalConf, EsResult)
	//插入数据库
	printInfos3(MysqlInfos)
	dao.InsertOriginalMysql(&MysqlInfos, conf.Conf.SqlTblName.EsDealResTblName)
}

func printInfos3(infos []model.MysqOriginalInfo) {

}

//从ES查询原始数据，对数据选择做处理（求avg，max等操作）还是不做处理，导入MYSQL
func GF10QueryHandleResToMysql(handleInfo model.HandleConf) {
	var postBody string
	switch handleInfo.TargetNum {
	case model.TarNum2:
		postBody = GetBodyByHandleGF10T2(handleInfo)
	default:
		log.Errorf("GF10QueryHandleResToMysql is TarNum error! GroupFilter:%s", handleInfo.TargetNum)
	}

	var MysqlInfos []model.MysqHandleInfo

	var EsResult model.HandleJsonRes

	SendEsPost(postBody, &EsResult, conf.Conf.EsHost.PullHost)

	dao.CollectEsHandleRes(&MysqlInfos, handleInfo, EsResult)
	//插入数据库
	printInfos2(MysqlInfos)
	dao.InsertHandleMysql(&MysqlInfos, conf.Conf.SqlTblName.EsDealResTblName)
}

func printInfos2(infos []model.MysqHandleInfo) {
	for _, info := range infos {
		fmt.Println(info)
	}
}
