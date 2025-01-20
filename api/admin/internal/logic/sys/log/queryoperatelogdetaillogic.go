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
		Id: req.Id, // 编号
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	item := types.QueryOperateLogDetailData{
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
	}

	return &types.QueryOperateLogDetailResp{
		Code:    "000000",
		Message: "查询操作日志列表",
		Data:    item,
	}, nil
}
