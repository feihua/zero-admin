package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptAddLogic) DeptAdd(in *sys.DeptAddReq) (*sys.DeptAddResp, error) {
	_, err := l.svcCtx.DeptModel.Insert(model.SysDept{
		Name:       in.Name,
		ParentId:   in.ParentId,
		OrderNum:   in.OrderNum,
		CreateBy:   in.CreateBy,
		CreateTime: time.Time{},
		DelFlag:    0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DeptAddResp{}, nil
}
