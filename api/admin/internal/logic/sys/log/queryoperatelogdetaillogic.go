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
	}

	return &types.QueryOperateLogDetailResp{
		Code:    "000000",
		Message: "查询操作日志详情成功",
		Data:    item,
	}, nil
}
