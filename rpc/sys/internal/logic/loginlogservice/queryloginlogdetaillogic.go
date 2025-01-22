package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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
func (l *QueryLoginLogDetailLogic) QueryLoginLogDetail(in *sysclient.QueryLoginLogDetailReq) (*sysclient.QueryLoginLogDetailResp, error) {
	item, err := query.SysLoginLog.WithContext(l.ctx).Where(query.SysLoginLog.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询系统登录日志详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询系统登录日志详情失败")
	}

	if item == nil {
		logc.Errorf(l.ctx, "查询系统登录日志详情失败,参数：%+v,登录日志不存在", in)
		return nil, errors.New("查询系统登录日志详情失败,登录日志不存在")
	}

	data := &sysclient.QueryLoginLogDetailResp{
		Id:          item.ID,                             // 编号
		UserName:    item.UserName,                       // 用户名
		LoginStatus: item.LoginStatus,                    // 登录状态
		IpAddress:   item.IPAddress,                      // IP地址
		Browser:     item.Browser,                        // 浏览器
		Os:          item.Os,                             // 操作系统
		ErrorMsg:    item.ErrorMsg,                       // 登录失败信息
		LoginTime:   time_util.TimeToStr(item.LoginTime), // 登录时间
	}

	return data, nil

}
