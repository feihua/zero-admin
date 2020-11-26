package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(model.SysMenu{
		Name:       in.Name,
		ParentId:   in.ParentId,
		Url:        in.Url,
		Perms:      in.Perms,
		Type:       in.Type,
		Icon:       in.Icon,
		OrderNum:   in.OrderNum,
		CreateBy:   in.CreateBy,
		CreateTime: time.Time{},
		DelFlag:    0,
	})
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
		return nil, err
	}

	return &sys.MenuAddResp{}, nil
}
