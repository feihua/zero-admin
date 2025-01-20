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
		OutPath:       "./rpc/ums/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	//var dsn = "12341qweqfsd2356@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/zero-ums?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("ums_growth_change_history"),
		g.GenerateModel("ums_integration_change_history"),
		g.GenerateModel("ums_integration_consume_setting"),
		g.GenerateModel("ums_member"),
		g.GenerateModel("ums_member_brand_attention"),
		g.GenerateModel("ums_member_level"),
		g.GenerateModel("ums_member_login_log"),
		g.GenerateModel("ums_member_member_tag_relation"),
		g.GenerateModel("ums_member_product_category_relation"),
		g.GenerateModel("ums_member_product_collection"),
		g.GenerateModel("ums_member_read_history"),

		g.GenerateModel("ums_member_receive_address"),
		g.GenerateModel("ums_member_rule_setting"),
		g.GenerateModel("ums_member_statistics_info"),
		g.GenerateModel("ums_member_tag"),
		g.GenerateModel("ums_member_task"),
	)

	// Execute the generator
	g.Execute()
}
