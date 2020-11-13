package service

import (
	"bytes"
	log "code.inke.cn/BackendPlatform/golang/logging"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendEsPost(postBody string, object interface{}, hostName string) {
	var jsonStr = []byte(postBody)
	req, err := http.NewRequest("POST", hostName, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &object)
	//fmt.Println(string(body))
}

func PostJoint(utConf model.UntreatedConf, level int) string {
	var (
		query string
		aggs  string
	)

	QueryJointRes(utConf, &query)
	AggsJointRes(utConf, &aggs, level)

	return JointRes(query, aggs)

}

func AggsJointRes(utConf model.UntreatedConf, aggs *string, level int) {

	targetJoint := TargetsJoint(utConf.Targets)
	if len(utConf.Targets) == 0 { // 用于处理不需要求值的值（无targets，单纯求个数），
		targetJoint = fmt.Sprintf(`"aggs": {
	                        "group_by_group": {
	                          "terms": {
	                            "field": "%s"
	                          }
	                        }
	                      }`, utConf.Group[len(utConf.Group)-1])
		//level -= 1
	}
	if utConf.DistinctiveMark == model.Mark1 { //特殊指标直接赋值
		targetJoint = GetSuccAndSecTargets()
	}

	//fmt.Println(targetJoint)
	*aggs = AggsJoint(utConf.Group, targetJoint, level)
}

//成功秒开指标只有一个，没必要特殊处理，以后可以优化
func GetSuccAndSecTargets() string {
	return `"aggs": {
                       "group_by_group": {
                         "terms": {
                           "field": "md_einfo_detail_time_open_status"
                         },
                         "aggs": {
                           "all_time_1000": {
                             "filter": {
                               "range": {
                                 "md_einfo_push_detail_time_all_time": {
                                   "lte": 1000
                                 }
                               }
                             }
                           },
                           "all_time_5000":{
                             "filter": {
                               "range": {
                                 "md_einfo_push_detail_time_all_time": {
                                   "lte": 5000
                                 }
                               }
                             }
                           }
                         }
                       }
                     }`
}

func AggsJoint(groups []string, targetJoint string, level int) string {
	var tempAggs string
	switch len(groups) {
	case 0:
		log.Errorf("AggsJoint no group! targets:%v", groups)
		return tempAggs
	default:
		tempAggs = AddAggsJoint(groups, targetJoint, level)

	}
	return tempAggs
}

func AddAggsJoint(groups []string, targetJoint string, level int) string {
	var tempGroups string
	for i := level - 1; i >= 0; i-- {
		if i == level-1 {
			tempGroups = fmt.Sprintf(`"aggs": {
        "group_by_group": {
          "terms": {
            "field": "%s"
          },
          %s
        }
      }`, groups[i], targetJoint)
			continue
		}

		tempGroups = fmt.Sprintf(`"aggs": {
        "group_by_group": {
          "terms": {
            "field": "%s"
          },
          %s
        }
      }`, groups[i], tempGroups)
	}

	//fmt.Println(tempGroups)
	return tempGroups
}

func TargetsJoint(targets []model.Target) string {
	var tempTarges string
	switch len(targets) {
	case 0:
		log.Errorf("TargetsJoint no target! targets:%v", targets)
		return tempTarges
	/*case 1:
			tempTarges = fmt.Sprintf(`"oper_target1": {
	                          "%s": {
	                            "field": "%s"
	                          }
	                        }`, targets[0].Method, targets[0].Tar)*/
	default:
		tempTarges = AddTargetsJoint(targets)
	}
	tempTarges = fmt.Sprintf(`"aggs": {%s
}`, tempTarges)
	return tempTarges
}

func AddTargetsJoint(targets []model.Target) string {
	var tempTarget string
	//生成求值目标对象
	for i, target := range targets {
		targetNum := i + 1
		if i == len(targets)-1 || len(targets) == 1 {
			tempTarget = tempTarget + fmt.Sprintf(`"oper_target%d": {
                          "%s": {
                            "field": "%s"
                          }
                        }`, targetNum, target.Method, target.Tar)
			break
		}

		tempTarget = tempTarget + fmt.Sprintf(`"oper_target%d": {
                          "%s": {
                            "field": "%s"
                          }
                        },`, targetNum, target.Method, target.Tar)
	}
	return tempTarget
}

func JointRes(query string, aggs string) string {
	return fmt.Sprintf(`{
		%s
		%s
	}`, query, aggs)
}

//拼接Query
func QueryJointRes(utConf model.UntreatedConf, query *string) {
	switch len(utConf.Filters) {
	case 0:
		log.Errorf("QueryJointRes no Filters! Filters:%v", utConf.Filters)
	case 1:
		*query = fmt.Sprintf(`"size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "%s": {
              "gte": "%d",
              "lte": "%d"
            }
          }
        }
      ]
    } 
  },`,utConf.Filters[0].Field, utConf.Filters[0].Gte, utConf.Filters[0].Lte)
	default:
		*query = AddQueryJoint(utConf)
	}
}

func AddQueryJoint(utConf model.UntreatedConf) string {
	var tempQuery string
	for i, filter := range utConf.Filters {
		if i == len(utConf.Filters)-1 {
			tempQuery = tempQuery + fmt.Sprintf(`
        {
          "range": {
            "%s": {
              "gte": %d,
              "lte": %d
            }
          }
        }`, filter.Field, filter.Gte, filter.Lte)
			break
		}

		tempQuery = tempQuery + fmt.Sprintf(`
        {
          "range": {
            "%s": {
              "gte": %d,
              "lte": %d
            }
          }
        },`, filter.Field, filter.Gte, filter.Lte)
	}
	tempQuery = fmt.Sprintf(`"size": 0,
  "query": {
    "bool": {
      "must": [
        %s
      ]
    } 
  },`, tempQuery)
	return tempQuery
}

func GetBodyByOriginal(handleInfo model.OriginalConf) string {
	postBody := fmt.Sprintf(`{
 "size": 15,
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "times": {
              "gte": "%s",
              "lte": "%s"
            }
          }
        }
      ]
    } 
  }
}`, handleInfo.StartTime, handleInfo.EndTime)
	return postBody
}

func GetBodyByHandleGF10T2(handleInfo model.HandleConf) string {
	postBody := fmt.Sprintf(`{
  "size": 0, 
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "times": {
              "gte": "%s",
              "lte": "%s"
            }
          }
        }
      ]
    } 
  },
  "aggs": {
    "group_by_group": {
      "terms": {
        "field": "%s"
      },
      "aggs": {
        "field1": {
          "%s": {
            "field": "%s"
          }
        },
        "field2":{
          "%s": {
            "field": "%s"
          }
        }
      }
    }
  }
}`, handleInfo.StartTime, handleInfo.EndTime, handleInfo.Group[0],
		handleInfo.Targets[0].Method, handleInfo.Targets[0].Tar, handleInfo.Targets[1].Method, handleInfo.Targets[1].Tar)
	return postBody
}
