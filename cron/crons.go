package cron

/**
 * @title cron
 * @author xiongshao
 * @date 2022-07-20 08:51:19
 */

import (
	"github.com/go-co-op/gocron"
	"httpParse/hs"
	"httpParse/open"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func taskCape() {
	platform := []string{"hot/topics?", "node/news?", "node/topics?type=1&nodeId=258&", "node/topics?type=7&", "node/topics?type=3&nodeId=0&", "node/topics?type=1&nodeId=0&"}
	wg.Add(len(platform))
	for _, pf := range platform {
		go func(str string) {
			for i := 1; i <= 200; i++ {
				new(hs.New).RequestPageInfo(i, str)
			}
			defer wg.Done()
		}(pf)
	}
	wg.Wait()
}

// paoyou第二版
func taskPaoYou2() {
	contentsMap := make(map[string]string)
	//contentsMap["1"] = "日本"
	contentsMap["2"] = "国产"
	//contentsMap["3"] = "其他"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			for i := 1; i < 30; i++ {
				hs.RequestPaoyou2(i, pageNumber)
			}
		}(content)
	}
}

func taskTyms() {
	contentsMap := make(map[string]string)
	contentsMap["1"] = "精品自拍"
	contentsMap["2"] = "国产偷拍"
	contentsMap["3"] = "优选短视频"
	contentsMap["4"] = "主播大秀"
	contentsMap["5"] = "网爆事件"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			for i := 1; i < 20; i++ {
				id, _ := strconv.Atoi(pageNumber)
				hs.Tyms74Request(id, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
}

// 42--网爆流出 24--国产精品 41--短视频(目前39) 38--自拍偷拍 25--直播主播
func taskJinrijp() {
	contentsMap := make(map[string]string)
	contentsMap["42"] = "网爆流出"
	contentsMap["24"] = "国产精品"
	contentsMap["41"] = "短视频"
	contentsMap["38"] = "自拍偷拍"
	contentsMap["25"] = "直播主播"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			for i := 1; i < 20; i++ {
				id, _ := strconv.Atoi(pageNumber)
				hs.JinrijpRequest(id, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
	wg.Wait()
}

// 11--国产原创 27--变态另类 24--制服黑丝 25--亚洲有码 10--精彩时刻-105 21-热门事件 8-自拍偷拍 6-国产情色
func taskLujiujin() {
	contentsMap := make(map[string]string)
	contentsMap["11"] = "国产原创"
	contentsMap["27"] = "变态另类"
	contentsMap["24"] = "制服黑丝"
	contentsMap["21"] = "热门事件"
	contentsMap["10"] = "精彩时刻"
	contentsMap["8"] = "自拍偷拍"
	contentsMap["6"] = "国产情色"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			id, _ := strconv.Atoi(pageNumber)
			for i := 1; i < 20; i++ {
				hs.LujiujiuRequest(id, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
	wg.Wait()
}

// 茎淫
func taskJinyuisland() {
	contentsMap := make(map[string]string)
	contentsMap["10"] = "AV中文视频"
	contentsMap["3"] = "经典国产"
	contentsMap["2"] = "国产传媒"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			id, _ := strconv.Atoi(pageNumber)
			for i := 1; i < 40; i++ {
				hs.JinyuislandRequest(id, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
	wg.Wait()
}

// 红会所
func taskRedCross88() {
	contentsMap := make(map[string]string)
	contentsMap["91"] = "经典国产"
	contentsMap["75"] = "传媒原创"
	contentsMap["98"] = "另类视频"
	//contentsMap["1"] = "视频"
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			id, _ := strconv.Atoi(pageNumber)
			for i := 1; i < 20; i++ {
				hs.RedCross88Request(id, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
	wg.Wait()
}

func task69YW() {
	contentsMap := hs.ParseClassInfo(hs.Url69yw)
	wg.Add(len(contentsMap))
	for content := range contentsMap {
		go func(pageNumber string) {
			for i := 150; i < 200; i++ {
				text := strings.ReplaceAll(pageNumber, "/", "")
				classId, _ := strconv.Atoi(text)
				hs.ParseListVideoInfo(hs.GetHtml(hs.IsPage(i, pageNumber)), classId, i, contentsMap[pageNumber])
			}
			defer wg.Done()
		}(content)
	}
	wg.Wait()
}

func taskGga666() {
	pages := []int{23}
	wg.Add(len(pages))
	for _, page := range pages {
		go func(pageNumber int) {
			for i := 1; i < 20; i++ {
				hs.Gga666Request(pageNumber, i, "福利嫩妹")
			}
			defer wg.Done()
		}(page)
	}
	wg.Wait()
}

func taskLi5apuu7() {
	// 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32
	pages := []int{20, 28}
	wg.Add(len(pages))
	for _, page := range pages {
		go func(page int) {
			for i := 1; i < 20; i++ {
				hs.ExampleScrape(page, i)
			}
			defer wg.Done()
		}(page)
	}
	wg.Wait()
}

func taskGdian() {
	open.GetHs(1, 101, 5, open.Platfrom_G)
}

func taskMaomi() {
	//, "中文字幕", "亚洲无码"
	array := []string{"国产精品", "美女主播", "短视频"}
	wg.Add(len(array))
	for _, name := range array {
		go func(str string) {
			for i := 1; i <= 10; i++ {
				new(hs.New).MaomiRequest(i, str)
			}
			defer wg.Done()
		}(name)
	}
	wg.Wait()
}

func taskMaodou() {
	stringList := []string{"https://nnp35.com/upload_json_live", "https://jsonmdtv.md29.tv/upload_json_live", "https://json.wtjfjil.cn/upload_json_live"}
	wg.Add(len(stringList))
	for _, name := range stringList {
		go func(str string) {
			for i := 1; i < 140; i++ {
				maDouDao, Type, urlType := new(hs.New).MaodouReq(i, str)
				hs.DataParseSave(maDouDao, Type, urlType)
			}
			wg.Done()
		}(name)
	}
	wg.Wait()
}

func taskXyzVideo() {
	for i := 1; i < 227; i++ {
		hs.Xyzrequest(i)
	}
}

func taskgjyb() {
	wg.Add(3)
	go func() {
		open.T1001()
		wg.Done()
	}()
	go func() {
		open.T1002()
		wg.Done()
	}()
	go func() {
		open.T1003()
		wg.Done()
	}()
	wg.Wait()
}

// CronStartHs Hs定时器
func CronStartHs() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(55).Minutes().Do(task69YW)
	scheduler.StartAsync()
	scheduler.StartBlocking()
}

// CronStartGJYB 国家医保数据目录
func CronStartGJYB() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Hour().Do(taskCape)
	scheduler.StartAsync()
	scheduler.StartBlocking()
}
