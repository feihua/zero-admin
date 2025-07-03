package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"
	"time"

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

	uri := strings.Split(in.OperateUrl, "?")[0]

	key := l.svcCtx.RedisKey + "background_url"
	name, _ := l.svcCtx.Redis.HgetCtx(l.ctx, key, uri)

	if name == "" {
		q := query.SysMenu
		_ = q.WithContext(l.ctx).Select(q.MenuName).Where(q.BackgroundURL.Eq(uri)).Scan(&name)
		if name == "" {
			_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, l.svcCtx.RedisKey+"background_url", uri)
			logc.Errorf(l.ctx, "添加操作日志失败,参数uri:%+v,异常:%s", uri, "菜单不存在,可能被删除了")
			return &sysclient.AddOperateLogResp{}, nil
		}
		_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, uri, name)
	}

	sysLog := &model.SysOperateLog{
		Title:           name,               // 模块标题
		BusinessType:    in.BusinessType,    // 业务类型（0其它 1新增 2修改 3删除）
		Method:          in.Method,          // 方法名称
		RequestMethod:   in.RequestMethod,   // 请求方式
		OperatorType:    in.OperatorType,    // 操作类别（0其它 1后台用户 2手机端用户）
		OperateName:     in.OperateName,     // 操作人员
		DeptName:        in.DeptName,        // 部门名称
		OperateURL:      in.OperateUrl,      // 请求URL
		OperateIP:       in.OperateIp,       // 主机地址
		OperateLocation: in.OperateLocation, // 操作地点
		OperateParam:    in.OperateParam,    // 请求参数
		JSONResult:      in.JsonResult,      // 返回参数
		Platform:        in.Platform,        // 平台信息
		Browser:         in.Browser,         // 浏览器类型
		Version:         in.Version,         // 浏览器版本
		Os:              in.Os,              // 操作系统
		Arch:            in.Arch,            // 体系结构信息
		Engine:          in.Engine,          // 渲染引擎信息
		EngineDetails:   in.EngineDetails,   // 渲染引擎详细信息
		Extra:           in.Extra,           // 其他信息（可选）
		Status:          in.Status,          // 操作状态(0:异常,正常)
		ErrorMsg:        in.ErrorMsg,        // 错误消息
		OperateTime:     time.Now(),         // 操作时间
		CostTime:        in.CostTime,        // 消耗时间
	}

	if strings.Contains(in.JsonResult, "000000") {
		sysLog.Status = 1
	}

	err := query.SysOperateLog.WithContext(l.ctx).Create(sysLog)
	if err != nil {
		logc.Errorf(l.ctx, "添加操作日志失败,参数:%+v,异常:%s", sysLog, err.Error())
		return nil, errors.New("添加操作日志失败")
	}

	return &sysclient.AddOperateLogResp{}, nil
}
