package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpParse/li5apuu7"
	"net/http"
	"strconv"
	"time"
)

// 定时任务
func TimeTask() {
	var ch chan int
	// 定时任务
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for range ticker.C {
			// 方法
		}
		ch <- 1
	}()
	<-ch
}

func SaveInfo(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	start, _ := strconv.Atoi(c.Param("start"))
	for i := start; i < start+5; i++ {
		_, i2 := li5apuu7.ExampleScrape(page, i)
		fmt.Printf("第%d页", i2)
	}
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"start": start,
		"msg":   "操作成功",
		"url":   "http://localhost:8500/apuu7/" + strconv.Itoa(page) + "/" + strconv.Itoa(start+5),
	})
}
