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
		PageNum:     req.Current,
		PageSize:    req.PageSize,
		OperateName: req.OperateName, // 操作人员
		OperateUrl:  req.OperateUrl,  // 请求URL
		OperateIp:   req.OperateIp,   // 主机地址
		Status:      req.Status,      // 操作状态(0:异常,正常)
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOperateLogListData

	for _, item := range result.List {
		list = append(list, &types.QueryOperateLogListData{
			Id:           item.Id,           // 日志主键
			OperateName:  item.OperateName,  // 操作人员
			OperateIp:    item.OperateIp,    // 主机地址
			OperateUrl:   item.OperateUrl,   // 请求url
			OperateParam: item.OperateParam, // 请求参数
			JsonResult:   item.JsonResult,   // 返回参数
			Extra:        item.Extra,        // 其他信息（可选）
			Status:       item.Status,       // 操作状态(0:异常,正常)
			CostTime:     item.CostTime,     // 消耗时间
			OperateTime:  item.OperateTime,  // 操作时间
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
