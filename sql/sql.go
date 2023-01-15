package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hotSearch/config"
	"hotSearch/model"
	"log"
	"strconv"
	"time"
)

func TestConnection() {

	address := config.Conf.Get("sql.address").(string)
	user := config.Conf.Get("sql.user").(string)
	passwd := config.Conf.Get("sql.passwd").(string)
	driver := config.Conf.Get("sql.driver").(string)
	_SQLConnection(driver, user, passwd, address)
}

func _SQLConnection(driver, user, passwd, address string) {
	db, err := sql.Open(driver, user+":"+passwd+"@tcp("+address+")/mysql")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}

func _GetDBConnect() *sql.DB {

	address := config.Conf.Get("sql.address").(string)
	user := config.Conf.Get("sql.user").(string)
	passwd := config.Conf.Get("sql.passwd").(string)
	driver := config.Conf.Get("sql.driver").(string)
	database := config.Conf.Get("sql.database").(string)
	db, err := sql.Open(driver, user+":"+passwd+"@tcp("+address+")/"+database)
	checkErr(err)
	return db
}

func InsertBaiDuHotList(baiduList []model.BDHotSearchItem) {
	dbConnection := _GetDBConnect()
	da := time.Now().Format("2006-01-02")
	for i, item := range baiduList {
		InsertValue(dbConnection, "baidu", "baidu_"+da+"_"+strconv.Itoa(i), item.Word, strconv.Itoa(item.HotRank), item.Desc,
			item.HotScore, item.HotTag, item.Url, da)
	}
}

func InsertZhiHuHotList(baiduList []model.ZhiHuQuestion) {
	dbConnection := _GetDBConnect()
	da := time.Now().Format("2006-01-02")
	for i, item := range baiduList {
		InsertValue(dbConnection, "zhihu", "zhihu_"+da+"_"+strconv.Itoa(i), item.Title, strconv.Itoa(i), item.Desc,
			item.HotScore, item.HotTag, item.Url, da)
	}
}

// 插入微博热搜
func InsertWeiboHotList(weiboList []model.WBHotBandItem) {
	dbConnection := _GetDBConnect()
	da := time.Now().Format("2006-01-02")
	for i, item := range weiboList {
		InsertValue(dbConnection, "weibo", "weibo_"+da+"_"+strconv.Itoa(i), item.Note, strconv.Itoa(item.Rank), item.WordScheme,
			strconv.Itoa(item.Num), item.HotTag+item.Category, item.Url, da)
	}

	dbConnection.Close()
}

func InsertValue(dbConnection *sql.DB, web_source, web_id, hot_note, hot_rank, hot_desc, hot_score, hot_tag, hot_url, hot_date string) {

	// 插入数据
	stmt, err := dbConnection.Prepare("INSERT hotSearchInfo SET web_source=?,web_id=?,hot_note=?,hot_rank=?,hot_desc=?,hot_score=?,hot_tag=?,hot_url=?,hot_date=?")
	checkErr(err)
	res, err := stmt.Exec(web_source, web_id, hot_note, hot_rank, hot_desc, hot_score, hot_tag, hot_url, hot_date)
	checkErr(err)
	id, err := res.LastInsertId()
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
