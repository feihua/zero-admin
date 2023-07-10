package deptservicelogic

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

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

func (l *DeptUpdateLogic) DeptUpdate(in *sysclient.DeptUpdateReq) (*sysclient.DeptUpdateResp, error) {
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
		DelFlag:    in.DelFlag,
		ParentIds:  strings.Replace(strings.Trim(fmt.Sprint(in.ParentIds), "[]"), " ", ",", -1),
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.DeptUpdateResp{}, nil
}
