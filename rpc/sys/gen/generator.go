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
		OutPath:       "./rpc/sys/gen/query", // output directory, default value is ./query
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
		g.GenerateModel("sys_dept"),
		g.GenerateModel("sys_dict_type"),
		g.GenerateModelAs("sys_dict_data", "SysDictData"),
		g.GenerateModel("sys_post"),
		g.GenerateModel("sys_user_post"),
		g.GenerateModel("sys_operate_log", gen.FieldType("extra", "datatypes.JSON")),
		g.GenerateModel("sys_login_log", gen.FieldType("extra", "datatypes.JSON")),
		g.GenerateModel("sys_menu"),
		g.GenerateModel("sys_role"),
		g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_user", gen.FieldType("last_login_info", "datatypes.JSON")),
		g.GenerateModel("sys_post"),
		g.GenerateModel("sys_notice"),
		g.GenerateModel("sys_user_role"),
	)

	// Execute the generator
	g.Execute()
}
