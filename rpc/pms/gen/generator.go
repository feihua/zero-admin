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
		OutPath:       "./rpc/pms/gen/query", // output directory, default value is ./query
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
		g.GenerateModel("pms_product_brand"),
		g.GenerateModel("pms_product_category"),

		g.GenerateModel("pms_product_attribute_group"),
		g.GenerateModel("pms_product_attribute"),
		g.GenerateModel("pms_product_attribute_value"),

		g.GenerateModel("pms_product_spec"),
		g.GenerateModel("pms_product_spec_value"),

		g.GenerateModel("pms_product_spu"),
		g.GenerateModel("pms_product_sku"),

		g.GenerateModel("pms_feight_template"),
		g.GenerateModel("pms_member_price"),
		g.GenerateModel("pms_product_category_attribute_relation"),
		g.GenerateModel("pms_product_full_reduction"),
		g.GenerateModel("pms_product_ladder"),
	)

	// Execute the generator
	g.Execute()
}
