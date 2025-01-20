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
		OutPath:       "./rpc/pms/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	//var dsn = "root:12341qweqfsd2356@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/zero-pms?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("pms_album"),
		g.GenerateModel("pms_album_pic"),
		g.GenerateModel("pms_brand"),
		g.GenerateModel("pms_comment"),
		g.GenerateModel("pms_comment_replay"),
		g.GenerateModel("pms_feight_template"),
		g.GenerateModel("pms_member_price"),
		g.GenerateModel("pms_product"),
		g.GenerateModel("pms_product_attribute"),
		g.GenerateModel("pms_product_attribute_category"),
		g.GenerateModel("pms_product_attribute_value"),
		g.GenerateModel("pms_product_category"),
		g.GenerateModel("pms_product_category_attribute_relation"),
		g.GenerateModel("pms_product_collect"),
		g.GenerateModel("pms_product_full_reduction"),
		g.GenerateModel("pms_product_ladder"),
		g.GenerateModel("pms_product_operate_log"),
		g.GenerateModel("pms_product_vertify_record"),
		g.GenerateModel("pms_sku_stock"),
	)

	// Execute the generator
	g.Execute()
}
