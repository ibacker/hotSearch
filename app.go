package main

import (
	"encoding/json"
	"fmt"
	"hotSearch/model"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var fileName = time.Now().Format("2006-1-2")

func main() {
	result := getZhiHuHot()
	question := getQuestion(result)
	//createRaw(question, fileName)
	//utils.CreateReadMe(question)
	//utils.CreateArchives(question, fileName)
	for i := range question {
		fmt.Println(question[i])
	}
	getBaiduHotList(getBaiduHotSearch())
	processWBHotSearch(getWBHotSearch())
}

func createRaw(data []model.Question, fileName string) {
	filePath := fmt.Sprintf("./raw/%v.json", fileName)
	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		bytes, _ := json.Marshal(data)
		file.Write(bytes)
	}
}

func getQuestion(result *model.HotList) []model.Question {
	var questionList []model.Question
	for _, v := range result.Data {
		question := model.Question{
			Title: v.Target.Title,
			Url:   fmt.Sprintf("https://www.zhihu.com/question/%v", v.Target.Id),
		}
		questionList = append(questionList, question)
	}
	return questionList
}

func getZhiHuHot() *model.HotList {
	resp, _ := http.Get("https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=100")
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
	resp, _ := http.Get("https://top.baidu.com/api/board?platform=wise&tab=realtime")
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
func getBaiduHotList(result *model.BDHotSearch) []model.BDHotSearchItem {
	var hotSearchList []model.BDHotSearchItem

	var cards = result.BDDate.BDCards[0]
	for _, v := range cards.BDHotSearchContent {
		fmt.Println(v)
		hotSearchList = append(hotSearchList, v)
	}

	return hotSearchList
}

/*
获取微博热搜榜
*/
func getWBHotSearch() *model.WBHotSearch {
	resp, _ := http.Get("https://weibo.com/ajax/statuses/hot_band")
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

//https://weibo.com/ajax/statuses/hot_band/
//https://m.weibo.cn/api/container/getIndex?containerid=106003type%3D25%26t%3D3%26disable_hot%3D1%26filter_type%3Drealtimehot
