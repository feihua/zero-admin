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
		Id:              item.ID,                               // 操作日志id
		Title:           item.Title,                            // 模块标题
		BusinessType:    item.BusinessType,                     // 业务类型（0其它 1新增 2修改 3删除）
		Method:          item.Method,                           // 方法名称
		RequestMethod:   item.RequestMethod,                    // 请求方式
		OperatorType:    item.OperatorType,                     // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     item.OperateName,                      // 操作人员
		DeptName:        item.DeptName,                         // 部门名称
		OperateUrl:      item.OperateURL,                       // 请求URL
		OperateIp:       item.OperateIP,                        // 主机地址
		OperateLocation: item.OperateLocation,                  // 操作地点
		OperateParam:    item.OperateParam,                     // 请求参数
		JsonResult:      item.JSONResult,                       // 返回参数
		Platform:        item.Platform,                         // 平台信息
		Browser:         item.Browser,                          // 浏览器类型
		Version:         item.Version,                          // 浏览器版本
		Os:              item.Os,                               // 操作系统
		Arch:            item.Arch,                             // 体系结构信息
		Engine:          item.Engine,                           // 渲染引擎信息
		EngineDetails:   item.EngineDetails,                    // 渲染引擎详细信息
		Extra:           item.Extra,                            // 其他信息（可选）
		Status:          item.Status,                           // 操作状态(0:异常,正常)
		ErrorMsg:        item.ErrorMsg,                         // 错误消息
		OperateTime:     time_util.TimeToStr(item.OperateTime), // 操作时间
		CostTime:        item.CostTime,                         // 消耗时间
	}

	return data, nil
}
