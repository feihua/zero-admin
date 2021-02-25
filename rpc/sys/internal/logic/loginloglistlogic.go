package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogListLogic) LoginLogList(in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	all, err := l.svcCtx.LoginLogModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.LoginLogModel.Count()

	if err != nil {
		return nil, err
	}

	var list []*sys.LoginLogListData
	for _, log := range *all {
		fmt.Println(log)
		list = append(list, &sys.LoginLogListData{
			Id:             log.Id,
			UserName:       log.UserName,
			Status:         log.Status,
			Ip:             log.Ip,
			CreateBy:       log.CreateBy,
			CreateTime:     log.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   log.LastUpdateBy,
			LastUpdateTime: log.LastUpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &sys.LoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
