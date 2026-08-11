// configuration.go
package main

import (
	"gorm.io/driver/postgres"
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

	dsn := "host=129.204.203.29 user=root password=1a2341q!weqfsd2356T dbname=gozero port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
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
	)

	// Execute the generator
	g.Execute()
}
