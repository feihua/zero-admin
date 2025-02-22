package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryLoginLogDetailLogic 查询登录日志详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:32
*/
type QueryLoginLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogDetailLogic {
	return &QueryLoginLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryLoginLogDetail 查询登录日志详情
func (l *QueryLoginLogDetailLogic) QueryLoginLogDetail(req *types.QueryLoginLogDetailReq) (resp *types.QueryLoginLogDetailResp, err error) {
	detail, err := l.svcCtx.LoginLogService.QueryLoginLogDetail(l.ctx, &sysclient.QueryLoginLogDetailReq{
		Id: req.Id, // 编号
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询登录日志详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	log := types.QueryLoginLogDetailData{
		Id:          detail.Id,          // 编号
		UserName:    detail.UserName,    // 用户名
		LoginStatus: detail.LoginStatus, // 登录状态
		IpAddress:   detail.IpAddress,   // IP地址
		Browser:     detail.Browser,     // 浏览器
		Os:          detail.Os,          // 操作系统
		ErrorMsg:    detail.ErrorMsg,    // 登录失败信息
		LoginTime:   detail.LoginTime,   // 登录时间
	}

	return &types.QueryLoginLogDetailResp{
		Code:    "000000",
		Message: "查询登录日志详情成功",
		Data:    log,
	}, nil
}
