package conf

import (
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/utils"
)

func AddOriginalConfs(originalConf *model.OriginalConf) {
	*originalConf = model.OriginalConf{
		HostName: "http://c2-media-mix02.bj:9200/new_cdn_pull_open_detail_online/_search?pretty",
		Minute:   1,
		TblName:  "media_realTime_capturefps",
	}
	AddTime2OriginalConf(originalConf)

}

func AddTime2OriginalConf(originalConf *model.OriginalConf) {
	//startTime, endTime := utils.StartAndEndTime(originalConf.Minute)
	//originalConf.StartTime = startTime
	originalConf.StartTime = "2020-11-12 11:18:14"
	//originalConf.EndTime = endTime
	originalConf.EndTime = "2020-11-12 11:20:14"


}

func AddHandleConfs(handleInfos *[]model.HandleConf) {
	handleInfoGF10 := model.HandleConf{
		HostName: Conf.EsHost.PullHost,
		Group:    []string{"isp"},
		Targets: []model.Target{
			{
				Method: "avg",
				Tar:    "open_success_t",
			},
			{
				Method: "max",
				Tar:    "info_duration_ms",
			}},
		Minute:      1,
		GroupFilter: model.GF10,
		TblName:     "media_realTime_capturefps",
	}
	AddTimeAndTargetNum2HandleConf(&handleInfoGF10)

	*handleInfos = append(*handleInfos, handleInfoGF10)
}

func AddTimeAndTargetNum2HandleConf(hadleInfo *model.HandleConf) {
	startTime, endTime := utils.StartAndEndTime(hadleInfo.Minute)
	hadleInfo.StartTime = startTime
	hadleInfo.EndTime = endTime
	hadleInfo.TargetNum = len(hadleInfo.Targets)
}

//添加信息到queryInfos
func AddUntreatedConfs(queryInfos *[]model.UntreatedConf) {

	//拉流部分实时统计
	/*pullGeneralConf := model.UntreatedConf{
		HostName: Conf.EsHost.PullHost, //字段含义备注见结构体
		Group:    []string{"md_einfo_PlayerLog.pull_type", "app", "md_einfo_PlayerLog.live_type", "cv", "md_einfo_PlayerLog.domain"},
		Filters: []model.Filter{
			{
				Field: "record_time",
				Gte:   0,
				Lte:   0,
			},
			{
				Field: "md_einfo_PlayerLog.custom.krnsPlay.bitrateA",
				Gte:   0,
				Lte:   512,
			},
			{
				Field: "md_einfo_PlayerLog.custom.krnsPlay.bitrateV",
				Gte:   0,
				Lte:   20480,
			},
			{
				Field: "md_einfo_PlayerLog.custom.krnsPlay.medFps",
				Gte:   0,
				Lte:   60,
			},
			{
				Field: "md_einfo_PlayerLog.lostA_avg",
				Gte:   0,
				Lte:   100,
			},
			{
				Field: "md_einfo_PlayerLog.lostV_avg",
				Gte:   0,
				Lte:   100,
			},
			{
				Field: "md_einfo_PlayerLog.rtt_avg",
				Gte:   0,
				Lte:   4999,
			},
		},
		Targets: []model.Target{
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.custom.krnsPlay.bitrateA.keyword",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.custom.krnsPlay.bitrateV.keyword",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.custom.krnsPlay.medFps.keyword",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.lostA_avg",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.lostV_avg",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_PlayerLog.rtt_avg",
			},
		},
		Minute:      1,
		GroupFilter: model.GF50,
		TblName:     "media_realTime_capturefps",
	}
	AddTime2UntreatedConf(&pullGeneralConf)*/

	//推流部分实时统计
	/*pushGeneralConf := model.UntreatedConf{
		HostName: Conf.EsHost.PushHost, //字段含义备注见结构体
		Group:    []string{"md_einfo_value_push_type", "app", "md_einfo_value_live_type", "cv", "md_einfo_value_domain"},
		Filters: []model.Filter{
			{
				Field: "record_time",
				Gte:   0,
				Lte:   0,
			},
			{
				Field: "md_einfo_value_audio_encode_bitrate",
				Gte:   0,
				Lte:   256,
			},
			{
				Field: "md_einfo_value_video_encode_bitrate",
				Gte:   0,
				Lte:   10240,
			},
			{
				Field: "md_einfo_value_capturefps",
				Gte:   0,
				Lte:   60,
			},
			{
				Field: "md_einfo_value_lostA_avg",
				Gte:   0,
				Lte:   100,
			},
			{
				Field: "md_einfo_value_lostV_avg",
				Gte:   0,
				Lte:   4999,
			},
			{
				Field: "md_einfo_value_custom_krnsPub.bitrateA",
				Gte:   0,
				Lte:   512,
			},
			{
				Field: "md_einfo_value_custom_krnsPub.bitrateV",
				Gte:   0,
				Lte:   20480,
			},
		},
		Targets: []model.Target{
			{
				Method: "avg",
				Tar:    "md_einfo_value_audio_encode_bitrate",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_video_encode_bitrate",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_capturefps",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_lostA_avg",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_lostV_avg",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_rtt_avg",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_custom_krnsPub.bitrateA",
			},
			{
				Method: "avg",
				Tar:    "md_einfo_value_custom_krnsPub.bitrateV",
			},
		},
		Minute:      1,
		GroupFilter: model.GF50,
		TblName:     "media_realTime_capturefps",
	}
	AddTime2UntreatedConf(&pushGeneralConf)*/

	/*pushVideoEncoderConf := model.UntreatedConf{  //区分h264还是265
		HostName: Conf.EsHost.PushHost, //字段含义备注见结构体
		Group:    []string{"md_einfo_value_push_type", "app", "md_einfo_value_live_type", "cv", "md_einfo_value_domain","md_einfo_value_video_encoder"},
		Filters: []model.Filter{
			{
				Field: "record_time",
				Gte:   0,
				Lte:   0,
			},
			{
				Field: "md_einfo_value_audio_encode_bitrate",
				Gte:   0,
				Lte:   256,
			},
			{
				Field: "md_einfo_value_video_encode_bitrate",
				Gte:   0,
				Lte:   10240,
			},
			{
				Field: "md_einfo_value_capturefps",
				Gte:   0,
				Lte:   60,
			},
			{
				Field: "md_einfo_value_lostA_avg",
				Gte:   0,
				Lte:   100,
			},
			{
				Field: "md_einfo_value_lostV_avg",
				Gte:   0,
				Lte:   4999,
			},
			{
				Field: "md_einfo_value_custom_krnsPub.bitrateA",
				Gte:   0,
				Lte:   512,
			},
			{
				Field: "md_einfo_value_custom_krnsPub.bitrateV",
				Gte:   0,
				Lte:   20480,
			},
		},
		Minute:      1,
		GroupFilter: model.GF50,
		TblName:     "media_realTime_capturefps",
	}
	AddTime2UntreatedConf(&pushVideoEncoderConf)*/

	succAndSecConf := model.UntreatedConf{  //统计秒开率，成功率
		HostName: Conf.EsHost.OpenHost, //字段含义备注见结构体
		Group:    []string{"md_einfo_pull_type", "app", "md_einfo_live_type", "cv", "md_einfo_domain"},
		Filters: []model.Filter{
			{
				Field: "record_time",
				Gte:   0,
				Lte:   0,
			},
		},
		Minute:      1,
		GroupFilter: model.GF50,
		DistinctiveMark:model.Mark1,
		TblName:     "media_realTime_capturefps",
	}
	AddTime2UntreatedConf(&succAndSecConf)

	//*queryInfos = append(*queryInfos, pullGeneralConf, pushGeneralConf)
	*queryInfos = append(*queryInfos,succAndSecConf)
}

func AddTime2UntreatedConf(queryInfo *model.UntreatedConf) {
	startTime, endTime := utils.StartAndEndTimestamp(queryInfo.Minute)
	queryInfo.Filters[0].Gte = startTime
	queryInfo.Filters[0].Lte = endTime
}
