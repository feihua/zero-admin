package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogDetailLogic 查询系统操作日志详情
/*
Author: LiuFeiHua
Date: 2024/5/30 11:11
*/
type QueryOperateLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogDetailLogic {
	return &QueryOperateLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOperateLogDetail 查询系统操作日志详情
// 1.判断操作日志是否存在
func (l *QueryOperateLogDetailLogic) QueryOperateLogDetail(in *sysclient.QueryOperateLogDetailReq) (*sysclient.QueryOperateLogDetailResp, error) {
	item, err := query.SysOperateLog.WithContext(l.ctx).Where(query.SysOperateLog.ID.Eq(in.Id)).First()

	// 1.判断操作日志是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "操作日志不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("操作日志不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询操作日志异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询操作日志异常")
	}

	data := &sysclient.QueryOperateLogDetailResp{
		Id:                item.ID,                                 // 编号
		Title:             item.Title,                              // 系统模块
		OperationType:     item.OperationType,                      // 操作类型
		OperationName:     item.OperationName,                      // 操作人员
		RequestMethod:     item.RequestMethod,                      // 请求方式
		OperationUrl:      item.OperationURL,                       // 操作方法
		OperationParams:   item.OperationParams,                    // 请求参数
		OperationResponse: item.OperationResponse,                  // 响应参数
		OperationStatus:   item.OperationStatus,                    // 操作状态
		DeptName:          item.DeptName,                           // 部门名称
		UseTime:           item.UseTime,                            // 执行时长(毫秒)
		Browser:           item.Browser,                            // 浏览器
		Os:                item.Os,                                 // 操作信息
		OperationIp:       item.OperationIP,                        // 操作地址
		OperationTime:     time_util.TimeToStr(item.OperationTime), // 操作时间
	}

	return data, nil
}
