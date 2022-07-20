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
		go func() {
			for i := 1; i <= 200; i++ {
				new(hs.New).RequestPageInfo(i, pf)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}

func taskPaoyou() {
	open.GetHs(1, 20, 1, open.Platfrom_paoyou)
}

func taskLi5apuu7() {
	open.GetHs(1, 21, 2, open.Platfrom_li5apuu7)
}

func taskGdian() {
	open.GetHs(1, 601, 10, open.Platfrom_G)
}

func taskMaodou() {
	stringList := []string{"https://nnp35.com/upload_json_live", "https://jsonmdtv.md29.tv/upload_json_live", "https://json.wtjfjil.cn/upload_json_live"}
	wg.Add(len(stringList))
	for _, name := range stringList {
		go func() {
			for i := 1; i < 121; i++ {
				maDouDao, Type, urlType := new(hs.New).MaodouReq(i, name)
				hs.DataParseSave(maDouDao, Type, urlType)
			}
			defer wg.Done()
		}()
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
	scheduler.Every(30).Seconds().Do(taskCape)
	//scheduler.Cron("*/5 * * * *").Do(taskMaodou)
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
