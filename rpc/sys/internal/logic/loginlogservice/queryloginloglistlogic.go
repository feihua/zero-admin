package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
	loginLog := query.SysLoginLog
	q := loginLog.WithContext(l.ctx)
	if len(in.UserName) > 0 {
		q = q.Where(loginLog.UserName.Like("%" + in.UserName + "%"))
	}
	if len(in.IpAddress) > 0 {
		q = q.Where(loginLog.IPAddress.Like("%" + in.IpAddress + "%"))
	}

	if len(in.Browser) > 0 {
		q = q.Where(loginLog.Browser.Like("%" + in.Browser + "%"))
	}
	if len(in.LoginStatus) > 0 {
		q = q.Where(loginLog.LoginStatus.Like("%" + in.LoginStatus + "%"))
	}
	if len(in.Os) > 0 {
		q = q.Where(loginLog.Os.Like("%" + in.Os + "%"))
	}

	offset := (in.PageNum - 1) * in.PageSize
	result, count, err := q.Order(loginLog.ID.Desc()).FindByPage(int(offset), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询登录记录列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询登录记录列表信息失败")
	}

	var list []*sysclient.LoginLogListData
	for _, item := range result {
		list = append(list, &sysclient.LoginLogListData{
			Id:          item.ID,                             // 编号
			UserName:    item.UserName,                       // 用户名
			LoginStatus: item.LoginStatus,                    // 登录状态
			IpAddress:   item.IPAddress,                      // IP地址
			Browser:     item.Browser,                        // 浏览器
			Os:          item.Os,                             // 操作系统
			ErrorMsg:    item.ErrorMsg,                       // 登录失败信息
			LoginTime:   time_util.TimeToStr(item.LoginTime), // 登录时间
		})
	}

	return &sysclient.QueryLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
