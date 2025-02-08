package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogListLogic 查询操作日志列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type QueryOperateLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryOperateLogListLogic {
	return QueryOperateLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOperateLogList 查询操作日志列表
func (l *QueryOperateLogListLogic) QueryOperateLogList(req *types.QueryOperateLogListReq) (*types.QueryOperateLogListResp, error) {
	result, err := l.svcCtx.Operatelogservice.QueryOperateLogList(l.ctx, &sysclient.QueryOperateLogListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		OperationIp:     req.OperationIp,     // 操作地址
		OperationName:   req.OperationName,   // 操作人员
		DeptName:        req.DeptName,        // 部门名称
		OperationStatus: req.OperationStatus, // 操作状态
		OperationType:   req.OperationType,   // 操作类型
		OperationUrl:    req.OperationUrl,    // 操作方法
		Title:           req.Title,           // 系统模块
		Browser:         req.Browser,         // 浏览器
		Os:              req.Os,              // 操作系统
		RequestMethod:   req.RequestMethod,   // 请求方式
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOperateLogListData

	for _, detail := range result.List {
		list = append(list, &types.QueryOperateLogListData{
			Id:                detail.Id,                // 编号
			Title:             detail.Title,             // 系统模块
			OperationType:     detail.OperationType,     // 操作类型
			OperationName:     detail.OperationName,     // 操作人员
			RequestMethod:     detail.RequestMethod,     // 请求方式
			OperationUrl:      detail.OperationUrl,      // 操作方法
			OperationParams:   detail.OperationParams,   // 请求参数
			OperationResponse: detail.OperationResponse, // 响应参数
			OperationStatus:   detail.OperationStatus,   // 操作状态
			DeptName:          detail.DeptName,          // 部门名称
			UseTime:           detail.UseTime,           // 执行时长(毫秒)
			Browser:           detail.Browser,           // 浏览器
			Os:                detail.Os,                // 操作系统
			OperationIp:       detail.OperationIp,       // 操作地址
			OperationTime:     detail.OperationTime,     // 操作时间
		})
	}

	return &types.QueryOperateLogListResp{
		Code:     "000000",
		Message:  "查询操作日志列表",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil

}
