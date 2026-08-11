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
		OutPath:       "./rpc/sms/gen/query", // output directory, default value is ./query
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
		g.GenerateModel("sms_coupon"),
		g.GenerateModel("sms_coupon_record"),
		g.GenerateModel("sms_coupon_scope"),
		g.GenerateModel("sms_coupon_type"),
		g.GenerateModel("sms_seckill_session"),
		g.GenerateModel("sms_seckill_activity"),
		g.GenerateModel("sms_seckill_product"),
		g.GenerateModel("sms_seckill_reservation"),
		g.GenerateModel("sms_home_advertise"),
	)

	// Execute the generator
	g.Execute()
}
