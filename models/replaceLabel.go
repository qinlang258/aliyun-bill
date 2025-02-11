package models

import (
	database "aliyun-bill/databases"
	"fmt"
)

func ReplaceLabel(tableName string) {
	db := database.ContentMysql(tableName)

	data := []database.DetailedBill{}

	_ = db.Table("detailed_bill").Find(&data)

	for key, values := range data {
		fmt.Println(key, values.Labels)
	}

}
