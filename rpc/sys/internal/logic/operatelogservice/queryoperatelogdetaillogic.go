package operatelogservicelogic

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
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
		Id:           item.ID,                               // 操作日志id
		Title:        item.Title,                            // 模块标题
		OperateName:  item.OperateName,                      // 操作人员
		OperateUrl:   item.OperateURL,                       // 请求URL
		OperateIp:    item.OperateIP,                        // 主机地址
		OperateParam: item.OperateParam,                     // 请求参数
		JsonResult:   item.JSONResult,                       // 返回参数
		Status:       item.Status,                           // 操作状态(0:异常,正常)
		OperateTime:  time_util.TimeToStr(item.OperateTime), // 操作时间
		CostTime:     item.CostTime,                         // 消耗时间
	}

	_ = sonic.Unmarshal(item.Extra, data.Extra)
	return data, nil
}
