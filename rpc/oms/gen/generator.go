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
		OutPath:       "./rpc/oms/gen/query", // output directory, default value is ./query
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
		g.GenerateModel("oms_company_address"),
		g.GenerateModel("oms_order_return_reason"),
		g.GenerateModel("oms_order_setting"),

		g.GenerateModel("oms_order_delivery"),
		g.GenerateModel("oms_order_return"),
		g.GenerateModel("oms_order_return_item"),
		g.GenerateModel("oms_order_main"),
		g.GenerateModel("oms_order_item"),
		g.GenerateModel("oms_order_operation_log"),
		g.GenerateModel("oms_order_payment"),
		g.GenerateModel("oms_order_promotion"),

		g.GenerateModel("oms_cart_item"),
		// g.GenerateModel("oms_order"),
		// g.GenerateModel("oms_order_item"),
		// g.GenerateModel("oms_order_operate_history"),
		// g.GenerateModel("oms_order_return_apply"),
	)

	// Execute the generator
	g.Execute()
}
