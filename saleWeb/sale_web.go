package saleWeb

import (
	"github.com/gin-gonic/gin"
	"httpParse/sale"
	"strconv"
)

/**
 * @title 销售查询服务端
 * @author xiongshao
 * @date 2022-07-20 08:53:54
 */

func ginTest(c *gin.Context) {
	shopId, _ := strconv.Atoi(c.Param("shopId"))
	Type := c.Param("Type")
	c.JSON(200, gin.H{
		"message": "成功",
		"data":    sale.SaleFind(shopId, Type),
	})
}

func SaleWebStart() {
	gin := gin.Default()
	gin.GET("/getData/:Type/:shopId", ginTest)
	gin.Run(":8521")
}
