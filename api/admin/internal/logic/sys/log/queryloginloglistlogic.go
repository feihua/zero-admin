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
		PageNum:       req.Current,
		PageSize:      req.PageSize,
		LoginName:     req.LoginName,     // 登录账号
		Ipaddr:        req.Ipaddr,        // 登录IP地址
		LoginLocation: req.LoginLocation, // 登录地点
		Platform:      req.Platform,      // 平台信息
		Browser:       req.Browser,       // 浏览器类型
		Os:            req.Os,            // 操作系统
		Status:        req.Status,        // 登录状态(0:失败,1:成功)
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询登录日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryLoginLogListData

	for _, detail := range result.List {
		list = append(list, &types.QueryLoginLogListData{
			Id:            detail.Id,            // 登录日志id
			LoginName:     detail.LoginName,     // 登录账号
			Ipaddr:        detail.Ipaddr,        // 登录IP地址
			LoginLocation: detail.LoginLocation, // 登录地点
			Platform:      detail.Platform,      // 平台信息
			Browser:       detail.Browser,       // 浏览器类型
			Version:       detail.Version,       // 浏览器版本
			Os:            detail.Os,            // 操作系统
			Arch:          detail.Arch,          // 体系结构信息
			Engine:        detail.Engine,        // 渲染引擎信息
			EngineDetails: detail.EngineDetails, // 渲染引擎详细信息
			Extra:         detail.Extra,         // 其他信息（可选）
			Status:        detail.Status,        // 登录状态(0:失败,1:成功)
			Msg:           detail.Msg,           // 提示消息
			LoginTime:     detail.LoginTime,     // 访问时间
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
