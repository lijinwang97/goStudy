package service

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/dao"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/manager"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/utils"
	"context"
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

const (
	GF20 = 20 // 20 表表示2个分组条件，0个过滤条件
)

//添加信息到queryInfos
func AddQueryEsInfos(queryInfos *[]model.QueryInfo) {

	queryByLostA2Avg := model.QueryInfo{
		HostName:    conf.Conf.EsConfig.CdnPullOpenHost,
		Group:       []string{"isp","cv"},
		Filter:      map[string]string{},
		Target:      "open_success_t",
		Method:      "avg",
		Minute:      1,
		GroupFilter: GF20,
	}
	startTime, endTime := utils.StartAndEndTime(queryByLostA2Avg.Minute)
	queryByLostA2Avg.StartTime = startTime
	queryByLostA2Avg.EndTime = endTime

	queryByLostA2DurationMs := model.QueryInfo{
		HostName:conf.Conf.EsConfig.CdnPullOpenHost,
		Group: []string{"isp","cv"},
		Filter: map[string]string{},
		Target: "info_duration_ms",
		Method: "avg",
		Minute: 1,
		GroupFilter:GF20,
	}

	*queryInfos = append(*queryInfos,queryByLostA2Avg,queryByLostA2DurationMs)
}


/*

适用于两个过滤条件+求avg，max等条件情况
例如：
queryInfo := HostName:conf.Conf.EsConfig.CdnPullOpenHost,
			Group1: "isp",
			Group2: "cv",
			Target: "open_success_t",
			Method: "avg",
			Minute: 1,
ExeQueryEs(queryInfo model.QueryInfo) = SQL: select avg(open_success_t) from tbl between 前(3+1)分钟 and 前3分钟 group by app,cv;
*/
func ExeQueryEs(queryInfo model.QueryInfo) string {

	switch queryInfo.GroupFilter {
	case GF20:
		GF10PostBody := OneGroupZeroFilter(queryInfo)
		GF20PostBody := TwoGroupZeroFilter(queryInfo)

		var GF20MysqlInfos []model.GF20MysqlInfo
		var EsResultGF10 model.QueryRes
		var EsResultGF20 model.QueryRes

		SendEsPost(GF10PostBody,&EsResultGF10,queryInfo.HostName)
		SendEsPost(GF20PostBody,&EsResultGF20,queryInfo.HostName)

		dao.CollectEsRes2Mysql(&GF20MysqlInfos,queryInfo, EsResultGF10, EsResultGF20)
		//插入数据库


	}

	return ""
}



func OneGroupZeroFilter(queryInfo model.QueryInfo) string {
	if len(queryInfo.Group)<1 {
		log.Errorf("TwoGroupZeroFilter error Group Num: %d", len(queryInfo.Group))
		return ""
	}
	postBody := GetBodyByGroupFilter10(queryInfo)
	return postBody
}

func TwoGroupZeroFilter(queryInfo model.QueryInfo) string {
	if len(queryInfo.Group)<2 {
		log.Errorf("TwoGroupZeroFilter error Group Num: %d", len(queryInfo.Group))
		return ""
	}

	postBody := GetBodyByGroupFilter20(queryInfo.Group[0], queryInfo.Group[1], queryInfo.Target, queryInfo.Method, queryInfo.Minute)
	return postBody
}


//遍历queryInfos，对每个queryInfo进行查询ES操作
func StartQueryES(queryInfos []model.QueryInfo) {
	for _, queryInfo := range queryInfos {
		ExeQueryEs(queryInfo)
	}
}


