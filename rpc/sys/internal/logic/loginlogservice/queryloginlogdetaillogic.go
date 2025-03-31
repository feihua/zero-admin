package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryLoginLogDetailLogic 查询系统登录日志表详情
/*
Author: LiuFeiHua
Date: 2025/1/20 17:46
*/
type QueryLoginLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLoginLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogDetailLogic {
	return &QueryLoginLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryLoginLogDetail 查询系统登录日志表详情
// 1.判断登录日志是否存在
func (l *QueryLoginLogDetailLogic) QueryLoginLogDetail(in *sysclient.QueryLoginLogDetailReq) (*sysclient.QueryLoginLogDetailResp, error) {
	item, err := query.SysLoginLog.WithContext(l.ctx).Where(query.SysLoginLog.ID.Eq(in.Id)).First()

	// 1.判断登录日志是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "登录日志不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("登录日志不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询登录日志异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询登录日志异常")
	}

	data := &sysclient.QueryLoginLogDetailResp{
		Id:            item.ID,                             // 登录日志id
		LoginName:     item.LoginName,                      // 登录账号
		Ipaddr:        item.Ipaddr,                         // 登录IP地址
		LoginLocation: item.LoginLocation,                  // 登录地点
		Platform:      item.Platform,                       // 平台信息
		Browser:       item.Browser,                        // 浏览器类型
		Version:       item.Version,                        // 浏览器版本
		Os:            item.Os,                             // 操作系统
		Arch:          item.Arch,                           // 体系结构信息
		Engine:        item.Engine,                         // 渲染引擎信息
		EngineDetails: item.EngineDetails,                  // 渲染引擎详细信息
		Extra:         item.Extra,                          // 其他信息（可选）
		Status:        item.Status,                         // 登录状态(0:失败,1:成功)
		Msg:           item.Msg,                            // 提示消息
		LoginTime:     time_util.TimeToStr(item.LoginTime), // 登录时间
	}

	return data, nil

}
