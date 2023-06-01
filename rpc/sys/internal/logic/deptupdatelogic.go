package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
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
	dept, err := l.svcCtx.DeptModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DeptModel.Update(l.ctx, &sysmodel.SysDept{
		Id:         in.Id,
		Name:       in.Name,
		ParentId:   in.ParentId,
		OrderNum:   in.OrderNum,
		CreateBy:   dept.CreateBy,
		CreateTime: dept.CreateTime,
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime: sql.NullTime{Time: time.Now()},
		DelFlag:    0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DeptUpdateResp{}, nil
}
