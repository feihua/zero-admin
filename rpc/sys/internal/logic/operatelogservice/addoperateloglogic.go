package operatelogservicelogic

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/ua-parser/uap-go/uaparser"
	"github.com/zeromicro/go-zero/core/logc"

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
		_ = q.WithContext(l.ctx).Select(q.MenuName).Where(q.APIURL.Eq(uri)).Scan(&name)
		if name == "" {
			_, _ = l.svcCtx.Redis.HdelCtx(l.ctx, l.svcCtx.RedisKey+"background_url", uri)
			return &sysclient.AddOperateLogResp{}, nil
		}
		_ = l.svcCtx.Redis.HsetCtx(l.ctx, key, uri, name)
	}

	parser := uaparser.NewFromSaved()
	ua := parser.Parse(in.Extra)

	browser := ua.UserAgent.Family + " " + ua.UserAgent.Major
	os := ua.Os.Family + " " + ua.Os.Major

	data, _ := sonic.Marshal(map[string]string{
		"loginOs":      os,      // 登录os
		"loginBrowser": browser, // 登录浏览器
	})
	sysLog := &model.SysOperateLog{
		OperateName:  in.OperateName,  // 操作人员
		OperateIP:    in.OperateIp,    // 主机地址
		OperateURL:   in.OperateUrl,   // 请求URL
		OperateParam: in.OperateParam, // 请求参数
		JSONResult:   in.JsonResult,   // 返回参数
		Extra:        data,            // 其他信息（可选）
		Status:       in.Status,       // 操作状态(0:异常,正常)
		CostTime:     in.CostTime,     // 消耗时间
		OperateTime:  time.Now(),      // 操作时间
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
