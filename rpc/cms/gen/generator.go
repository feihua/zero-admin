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
		OutPath:       "./rpc/cms/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	//var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:oMbPi5munxCsBSsiLoPV@tcp(110.41.179.89:3306)/zero-cms?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("cms_help"),
		g.GenerateModel("cms_help_category"),
		g.GenerateModel("cms_member_report"),
		g.GenerateModel("cms_preferred_area"),
		g.GenerateModel("cms_preferred_area_product_relation"),
		g.GenerateModel("cms_subject"),
		g.GenerateModel("cms_subject_category"),
		g.GenerateModel("cms_subject_comment"),
		g.GenerateModel("cms_subject_product_relation"),
		g.GenerateModel("cms_topic"),
		g.GenerateModel("cms_topic_category"),
		g.GenerateModel("cms_topic_comment"),
	)

	// Execute the generator
	g.Execute()
}
