package loginlogservicelogic

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
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
	if len(in.LoginName) > 0 {
		q = q.Where(loginLog.LoginName.Like("%" + in.LoginName + "%"))
	}
	if len(in.Ipaddr) > 0 {
		q = q.Where(loginLog.IPAddr.Like("%" + in.Ipaddr + "%"))
	}
	if in.Status != 2 {
		q = q.Where(loginLog.Status.Eq(in.Status))
	}

	result, count, err := q.Order(loginLog.ID.Desc()).FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))
	if err != nil {
		logc.Errorf(l.ctx, "查询登录记录列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询登录记录列表信息失败")
	}

	var list []*sysclient.LoginLogListData
	for _, item := range result {
		data := &sysclient.LoginLogListData{
			Id:        item.ID,                             // 登录日志id
			LoginName: item.LoginName,                      // 登录账号
			Ipaddr:    item.IPAddr,                         // 登录IP地址
			Status:    item.Status,                         // 登录状态(0:失败,1:成功)
			Extra:     string(item.Extra),                  // 其他信息（可选）
			Msg:       item.Msg,                            // 提示消息
			LoginTime: time_util.TimeToStr(item.LoginTime), // 登录时间
		}
		_ = sonic.Unmarshal(item.Extra, data.Extra)
		list = append(list, data)
	}

	return &sysclient.QueryLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
