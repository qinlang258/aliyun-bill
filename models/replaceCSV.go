package models

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReplaceCSV(csvName string) error {
	file, err := os.Open(csvName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	// 读取 CSV 数据
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// 在内存中修改字段
	for _, record := range records {
		// 假设要修改第二列（索引为1）的字段
		if len(record) > 1 {
			for key, values := range record {
				if values == "产品" {
					record[key] = "product_name"
				}
				if values == "产品Code" {
					record[key] = "product_code"
				}
				if values == "实例ID" {
					record[key] = "resource_id" // 将字段修改为 "NewValue"
				}
				if values == "实例昵称" {
					record[key] = "Name" // 将字段修改为 "NewValue"
				}
				if values == "资源组" {
					record[key] = "resource_group" // 将字段修改为 "NewValue"
				}
				// if values == "实例标签" {
				// 	record[key] = "labels" // 将字段修改为 "NewValue"
				// }
				if values == "实例规格" {
					record[key] = "instance_type" // 将字段修改为 "NewValue"
				}
				if values == "公网IP" {
					record[key] = "public_ip" // 将字段修改为 "NewValue"
				}
				if values == "私网IP" {
					record[key] = "private_ip" // 将字段修改为 "NewValue"
				}
				if values == "地域" {
					record[key] = "region"
				}
				if values == "应付金额" {
					record[key] = "un_blended_cost"
				}

			}
		}
	}

	newFile, err := os.Create("detailed_bill.csv")
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
