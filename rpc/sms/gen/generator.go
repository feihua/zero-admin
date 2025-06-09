// configuration.go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./rpc/sms/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("sms_coupon"),
		g.GenerateModel("sms_coupon_record"),
		g.GenerateModel("sms_coupon_scope"),
		g.GenerateModel("sms_coupon_type"),
		g.GenerateModel("sms_seckill_session"),
		g.GenerateModel("sms_seckill_activity"),
		g.GenerateModel("sms_seckill_product"),
		g.GenerateModel("sms_seckill_reservation"),
		g.GenerateModel("sms_home_advertise"),
	)

	// Execute the generator
	g.Execute()
}
