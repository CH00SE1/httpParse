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

func taskPaoyou() {
	map1, array1 := hs.PaoyouFindClass()
	wg.Add(len(array1))
	for _, arr := range array1 {
		go func(classname string) {
			for i := 1; i < 21; i++ {
				hs.Paoyou(i, classname, map1)
			}
			defer wg.Done()
		}(arr)
	}
	wg.Wait()
}

func taskLi5apuu7() {
	// 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32
	pages := []int{20, 28}
	wg.Add(len(pages))
	for _, page := range pages {
		go func(page int) {
			for i := 1; i < 11; i++ {
				hs.ExampleScrape(page, i)
			}
			time.Sleep(5 * time.Second)
			defer wg.Done()
		}(page)
	}
	wg.Wait()
}

func taskGdian() {
	open.GetHs(1, 101, 5, open.Platfrom_G)
}

func taskMaomi() {
	array := []string{"国产精品", "美女主播", "短视频", "中文字幕"}
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
			for i := 1; i < 141; i++ {
				maDouDao, Type, urlType := new(hs.New).MaodouReq(i, str)
				hs.DataParseSave(maDouDao, Type, urlType)
			}
			wg.Done()
		}(name)
	}
	wg.Wait()
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

// Hs定时器
func CronStartHs() {
	scheduler := gocron.NewScheduler(time.UTC)
	//scheduler.Cron("0 */1 * * * ").Seconds().Do(taskCape)
	//scheduler.Cron("*/5 * * * *").Do(taskMaodou)
	scheduler.Every(40).Minutes().Do(taskLi5apuu7)
	scheduler.StartAsync()
	scheduler.StartBlocking()
}

// 国家医保数据目录
func CronStartGJYB() {
	scheduler := gocron.NewScheduler(time.UTC)
	//scheduler.Cron("*/10 * * * *").Do(open.T1001)
	scheduler.Every(1).Hour().Do(taskgjyb)
	scheduler.StartAsync()
	scheduler.StartBlocking()
}
