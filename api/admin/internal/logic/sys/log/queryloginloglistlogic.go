package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
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
		PageNum:     req.Current,
		PageSize:    req.PageSize,
		Browser:     strings.TrimSpace(req.Browser),   // 浏览器
		IpAddress:   strings.TrimSpace(req.IpAddress), // IP地址
		LoginStatus: req.LoginStatus,                  // 登录状态
		Os:          strings.TrimSpace(req.Os),        // 操作系统
		UserName:    strings.TrimSpace(req.UserName),  // 用户名
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询登录日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryLoginLogListData

	for _, item := range result.List {
		list = append(list, &types.QueryLoginLogListData{
			Id:          item.Id,          // 编号
			UserName:    item.UserName,    // 用户名
			LoginStatus: item.LoginStatus, // 登录状态
			IpAddress:   item.IpAddress,   // IP地址
			Browser:     item.Browser,     // 浏览器
			Os:          item.Os,          // 操作系统
			ErrorMsg:    item.ErrorMsg,    // 登录失败信息
			LoginTime:   item.LoginTime,   // 登录时间
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
