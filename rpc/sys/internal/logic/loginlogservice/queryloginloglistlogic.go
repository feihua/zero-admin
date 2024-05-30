package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryLoginLogListLogic 登录日志列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type QueryLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogListLogic {
	return &QueryLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryLoginLogList 登录日志列表
func (l *QueryLoginLogListLogic) QueryLoginLogList(in *sysclient.QueryLoginLogListReq) (*sysclient.QueryLoginLogListResp, error) {
	q := query.SysLoginLog.WithContext(l.ctx)
	if len(in.UserName) > 0 {
		q = q.Where(query.SysLoginLog.UserName.Like("%" + in.UserName + "%"))
	}
	if len(in.IpAddress) > 0 {
		q = q.Where(query.SysLoginLog.IPAddress.Like("%" + in.IpAddress + "%"))
	}

	if len(in.Browser) > 0 {
		q = q.Where(query.SysLoginLog.Browser.Like("%" + in.Browser + "%"))
	}
	if len(in.LoginStatus) > 0 {
		q = q.Where(query.SysLoginLog.LoginStatus.Like("%" + in.LoginStatus + "%"))
	}
	if len(in.Os) > 0 {
		q = q.Where(query.SysLoginLog.Os.Like("%" + in.Os + "%"))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询登录记录列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询登录记录列表信息失败")
	}

	var list []*sysclient.LoginLogListData
	for _, log := range result {
		list = append(list, &sysclient.LoginLogListData{
			Browser:     log.Browser,
			Id:          log.ID,
			IpAddress:   log.IPAddress,
			LoginStatus: log.LoginStatus,
			LoginTime:   log.LoginTime.Format("2006-01-02 15:04:05"),
			Os:          log.Os,
			UserName:    log.UserName,
		})
	}

	logc.Infof(l.ctx, "查询登录记录列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
