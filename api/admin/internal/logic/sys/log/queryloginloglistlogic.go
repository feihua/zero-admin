package log

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryLoginLogListLogic 查询登录日志列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type QueryLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryLoginLogListLogic {
	return QueryLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryLoginLogList 查询登录日志列表
func (l *QueryLoginLogListLogic) QueryLoginLogList(req *types.QueryLoginLogListReq) (*types.QueryLoginLogListResp, error) {
	result, err := l.svcCtx.LoginLogService.QueryLoginLogList(l.ctx, &sysclient.QueryLoginLogListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		LoginName: req.LoginName, // 登录账号
		Ipaddr:    req.Ipaddr,    // 登录IP地址
		Status:    req.Status,    // 登录状态(0:失败,1:成功)
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询登录日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryLoginLogListData

	for _, item := range result.List {
		list = append(list, &types.QueryLoginLogListData{
			Id:        item.Id,        // 访问id
			LoginName: item.LoginName, // 登录账号
			IpAddr:    item.Ipaddr,    // 登录ip地址
			Extra:     item.Extra,     // 其他信息（可选）
			Status:    item.Status,    // 登录状态(0:失败,1:成功)
			Msg:       item.Msg,       // 提示消息
			LoginTime: item.LoginTime, // 访问时间
		})
	}

	return &types.QueryLoginLogListResp{
		Code:     "000000",
		Message:  "查询登录日志成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
