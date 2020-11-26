package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(model.SysMenu{
		Id:             0,
		Name:           in.Name,
		ParentId:       in.ParentId,
		Url:            in.Url,
		Perms:          in.Perms,
		Type:           in.Type,
		Icon:           in.Icon,
		OrderNum:       in.OrderNum,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Time{},
		DelFlag:        0,
	})
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
		return nil, err
	}

	return &sys.MenuUpdateResp{}, nil
}
