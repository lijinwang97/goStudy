package dao

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/conf"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/utils"
	"context"
	"time"
)

// Dao represents data access object
type Dao struct {
	c *conf.Config
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c: c,
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

func CollectEsRes2Mysql(MysqlGF20Infos *[]model.GF20MysqlInfo,queryInfo model.QueryInfo, EsResultGF10, EsResultGF20 model.QueryRes) {



	var MysqlGF20Info model.GF20MysqlInfo

	MysqlGF20Info.Ymd = time.Now().AddDate(0, 0, -0).Format("2006-01-02")
	MysqlGF20Info.Event_Time = utils.String2TimeStampe(queryInfo.StartTime)

	MysqlGF20Info.All_Total = EsResultGF10.Hits.Total
	for _, bucket := range EsResultGF10.Aggregations.Group1.Buckets {
		MysqlGF20Info.AppName = bucket.Key
		MysqlGF20Info.Cv = "all"
		MysqlGF20Info.Sec_Total = bucket.Doc_count
		MysqlGF20Info.Res_Total = bucket.Doc_count
		MysqlGF20Info.Value = bucket.Oper_target.Value
		*MysqlGF20Infos = append(*MysqlGF20Infos, MysqlGF20Info)
	}

	MysqlGF20Info.All_Total = EsResultGF20.Hits.Total
	for _, bucket := range EsResultGF20.Aggregations.Group1.Buckets {
		MysqlGF20Info.AppName = bucket.Key
		MysqlGF20Info.Sec_Total = bucket.Doc_count
		for _, coreBucket := range bucket.Group2.CoreBuckets {
			MysqlGF20Info.Cv = coreBucket.CoreKey
			MysqlGF20Info.Res_Total = coreBucket.CoreCount
			MysqlGF20Info.Value = coreBucket.CoreOperTarget.Value
			*MysqlGF20Infos = append(*MysqlGF20Infos, MysqlGF20Info)
		}

		MysqlGF20Info.Cv = "all"
		MysqlGF20Info.Sec_Total = bucket.Doc_count
		MysqlGF20Info.Res_Total = bucket.Doc_count
		MysqlGF20Info.Value = bucket.Oper_target.Value
		*MysqlGF20Infos = append(*MysqlGF20Infos, MysqlGF20Info)
	}
}




