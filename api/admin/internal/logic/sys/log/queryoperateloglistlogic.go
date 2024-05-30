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
		OperationIp:     req.OperationIp,
		OperationName:   req.OperationName,
		DeptName:        req.DeptName,
		OperationStatus: req.OperationStatus,
		OperationType:   req.OperationType,
		OperationUrl:    req.OperationUrl,
		Title:           req.Title,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOperateLogListData

	for _, item := range result.List {
		list = append(list, &types.QueryOperateLogListData{
			DeptName:          item.DeptName,
			Id:                item.Id,
			OperationIp:       item.OperationIp,
			OperationName:     item.OperationName,
			OperationParams:   item.OperationParams,
			OperationResponse: item.OperationResponse,
			OperationStatus:   item.OperationStatus,
			OperationTime:     item.OperationTime,
			OperationType:     item.OperationType,
			OperationUrl:      item.OperationUrl,
			RequestMethod:     item.RequestMethod,
			Title:             item.Title,
			UseTime:           item.UseTime,
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
