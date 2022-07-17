package main

import "httpParse/open"

func main() {
	//gin := gin.Default()
	//gin.GET("/getData/:page/:start", src.SaveInfo)
	//gin.Run(":8500")
	//db.AutoCreateTable(xml.XmlInfo{})
	//hs.Mysql2Redis()
	open.GetHs(1, 20, 1, open.Platfrom_paoyou)
	//open.GetHs(1, 11, 1, open.Platfrom_li5apuu7)
	//open.GetHs(1, 121, 1, open.Platfrom_madou)
	//open.GetHs(1, 9, 1, open.Platfrom_maomi)
	// G. max page 25865 / 20 + 1 ==1294
	//open.GetHs(1, 601, 10, open.Platfrom_G)
	//open.GetHs(1, 101, 10, open.Platfrom_cape)
}
