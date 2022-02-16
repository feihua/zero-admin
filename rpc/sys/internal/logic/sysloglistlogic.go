package logic

import (
	"context"
	"fmt"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogListLogic) SysLogList(in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	all, err := l.svcCtx.SysLogModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SysLogModel.Count()

	if err != nil {
		return nil, err
	}
	var list []*sys.SysLogListData
	for _, log := range *all {
		fmt.Println(log)
		list = append(list, &sys.SysLogListData{
			Id:             log.Id,
			UserName:       log.UserName,
			Operation:      log.Operation,
			Method:         log.Method,
			Params:         log.Params,
			Time:           log.Time,
			Ip:             log.Ip,
			CreateBy:       log.CreateBy,
			CreateTime:     log.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   log.LastUpdateBy,
			LastUpdateTime: log.LastUpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &sys.SysLogListResp{
		Total: count,
		List:  list,
	}, nil

}
