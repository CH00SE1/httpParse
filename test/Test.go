package main

import (
	"gorm.io/gorm"
	"httpParse/db"
)

/**
 * @title xxx
 * @author xiongshao
 * @date 2022-07-05 08:48:51
 */

type Give struct {
	gorm.Model
	ClassId  int
	DrugName string
	DrugCode string
	DrugId   int
}

func GenTable() {
	db.AutoCreateTable(Give{})
}
