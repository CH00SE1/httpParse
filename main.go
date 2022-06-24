package main

/**
 * @title 主程序启动
 * @author xiongshao
 * @date 2022-06-24 09:36:54
 */

import (
	"github.com/gin-gonic/gin"
	"httpParse/src"
)

func main() {
	gin := gin.Default()
	gin.PUT("/getData/:page/:start", src.SaveInfo)
	gin.Run(":8500")
}
