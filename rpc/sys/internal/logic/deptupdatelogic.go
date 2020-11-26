package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *sys.DeptUpdateReq) (*sys.DeptUpdateResp, error) {
	err := l.svcCtx.DeptModel.Update(model.SysDept{
		Id:             in.Id,
		Name:           in.Name,
		ParentId:       in.ParentId,
		OrderNum:       in.OrderNum,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &sys.DeptUpdateResp{}, nil
}
