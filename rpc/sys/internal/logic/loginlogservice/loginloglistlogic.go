package loginlogservicelogic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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
		fmt.Println(log)
		list = append(list, &sysclient.LoginLogListData{
			Id:         log.Id,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.Ip,
			CreateBy:   log.CreateBy,
			CreateTime: log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询登录记录列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sysclient.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
