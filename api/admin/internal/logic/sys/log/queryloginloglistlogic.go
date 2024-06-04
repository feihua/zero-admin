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
		Browser:     strings.TrimSpace(req.Browser),
		IpAddress:   strings.TrimSpace(req.IpAddress),
		LoginStatus: req.LoginStatus,
		Os:          strings.TrimSpace(req.Os),
		UserName:    strings.TrimSpace(req.UserName),
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询登录日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryLoginLogListData

	for _, item := range result.List {
		list = append(list, &types.QueryLoginLogListData{
			Browser:     item.Browser,
			Id:          item.Id,
			IpAddress:   item.IpAddress,
			LoginStatus: item.LoginStatus,
			LoginTime:   item.LoginTime,
			Os:          item.Os,
			UserName:    item.UserName,
			ErrorMsg:    item.ErrorMsg,
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
