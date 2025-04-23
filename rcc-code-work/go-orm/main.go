package main

import (
	"fmt"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SysConfig struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	ConfigKey   string    `gorm:"column:config_key;unique;not null" json:"config_key"`
	ConfigValue string    `gorm:"column:config_value" json:"config_value"`
	Description string    `gorm:"column:description" json:"description"`
	Status      byte      `gorm:"column:status" json:"status"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

func main() {
	dsn := "root:duoguan@mysql@123@(47.110.151.225:3307)/trx?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 开启单数表名
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Create
	// db.Table("sys_config").Create(&SysConfig{ConfigKey: "index_switch", ConfigValue: "1", Description: "首页开关", Status: 1, CreateTime: time.Now(), UpdateTime: time.Now()})

	// var ss SysConfig
	// db.Table("sys_config").Find(&ss, "config_key=?", "test")

	// fmt.Println(ss)

	// db.Table("sys_config").Where("config_key=?", "index_switch").Updates(SysConfig{ConfigValue: "2"})

	// var sss SysConfig
	// db.Table("sys_config").Where("config_key=?", "index_switch").Find(&sss)

	// fmt.Println(sss)

	// sys := make([]SysConfig, 10)

	// db.Table("sys_config").Not([]int{1, 2}).Find(&sys)

	// fmt.Println(sys)

	// var sys SysConfig
	// db.Table("sys_config").Where(SysConfig{ConfigKey: "my_switch"}).Attrs(SysConfig{ConfigValue: "1", Description: "我的开关", Status: 1, CreateTime: time.Now(), UpdateTime: time.Now()}).FirstOrCreate(&sys)

	// db.Table("sys_config").Where(SysConfig{ConfigKey: "my_switch"}).First(&sys)

	// db.Table("sys_config").Where(SysConfig{ConfigKey: "my_switch"}).Assign(SysConfig{ConfigValue: "2"}).FirstOrCreate(&sys)

	// rows, err := db.Table("sys_config").Where("config_key is not null").Rows()
	// if err == nil {
	// 	for rows.Next() {
	// 		var sys SysConfig
	// 		db.ScanRows(rows, &sys)
	// 		fmt.Println(sys)
	// 	}
	// }

	var syss []SysConfig
	syss = make([]SysConfig, 0)
	db.Table("sys_config").Where("config_key is not null").Find(&syss)
	fmt.Println(syss)
	// fmt.Println(sys)

	// tx := db.Table("sys_config").Where("id is not null")
	// tx.Where("config_key is not null").Find(&syss)

	// fmt.Println(syss)

	// db.Scopes(KeyIsNotNull, ValueIsNotNull, IdIsNotNull).Find(&syss)
	// fmt.Println(syss)

	// count := 0
	// db.Table("sys_config").Where("config_key is not null").Find(&syss, "id in (?)", []int{10, 11, 12, 13, 14, 16}).Count(&count)
	// fmt.Println(count)
	// fmt.Println(len(syss))

	// db.Table("sys_config").Where("id = 14").Find(&syss)
	// syss[0].ConfigValue = "14"
	// db.Table("sys_config").Save(&syss[0])

	// db.Table("sys_config").Save(&SysConfig{ConfigKey: "14", ConfigValue: "14", Description: "14", Status: 1, CreateTime: time.Now(), UpdateTime: time.Now()})

	// var sys SysConfig
	// db.Table("sys_config").First(&sys)
	// fmt.Println(sys)

	// sys := &SysConfig{ID: 14, ConfigValue: "14"}
	// syss = append(syss, SysConfig{ID: 14, ConfigValue: "14"})
	// db.Table("sys_config").Model(&SysConfig{}).Where("id = ?", 14).Update("status", 1)
	// db.Table("sys_config").Model(&SysConfig{}).Where("id = ?", 14).Updates(map[string]interface{}{"config_value": "17", "description": "17"})
	// db.Transaction(func(tx *gorm.DB) error {
	// 	var sys SysConfig
	// 	tx.Table("sys_config").Where("id = ?", 14).Find(&sys)
	// 	sys.ConfigValue = "18"
	// 	tx.Table("sys_config").Save(&sys)

	// 	if err := tx.Table("sys_config").Where("id = ?", 16).Update("config_key", "15").Error; err != nil {
	// 		return err // 触发回滚
	// 	}
	// 	return nil
	// })

	// db.CreateInBatches([]SysConfig{
	// 	{ConfigKey: "1", ConfigValue: "1", Description: "1", Status: 1, CreateTime: time.Now(), UpdateTime: time.Now()},
	// 	{ConfigKey: "2", ConfigValue: "2", Description: "2", Status: 1, CreateTime: time.Now(), UpdateTime: time.Now()},
	// }, 2)
}

func (s *SysConfig) AfterFind(tx *gorm.DB) (err error) {
	fmt.Printf("配置 %s 查询成功\n", s.ConfigKey)
	return
}

func KeyIsNotNull(tx *gorm.DB) *gorm.DB {
	return tx.Where("config_key is not null")
}

func ValueIsNotNull(tx *gorm.DB) *gorm.DB {
	return tx.Where("config_value is not null")
}

func IdIsNotNull(tx *gorm.DB) *gorm.DB {
	return tx.Where("id is not null")
}
