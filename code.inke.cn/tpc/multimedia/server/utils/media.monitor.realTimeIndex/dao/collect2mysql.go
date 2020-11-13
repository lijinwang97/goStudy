package dao

import (
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/utils"
	"time"
)

func CollectEsOriginalRes(mysqlInfos *[]model.MysqOriginalInfo, originalConf model.OriginalConf, esResult model.OriginalJsonRes) {
	var tempMysqlInfo model.MysqOriginalInfo

	tempMysqlInfo.Ymd = time.Now().AddDate(0, 0, -0).Format("2006-01-02")
	tempMysqlInfo.Event_Time = utils.String2TimeStampe(originalConf.StartTime)
	tempMysqlInfo.All = esResult.Hits.Total
	for _, hits := range esResult.Hits.Hits2 {
		tempMysqlInfo.Country = hits.Source.Country
		tempMysqlInfo.Province = hits.Source.Province
		tempMysqlInfo.City = hits.Source.City
		tempMysqlInfo.Isp = hits.Source.Isp
		tempMysqlInfo.Info_duration_ms = hits.Source.Info_duration_ms
		*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
	}
}

func CollectEsHandleRes(mysqlInfos *[]model.MysqHandleInfo, handleInfo model.HandleConf, esResult model.HandleJsonRes) {
	var tempMysqlInfo model.MysqHandleInfo

	tempMysqlInfo.Ymd = time.Now().AddDate(0, 0, -0).Format("2006-01-02")
	tempMysqlInfo.Event_Time = utils.String2TimeStampe(handleInfo.StartTime)
	tempMysqlInfo.All = esResult.Hits.Total

	for _, bucket := range esResult.Aggregations.Group1.Buckets {
		tempMysqlInfo.KeyG1 = bucket.KeyG1
		tempMysqlInfo.DocCountG1 = bucket.DocCountG1
		tempMysqlInfo.Field1 = bucket.Field1.Value
		tempMysqlInfo.Field2 = bucket.Field2.Value
		*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
	}
}

func CollectEsRes2Mysql(mysqlInfos *[]model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, level int) {
	if len(utConf.Targets) == 0 && utConf.DistinctiveMark!=model.Mark1 {
		level += 1
	}

	var tempMysqlInfo model.MysqUntreatedInfo
	tempMysqlInfo.Ymd = time.Now().AddDate(0, 0, -0).Format("2006-01-02")
	tempMysqlInfo.Event_Time = utConf.Filters[0].Gte
	tempMysqlInfo.All = esResult.Hits.Total

	switch level {
	case model.Level1:
		AssignmentLevel1(tempMysqlInfo,utConf,esResult,mysqlInfos)
	case model.Level2:
		AssignmentLevel2(tempMysqlInfo,utConf,esResult,mysqlInfos)
	case model.Level3:
		AssignmentLevel3(tempMysqlInfo,utConf,esResult,mysqlInfos)
	case model.Level4:
		AssignmentLevel4(tempMysqlInfo,utConf,esResult,mysqlInfos)
	case model.Level5:
		AssignmentLevel5(tempMysqlInfo,utConf,esResult,mysqlInfos)
	case model.Level6:
		AssignmentLevel6(tempMysqlInfo,utConf,esResult,mysqlInfos)
	default:
		log.Errorf("CollectEsRes2Mysql the error level,can not parse Json! level:%s", level)
	}
}

func AssignmentLevel1(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		tempMysqlInfo.Group2 = "all"
		tempMysqlInfo.Group2_Num = bucketG1.CountG1
		tempMysqlInfo.Group3 = "all"
		tempMysqlInfo.Group3_Num = bucketG1.CountG1
		tempMysqlInfo.Group4 = "all"
		tempMysqlInfo.Group4_Num = bucketG1.CountG1
		tempMysqlInfo.Group5 = "all"
		tempMysqlInfo.Group5_Num = bucketG1.CountG1
		tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG1.OperTarget1G1.Value)
		tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG1.OperTarget2G1.Value)
		tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG1.OperTarget3G1.Value)
		tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG1.OperTarget4G1.Value)
		tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG1.OperTarget5G1.Value)
		tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG1.OperTarget6G1.Value)
		tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG1.OperTarget7G1.Value)
		tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG1.OperTarget8G1.Value)
		tempMysqlInfo.Remark = ""
		*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
	}
}

func AssignmentLevel2(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		for _, bucketG2 := range bucketG1.Group2.BucketsG2 {
			tempMysqlInfo.Group2 = bucketG2.KeyG2
			tempMysqlInfo.Group2_Num = bucketG2.CountG2
			tempMysqlInfo.Group3 = "all"
			tempMysqlInfo.Group3_Num = bucketG2.CountG2
			tempMysqlInfo.Group4 = "all"
			tempMysqlInfo.Group4_Num = bucketG2.CountG2
			tempMysqlInfo.Group5 = "all"
			tempMysqlInfo.Group5_Num = bucketG2.CountG2
			tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG2.OperTarget1G2.Value)
			tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG2.OperTarget2G2.Value)
			tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG2.OperTarget3G2.Value)
			tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG2.OperTarget4G2.Value)
			tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG2.OperTarget5G2.Value)
			tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG2.OperTarget6G2.Value)
			tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG2.OperTarget7G2.Value)
			tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG2.OperTarget8G2.Value)
			tempMysqlInfo.Remark = ""
			if len(utConf.Targets) == 0 { //为了计算比例
				if bucketG1.CountG1 == 0 {  //除数不能为0
					*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
					continue
				}
				tempMysqlInfo.Value1 = float64(bucketG2.CountG2) / float64(bucketG1.CountG1)
			}
			*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
		}
	}
}

func AssignmentLevel3(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		for _, bucketG2 := range bucketG1.Group2.BucketsG2 {
			tempMysqlInfo.Group2 = bucketG2.KeyG2
			tempMysqlInfo.Group2_Num = bucketG2.CountG2
			for _, bucketG3 := range bucketG2.Group3.BucketsG3 {
				tempMysqlInfo.Group3 = bucketG3.KeyG3
				tempMysqlInfo.Group3_Num = bucketG3.CountG3
				tempMysqlInfo.Group4 = "all"
				tempMysqlInfo.Group4_Num = bucketG3.CountG3
				tempMysqlInfo.Group5 = "all"
				tempMysqlInfo.Group5_Num = bucketG3.CountG3
				tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG3.OperTarget1G3.Value)
				tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG3.OperTarget2G3.Value)
				tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG3.OperTarget3G3.Value)
				tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG3.OperTarget4G3.Value)
				tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG3.OperTarget5G3.Value)
				tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG3.OperTarget6G3.Value)
				tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG3.OperTarget7G3.Value)
				tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG3.OperTarget8G3.Value)
				tempMysqlInfo.Remark = ""
				if len(utConf.Targets) == 0 {
					if bucketG2.CountG2 == 0 {  //除数不能为0
						*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
						continue
					}
					tempMysqlInfo.Value1 = float64(bucketG3.CountG3) / float64(bucketG2.CountG2)
				}
				tempMysqlInfo.Remark = ""
				*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
			}
		}
	}
}

func AssignmentLevel4(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		for _, bucketG2 := range bucketG1.Group2.BucketsG2 {
			tempMysqlInfo.Group2 = bucketG2.KeyG2
			tempMysqlInfo.Group2_Num = bucketG2.CountG2
			for _, bucketG3 := range bucketG2.Group3.BucketsG3 {
				tempMysqlInfo.Group3 = bucketG3.KeyG3
				tempMysqlInfo.Group3_Num = bucketG3.CountG3
				for _, bucketG4 := range bucketG3.Group4.BucketsG4 {
					tempMysqlInfo.Group4 = bucketG4.KeyG4
					tempMysqlInfo.Group4_Num = bucketG4.CountG4
					tempMysqlInfo.Group5 = "all"
					tempMysqlInfo.Group5_Num = bucketG4.CountG4
					tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG4.OperTarget1G4.Value)
					tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG4.OperTarget2G4.Value)
					tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG4.OperTarget3G4.Value)
					tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG4.OperTarget4G4.Value)
					tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG4.OperTarget5G4.Value)
					tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG4.OperTarget6G4.Value)
					tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG4.OperTarget7G4.Value)
					tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG4.OperTarget8G4.Value)
					tempMysqlInfo.Remark = ""
					if len(utConf.Targets) == 0 {
						if bucketG3.CountG3 == 0 {  //除数不能为0
							*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
							continue
						}
						tempMysqlInfo.Value1 = float64(bucketG4.CountG4) / float64(bucketG3.CountG3)
					}
					*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
				}
			}
		}
	}
}

func AssignmentLevel5(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		for _, bucketG2 := range bucketG1.Group2.BucketsG2 {
			tempMysqlInfo.Group2 = bucketG2.KeyG2
			tempMysqlInfo.Group2_Num = bucketG2.CountG2
			for _, bucketG3 := range bucketG2.Group3.BucketsG3 {
				tempMysqlInfo.Group3 = bucketG3.KeyG3
				tempMysqlInfo.Group3_Num = bucketG3.CountG3
				for _, bucketG4 := range bucketG3.Group4.BucketsG4 {
					tempMysqlInfo.Group4 = bucketG4.KeyG4
					tempMysqlInfo.Group4_Num = bucketG4.CountG4
					for _, bucketG5 := range bucketG4.Group5.BucketsG5 {
						tempMysqlInfo.Group5 = bucketG5.KeyG5
						tempMysqlInfo.Group5_Num = bucketG5.CountG5
						tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG5.OperTarget1G5.Value)
						tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG5.OperTarget2G5.Value)
						tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG5.OperTarget3G5.Value)
						tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG5.OperTarget4G5.Value)
						tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG5.OperTarget5G5.Value)
						tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG5.OperTarget6G5.Value)
						tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG5.OperTarget7G5.Value)
						tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG5.OperTarget8G5.Value)
						tempMysqlInfo.Remark = ""
						if len(utConf.Targets) == 0 {
							if bucketG4.CountG4 == 0 {  //除数不能为0
								*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
								continue
							}
							tempMysqlInfo.Value1 = float64(bucketG5.CountG5) / float64(bucketG4.CountG4)
						}
						*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
					}
				}
			}
		}
	}
}

func AssignmentLevel6(tempMysqlInfo model.MysqUntreatedInfo, utConf model.UntreatedConf, esResult *model.UntreatedJsonRes, mysqlInfos *[]model.MysqUntreatedInfo) {
	for _, bucketG1 := range esResult.Aggregations.Group1.BucketsG1 {
		tempMysqlInfo.Group1 = utils.Int2String(bucketG1.KeyG1)
		tempMysqlInfo.Group1_Num = bucketG1.CountG1
		for _, bucketG2 := range bucketG1.Group2.BucketsG2 {
			tempMysqlInfo.Group2 = bucketG2.KeyG2
			tempMysqlInfo.Group2_Num = bucketG2.CountG2
			for _, bucketG3 := range bucketG2.Group3.BucketsG3 {
				tempMysqlInfo.Group3 = bucketG3.KeyG3
				tempMysqlInfo.Group3_Num = bucketG3.CountG3
				for _, bucketG4 := range bucketG3.Group4.BucketsG4 {
					tempMysqlInfo.Group4 = bucketG4.KeyG4
					tempMysqlInfo.Group4_Num = bucketG4.CountG4
					for _, bucketG5 := range bucketG4.Group5.BucketsG5 {
						tempMysqlInfo.Group5 = bucketG5.KeyG5
						tempMysqlInfo.Group5_Num = bucketG5.CountG5
						for _, bucketG6 := range bucketG5.Group6.BucketsG6 {
							tempMysqlInfo.Group6 = bucketG6.KeyG6
							tempMysqlInfo.Group6_Num = bucketG6.CountG6
							tempMysqlInfo.Value1 = utils.FloatRetain2(bucketG6.OperTarget1G6.Value)
							tempMysqlInfo.Value2 = utils.FloatRetain2(bucketG6.OperTarget2G6.Value)
							tempMysqlInfo.Value3 = utils.FloatRetain2(bucketG6.OperTarget3G6.Value)
							tempMysqlInfo.Value4 = utils.FloatRetain2(bucketG6.OperTarget4G6.Value)
							tempMysqlInfo.Value5 = utils.FloatRetain2(bucketG6.OperTarget5G6.Value)
							tempMysqlInfo.Value6 = utils.FloatRetain2(bucketG6.OperTarget6G6.Value)
							tempMysqlInfo.Value7 = utils.FloatRetain2(bucketG6.OperTarget7G6.Value)
							tempMysqlInfo.Value8 = utils.FloatRetain2(bucketG6.OperTarget8G6.Value)
							tempMysqlInfo.Remark = ""
							if len(utConf.Targets) == 0 {
								if bucketG5.CountG5 == 0 {  //除数不能为0
									*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
									continue
								}
								tempMysqlInfo.Value1 = float64(bucketG6.CountG6) / float64(bucketG5.CountG5)
							}
							*mysqlInfos = append(*mysqlInfos, tempMysqlInfo)
						}
					}
				}
			}
		}
	}
}