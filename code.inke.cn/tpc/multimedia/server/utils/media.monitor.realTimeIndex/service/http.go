package service

import (
	"bytes"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/model"
	"code.inke.cn/tpc/multimedia/server/utils/media.monitor.realTimeIndex/utils"
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
	json.Unmarshal([]byte(body),&object)
}

func GetBodyByGroupFilter10(queryInfo model.QueryInfo) string {
	postBody := fmt.Sprintf(`{
  "size": 0,
  "query": {
    "constant_score": {
      "filter": {
        "range": {
          "times": {
            "gte": "%s",
            "lte": "%s"
          }
        }
      }
    }
  },
  "aggs": {
    "group_by_group1": {
      "terms": {
        "field": "%s"
      },
      "aggs": {
        "oper_target": {
          "%s": {
            "field": "%s"
          }
        }
      }
    }
  }
}`, queryInfo.StartTime, queryInfo.EndTime,queryInfo.Group[0],queryInfo.Method,queryInfo.Target)
	fmt.Println(postBody)
	return postBody
}

func GetBodyByGroupFilter20(group1 string, group2 string, target string, method string, minute int) string {
	startTime, endTime := utils.StartAndEndTime(minute)
	postBody := fmt.Sprintf(`{
    "size":0,
    "query":{
        "constant_score":{
            "filter":{
                "range":{
                    "times":{
                        "gte":"%s",
                        "lte":"%s"
                    }
                }
            }
        }
    },
    "aggs":{
        "group_by_group1":{
            "terms":{
                "field":"%s"
            },
            "aggs":{
                "group_by_group2":{
                    "terms":{
                        "field":"%s"
                    },
                    "aggs":{
                        "oper_target":{
                            "%s":{
                                "field":"%s"
                            }
                        }
                    }
                }
            }
        }
    }
}`, startTime, endTime,group1,group2,method,target)
	fmt.Println(postBody)
	return postBody
}