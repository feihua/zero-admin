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

	//var dsn = "12341qweqfsd2356@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/zero-sms?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("sms_coupon"),
		g.GenerateModel("sms_coupon_history"),
		g.GenerateModel("sms_coupon_product_category_relation"),
		g.GenerateModel("sms_coupon_product_relation"),
		g.GenerateModel("sms_flash_promotion_log"),
		g.GenerateModel("sms_flash_promotion"),
		g.GenerateModel("sms_flash_promotion_product_relation"),
		g.GenerateModel("sms_flash_promotion_session"),
		g.GenerateModel("sms_home_advertise"),
		g.GenerateModel("sms_home_brand"),
		g.GenerateModel("sms_home_new_product"),
		g.GenerateModel("sms_home_recommend_product"),
		g.GenerateModel("sms_home_recommend_subject"),
	)

	// Execute the generator
	g.Execute()
}
