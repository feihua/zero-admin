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
		Title:           req.Title,           // 模块标题
		BusinessType:    req.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          req.Method,          // 方法名称
		RequestMethod:   req.RequestMethod,   // 请求方式
		OperatorType:    req.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     req.OperateName,     // 操作人员
		DeptName:        req.DeptName,        // 部门名称
		OperateUrl:      req.OperateUrl,      // 请求URL
		OperateIp:       req.OperateIp,       // 主机地址
		OperateLocation: req.OperateLocation, // 操作地点
		Platform:        req.Platform,        // 平台信息
		Browser:         req.Browser,         // 浏览器类型
		Os:              req.Os,              // 操作系统
		Status:          req.Status,          // 操作状态(0:异常,正常)
		OperateTime:     req.OperateTime,     // 操作时间
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOperateLogListData

	for _, detail := range result.List {
		list = append(list, &types.QueryOperateLogListData{
			Id:              detail.Id,              // 操作日志id
			Title:           detail.Title,           // 模块标题
			BusinessType:    detail.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
			Method:          detail.Method,          // 方法名称
			RequestMethod:   detail.RequestMethod,   // 请求方式
			OperatorType:    detail.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
			OperateName:     detail.OperateName,     // 操作人员
			DeptName:        detail.DeptName,        // 部门名称
			OperateUrl:      detail.OperateUrl,      // 请求URL
			OperateIp:       detail.OperateIp,       // 主机地址
			OperateLocation: detail.OperateLocation, // 操作地点
			OperateParam:    detail.OperateParam,    // 请求参数
			JsonResult:      detail.JsonResult,      // 返回参数
			Platform:        detail.Platform,        // 平台信息
			Browser:         detail.Browser,         // 浏览器类型
			Version:         detail.Version,         // 浏览器版本
			Os:              detail.Os,              // 操作系统
			Arch:            detail.Arch,            // 体系结构信息
			Engine:          detail.Engine,          // 渲染引擎信息
			EngineDetails:   detail.EngineDetails,   // 渲染引擎详细信息
			Extra:           detail.Extra,           // 其他信息（可选）
			Status:          detail.Status,          // 操作状态(0:异常,正常)
			ErrorMsg:        detail.ErrorMsg,        // 错误消息
			OperateTime:     detail.OperateTime,     // 操作时间
			CostTime:        detail.CostTime,        // 消耗时间
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
