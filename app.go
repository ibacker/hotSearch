package main

import (
	"encoding/json"
	"fmt"
	"hotSearch/config"
	"hotSearch/model"
	"hotSearch/sql"
	"io/ioutil"
	"net/http"
)

func main() {
	sql.TestConnection()

	sql.InsertZhiHuHotList(processZHHotList(getZhiHuHot()))
	sql.InsertWeiboHotList(processWBHotSearch(getWBHotSearch()))
	sql.InsertBaiDuHotList(processBDHotList(getBaiduHotSearch()))
}

func processZHHotList(result *model.HotList) []model.ZhiHuQuestion {
	var questionList []model.ZhiHuQuestion
	for _, v := range result.Data {
		var question = model.ZhiHuQuestion{
			Title:    v.Target.Title,
			Url:      fmt.Sprintf("https://www.zhihu.com/question/%v", v.Target.Id),
			Desc:     v.Target.Desc,
			HotTag:   v.CardLabel.Type,
			HotScore: v.DetailText,
		}
		questionList = append(questionList, question)
	}
	return questionList
}

func getZhiHuHot() *model.HotList {
	url := config.Conf.Get("url.zhihu").(string)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var list model.HotList
		if err := json.Unmarshal(body, &list); err == nil {
			return &list
		}
	}
	return nil
}

/*
*
获取百度热搜内容
*/
func getBaiduHotSearch() *model.BDHotSearch {
	url := config.Conf.Get("url.baidu").(string)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var bdHotSearch model.BDHotSearch
		if err := json.Unmarshal(body, &bdHotSearch); err == nil {
			return &bdHotSearch
		}
	}
	return nil
}

/*
*
处理百度热搜内容
*/
func processBDHotList(result *model.BDHotSearch) []model.BDHotSearchItem {
	var hotSearchList []model.BDHotSearchItem

	var cards = result.BDDate.BDCards[0]
	for _, v := range cards.BDHotSearchContent {
		fmt.Println(v)
		hotSearchList = append(hotSearchList, v)
	}

	return hotSearchList
}

func getWBHotSearch() *model.WBHotSearch {
	url := config.Conf.Get("url.sina").(string)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var wbHotSearch model.WBHotSearch
		if err := json.Unmarshal(body, &wbHotSearch); err == nil {
			return &wbHotSearch
		}
	}
	return nil
}

func processWBHotSearch(result *model.WBHotSearch) []model.WBHotBandItem {
	var hotSearchList []model.WBHotBandItem

	var bandList = result.WBDate

	for _, v := range bandList.WBBandList {
		v.Url = "https://s.weibo.com/weibo?q=%23" + v.Word + "%23"
		fmt.Println(v)
		hotSearchList = append(hotSearchList, v)
	}
	return hotSearchList
}
