package loginlogservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogList 登录日志列表
func (l *LoginLogListLogic) LoginLogList(in *sysclient.LoginLogListReq) (*sysclient.LoginLogListResp, error) {
	all, err := l.svcCtx.LoginLogModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.LoginLogModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询登录记录列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sysclient.LoginLogListData
	for _, log := range *all {
		list = append(list, &sysclient.LoginLogListData{
			Id:         log.Id,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.Ip,
			CreateBy:   log.CreateBy,
			CreateTime: log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询登录记录列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
