package main

import "httpParse/hs"

func main() {
	//hs.Mysql2Redis()
	//cron.CronStartHs()
	//hs.Redis2Mysql()
	for i := 1; i < 227; i++ {
		hs.Xyzrequest(i)
	}
}
