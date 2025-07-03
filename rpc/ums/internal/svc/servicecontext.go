package svc

import (
	"fmt"
	"github.com/feihua/zero-admin/pkg/mq"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type ServiceContext struct {
	Config                             config.Config
	DB                                 *gorm.DB
	RabbitMQ                           *mq.RabbitMQ
	MemberBrandAttentionModel          model.MemberBrandAttentionModel
	MemberBrowseRecordModel            model.MemberBrowseRecordModel
	MemberProductCategoryRelationModel model.MemberProductCategoryRelationModel
	MemberProductCollectionModel       model.MemberProductCollectionModel
	Redis                              *redis.Redis
	RedisKey                           string // redis的模块统一前缀
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisKey := c.RpcServerConf.Redis.Key
	rds := redis.MustNewRedis(c.RpcServerConf.Redis.RedisConf)
	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 settingLogConfig(),
	})
	if err != nil {
		logx.Errorf("mysql连接失败：%+v", err)
		panic(err)
	}

	logx.Info("mysql连接成功")
	query.SetDefault(db)

	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Rabbitmq.UserName, c.Rabbitmq.Password, c.Rabbitmq.Host, c.Rabbitmq.Port)
	rabbitmq := mq.NewRabbitMQSimple(mqUrl)

	_, err = mon.NewModel(c.Mongo.Datasource, c.Mongo.Db, "test")
	if err != nil {
		logx.Errorf("mongo连接失败：%+v", err)
		panic(err)
	}
	logx.Info("mongo连接成功")
	MemberBrandAttention := model.NewMemberBrandAttentionModel(c.Mongo.Datasource, c.Mongo.Db, "ums_member_brand_attention")
	MemberBrowseRecordModel := model.NewMemberBrowseRecordModel(c.Mongo.Datasource, c.Mongo.Db, "ums_member_browse_record")
	MemberProductCategoryRelationModel := model.NewMemberProductCategoryRelationModel(c.Mongo.Datasource, c.Mongo.Db, "ums_member_product_category_relation")
	MemberProductCollectionModel := model.NewMemberProductCollectionModel(c.Mongo.Datasource, c.Mongo.Db, "ums_member_product_collection")
	return &ServiceContext{
		Config:                             c,
		DB:                                 db,
		RabbitMQ:                           rabbitmq,
		MemberBrandAttentionModel:          MemberBrandAttention,
		MemberBrowseRecordModel:            MemberBrowseRecordModel,
		MemberProductCategoryRelationModel: MemberProductCategoryRelationModel,
		MemberProductCollectionModel:       MemberProductCollectionModel,
		Redis:                              rds,
		RedisKey:                           redisKey,
	}
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	logx.Infof(format, args...)
}

// init log config
func settingLogConfig() logger.Interface {
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return newLogger
}
