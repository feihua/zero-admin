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

// QueryOperateLogDetailLogic 查询操作日志详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:41
*/
type QueryOperateLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogDetailLogic {
	return &QueryOperateLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOperateLogDetail 查询操作日志详情
func (l *QueryOperateLogDetailLogic) QueryOperateLogDetail(req *types.QueryOperateLogDetailReq) (resp *types.QueryOperateLogDetailResp, err error) {
	detail, err := l.svcCtx.Operatelogservice.QueryOperateLogDetail(l.ctx, &sysclient.QueryOperateLogDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	item := types.QueryOperateLogDetailData{
		DeptName:          detail.DeptName,
		Id:                detail.Id,
		OperationIp:       detail.OperationIp,
		OperationName:     detail.OperationName,
		OperationParams:   detail.OperationParams,
		OperationResponse: detail.OperationResponse,
		OperationStatus:   detail.OperationStatus,
		OperationTime:     detail.OperationTime,
		OperationType:     detail.OperationType,
		OperationUrl:      detail.OperationUrl,
		RequestMethod:     detail.RequestMethod,
		Title:             detail.Title,
		UseTime:           detail.UseTime,
	}

	return &types.QueryOperateLogDetailResp{
		Code:    "000000",
		Message: "查询操作日志列表",
		Data:    item,
	}, nil
}
