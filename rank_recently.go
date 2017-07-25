package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/configor"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Config = struct {
	DB struct {
		   Host     string
		   User     string `default:"root"`
		   Password string
		   Port     string `default:"3306"`
		   Database string
	   }
}{}

var engine *xorm.Engine

type RankRecently struct {
	Author       string
	Video_review int64
	Aid          int64
	Coins        int64
	Duration     int64
	Play         int64
	Pts          int64
	Mid          int64
	Type         int64
	Day          int64
	Title        string
	Time         int64
}

func main() {
	for {
		 users:= []RankRecently{}
		days := [...] int{1, 3, 7}
		typeids := [...]int{0, 1, 168, 3, 129, 4, 36, 160, 23, 119, 11, 155, 5}
		for _, day := range days {
			for _, typeid := range typeids {
				req := fasthttp.AcquireRequest()
				resp := fasthttp.AcquireResponse()
				req.SetRequestURI("http://www.bilibili.com/index/rank/all-0" + strconv.Itoa(day) + "-" + strconv.Itoa(typeid) + ".json")
				fasthttp.Do(req, resp)
				body, _ := resp.BodyGunzip()
				gjson.GetBytes(body, "rank").Get("list").ForEach(func(key, value gjson.Result) bool {
					minute, _ := strconv.Atoi(strings.Split(value.Get("duration").String(), ":")[0])
					second, _ := strconv.Atoi(strings.Split(value.Get("duration").String(), ":")[1])
					length := int64(minute * 60) + int64(second)
					user := RankRecently{Author:value.Get("author").String(), Duration:length, Video_review:value.Get("video_review").Int(), Aid:value.Get("aid").Int(), Mid:value.Get("mid").Int(), Coins:value.Get("coins").Int(), Pts:value.Get("pts").Int(), Play:value.Get("play").Int(), Day:int64(day), Type:int64(typeid), Title:value.Get("title").String(),Time:time.Now().Unix()}
					users = append(users, user)
					return true
				})

			}
		}
		configor.Load(&Config, "./config.yml")
		engine, _ = xorm.NewEngine("mysql", Config.DB.User + ":" + Config.DB.Password + "@tcp(" + Config.DB.Host + ":" + Config.DB.Port + ")/" + Config.DB.Database)
		engine.Sync2(new(RankRecently))
		var err error
		fmt.Println(len(users))
		affected, err := engine.Insert(&users)
		fmt.Println(affected,err)
		time.Sleep(700 *time.Second )
	}
}
