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

	var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("ums_member_info"),
		g.GenerateModel("ums_member_level"),
		g.GenerateModel("ums_member_address"),
		g.GenerateModel("ums_member_tag"),
		g.GenerateModel("ums_member_tag_relation"),
		g.GenerateModel("ums_member_task"),
		g.GenerateModel("ums_member_task_relation"),
		g.GenerateModel("ums_member_sign_log"),
		g.GenerateModel("ums_member_points_log"),
		g.GenerateModel("ums_member_growth_log"),
		g.GenerateModel("ums_member_consume_setting"),

		// g.GenerateModel("ums_integration_consume_setting"),
		// g.GenerateModel("ums_member_brand_attention"),
		g.GenerateModel("ums_member_login_log"),
		// g.GenerateModel("ums_member_product_category_relation"),
		// g.GenerateModel("ums_member_product_collection"),
		// g.GenerateModel("ums_member_read_history"),
		g.GenerateModel("ums_member_rule_setting"),
		g.GenerateModel("ums_member_statistics_info"),
	)

	// Execute the generator
	g.Execute()
}
