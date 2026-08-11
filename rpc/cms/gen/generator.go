// configuration.go
package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
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

	dsn := "host=129.204.203.29 user=root password=1a2341q!weqfsd2356T dbname=gozero port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	var curDB, curSchema string
	if err := db.Raw("SELECT current_database()").Scan(&curDB).Error; err != nil {
		fmt.Printf("current_database() query failed: %v", err)
	}
	if err := db.Raw("SELECT current_schema()").Scan(&curSchema).Error; err != nil {
		fmt.Printf("current_schema() query failed: %v", err)
	}

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

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}
