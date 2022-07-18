package main

import (
	"github.com/gin-gonic/gin"
	"httpParse/open"
	"httpParse/sale"
	"strconv"
)

func taskCape() {
	open.GetHs(1, 101, 10, open.Platfrom_cape)
}

func taskPaoyou() {
	open.GetHs(1, 20, 1, open.Platfrom_paoyou)
}

func taskLi5apuu7() {
	open.GetHs(1, 11, 1, open.Platfrom_li5apuu7)
}

func taskGdian() {
	open.GetHs(1, 601, 10, open.Platfrom_G)
}

func ginTest(c *gin.Context) {
	shopId, _ := strconv.Atoi(c.Param("shopId"))
	Type := c.Param("Type")
	c.JSON(200, gin.H{
		"message": "成功",
		"data":    sale.SaleFind(shopId, Type),
	})

}

func main() {
	//db.AutoCreateTable(xml.XmlInfo{})
	//hs.Mysql2Redis()
	//open.GetHs(1, 20, 1, open.Platfrom_paoyou)
	//open.GetHs(1, 11, 1, open.Platfrom_li5apuu7)
	//open.GetHs(1, 121, 1, open.Platfrom_madou)
	//open.GetHs(1, 9, 1, open.Platfrom_maomi)
	// G. max page 25865 / 20 + 1 ==1294
	//open.GetHs(1, 101, 10, open.Platfrom_cape)
	//scheduler := gocron.NewScheduler(time.UTC)
	//scheduler.Every(5).Minutes().Do(taskCape)
	//scheduler.Every(3).Minutes().Do(taskPaoyou)
	//scheduler.Every(3).Minutes().Do(taskLi5apuu7)
	//scheduler.Every(3).Minutes().Do(taskGdian)
	//scheduler.StartBlocking()
	//data := orm.GetOracleData(`select * from djwk_b2c_order_doc a where a.internal_number = 240739`, DB.DS)
	//marshal, _ := json.Marshal(data)
	//fmt.Println(marshal)
	gin := gin.Default()
	gin.GET("/getData/:Type/:shopId", ginTest)
	gin.Run(":8521")
}
