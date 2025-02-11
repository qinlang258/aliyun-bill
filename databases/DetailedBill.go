package database

type DetailedBill struct {
	ID            int     `gorm:"column:id"`
	Financial     string  `gorm:"column:financial;type:varchar(255)"`
	ProductCode   string  `gorm:"column:product_code;type:varchar(255)"`
	ProductName   string  `gorm:"column:product_name;type:varchar(255)"`
	ResourceID    string  `gorm:"column:resource_id;type:varchar(255)"`
	Name          string  `gorm:"column:name;type:varchar(255)"`
	ResourceGroup string  `gorm:"column:resource_group;type:varchar(255)"`
	UnBlendedCost float64 `gorm:"column:un_blended_cost;type:decimal(10,5)"`
	UsageType     string  `gorm:"column:usage_type;type:varchar(255)"`
	Labels        string  `gorm:"column:labels;type:varchar(255)"`
	Region        string  `gorm:"column:region;type:varchar(255)"`
	PublicIp      string  `gorm:"column:public_ip;type:varchar(255)"`
	PrivateIp     string  `gorm:"column:private_ip;type:varchar(255)"`
}

func (DetailedBill) TableName() string {
	return "detailed_bill"
}
