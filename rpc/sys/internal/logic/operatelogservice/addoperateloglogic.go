package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOperateLogLogic 添加操作日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
type AddOperateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperateLogLogic {
	return &AddOperateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOperateLog 添加操作日志
func (l *AddOperateLogLogic) AddOperateLog(in *sysclient.AddOperateLogReq) (*sysclient.AddOperateLogResp, error) {

	uri := strings.Split(in.OperationUrl, "?")[0]
	// todo 待优化 量大的时候 ，把它们的关联缓存起来(redis)
	q := query.SysMenu
	menu, _ := q.WithContext(l.ctx).Select(q.MenuName).Where(q.BackgroundURL.Like("%" + uri + "%")).First()

	sysLog := &model.SysOperateLog{
		Title:             menu.MenuName,        // 系统模块
		OperationType:     in.OperationType,     // 操作类型
		OperationName:     in.OperationName,     // 操作人员
		RequestMethod:     in.RequestMethod,     // 请求方式
		OperationURL:      in.OperationUrl,      // 操作方法
		OperationParams:   in.OperationParams,   // 请求参数
		OperationResponse: in.OperationResponse, // 响应参数
		OperationStatus:   in.OperationStatus,   // 操作状态
		DeptName:          in.DeptName,          // 部门名称
		UseTime:           in.UseTime,           // 执行时长(毫秒)
		Browser:           in.Browser,           // 浏览器
		Os:                in.Os,                // 操作信息
		OperationIP:       in.OperationIp,       // 操作地址
		OperationTime:     time.Now(),           // 操作时间
	}

	if strings.Contains(in.OperationResponse, "000000") {
		sysLog.OperationStatus = 1
	}

	err := query.SysOperateLog.WithContext(l.ctx).Create(sysLog)
	if err != nil {
		logc.Errorf(l.ctx, "添加操作日志失败,参数:%+v,异常:%s", sysLog, err.Error())
		return nil, errors.New("添加操作日志失败")
	}

	return &sysclient.AddOperateLogResp{}, nil
}
