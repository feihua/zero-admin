package svc

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin/rpc/model/paymodel"
	"zero-admin/rpc/pay/internal/config"
)

type ServiceContext struct {
	c                config.Config
	WxRecordModel    paymodel.PayWxRecordModel
	WxMerchantsModel paymodel.PayWxMerchantsModel
	AlipayClient     *alipay.Client
}

const (
	//kAppId      = "2021000122672643"
	//kPrivateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCG/iWUj0IScEpQWeOxEdC8muaPNZdvmnXTeMXTLARMbOIlg3LsED1Nt4VqjZ/K2J+3jQGcGMtCGbIOTlJpBnjpM5n/IA+oiUmhooH6ftBJcDEpbixDAk9pvrW/LXE52V3+MMrpKKEwJMt+YOUKf//m4svx5aXgxzb7QlMk1F7TaYjxxbULcb2eZKhRfhfJwQYrt8DP8jRQq1Jo5m52tnvMrKQjyVJwJy44YTsEMrXRaPWZ0+XKMjUWTbnqdVECt9r2z/5M3r6x2j0qxiQAYBAGgeD/iazDtvhXCrrgb4nG7yEjOKhuVSN1mRItdmqW4TferOGCJZanxCFcdzrQkf2BAgMBAAECggEAcrbOIKyMrTaXMCjzAKnvBBduDgywn7pWnlpnYchp7rgohVBq/IfgUIa/7YhkXfAv6b79uzSmpYlIcjfEeFNztFiRaOhJ5iKkW6LJaaESRxX78QUav+barTXPJKLtMQeyhCvagsBwGYVrF/4nJQEY6Y+ZV/qbN6SS6Hm4RffijSxJKJv6CzUjkajxnDzJBz3ySR0z5L8Q0abRkb+PS8cgUO9VbiGcD4leQV3wXM1e3BFDBsQ7POVe2EAYWcGbDVHqV8hFM186zs2zsgMp4VzUODuPuBPaVbKnZcml3Z5aEIamJH52YtES3R5WvWWQ2We6xwmoGXmuGENHejtdqB21AQKBgQDE0Tu6yaP76lWxlq15TCeQORzVlWJ7zhcvoxEj4BTHTggs1r9dRLqLai9Cd2+cciirh+VeOFIUKSfX/yF+1iTYPO0qXVXtPVBkc8iVrRU1bDO7UCAWAMz36DteOSYsPL6roJrudmLuke1AQQz6I5ydRlupZpDYFsuxVNs2XhBtyQKBgQCvlbtCr8LEVpPP5Cye/KhBT3q8vORijFyz/U0Tq4IK4DRjKR2MSSrueR5Osq+qcAmqi7Fdn89XtjHQE1QFABLbHPE9F3f79M1/pndHABkOGXOHIxq8K+X9AshSRK1MoQp1pJLm2E4KiMxBSzaIvneug3ukPqeqm44SKG9uBuQN+QKBgQCE1peSzY+xcoseDo3NJZo6XGHawjWzS/kYPN5PsWk0z7Ty1opYYA/sEuIM4WHiXKaYh2NHAYpccx6iSV+JJO2/SPfltRNOyShedEs4wpZi9UHBNiZB046D8ClJwhbCmskyO3b2Zc8GKFXSHVWt6qVE/XzWTBSM1G3spVJDUp+SCQKBgQCpsNJmU4iuyWFW1BTPnixZ2h8rUn6CQ1bAWHfqH6GxMxdOEglNb9T+3a0Nr6EX3elpmlHSwsTW5uzjRBq6LmUKv8DhItJBfUgxKscxpgWQ28YL/0AyRVajG9JPt7GoUibSpTeXw8pAYg7Mt4y/wRvXW5jdlfPibS1znQJ72ksCuQKBgHOTEVs6M42vjXL/ijZPR1RBsY7rROKSsklpPjbyiHGrQeeJMWh4n1SNAVIrlaQmj3B8gCaCAFEJX702ajOEfnPrnDTO3V0w9rSP9S1PGoRr86JTOT8g3A+0d7Vp2BGnHZ/9s9Ed3cZzTev2TBZHdKeoTlYQfYaD7n/kttpI0Xx2"
	kAppId      = "2021003185602741"
	kPrivateKey = "123456"

	kServerPort = "9989"
	// TODO 设置回调地址域名
	kServerDomain = ""
)

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	client, err := alipay.New(kAppId, kPrivateKey, true)
	if err != nil {
		panic("初始化支付宝失败")
	}
	return &ServiceContext{
		c:                c,
		WxRecordModel:    paymodel.NewPayWxRecordModel(sqlConn),
		WxMerchantsModel: paymodel.NewPayWxMerchantsModel(sqlConn),
		AlipayClient:     client,
	}
}
