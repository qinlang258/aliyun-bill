# aliyun-bill
阿里云账单分析

## 说明
这里仅仅提取了账单的CSV 生成一个新的detailed_bill.csv文件，在MySQL里导入该账单，选择转化成英语的字段  

使用 select product_name,product_code,resource_id,name,SUM(un_blended_cost) from detailed_bill GROUP BY product_name,product_code,resource_id,name ORDER BY product_name 获取查询内容，并输出至新的汇总表格

