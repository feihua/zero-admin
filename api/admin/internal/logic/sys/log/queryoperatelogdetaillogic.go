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
		logc.Errorf(l.ctx, "查询操作日志详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	item := types.QueryOperateLogDetailData{
		Id:           detail.Id,           // 日志主键
		Title:        detail.Title,        // 模块标题
		OperateName:  detail.OperateName,  // 操作人员
		OperateIp:    detail.OperateIp,    // 主机地址
		OperateUrl:   detail.OperateUrl,   // 请求url
		OperateParam: detail.OperateParam, // 请求参数
		JsonResult:   detail.JsonResult,   // 返回参数
		Extra:        detail.Extra,        // 其他信息（可选）
		Status:       detail.Status,       // 操作状态(0:异常,正常)
		CostTime:     detail.CostTime,     // 消耗时间
		OperateTime:  detail.OperateTime,  // 操作时间
	}

	return &types.QueryOperateLogDetailResp{
		Code:    "000000",
		Message: "查询操作日志详情成功",
		Data:    item,
	}, nil
}
