package main

import "httpParse/cron"

func init() {

}

func main() {
	//hs.Mysql2Redis()
	cron.CronStartHs()
	//hs.Redis2Mysql()
}
