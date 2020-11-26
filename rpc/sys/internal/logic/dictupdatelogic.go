package logic

import (
	"context"
	"go-zero-admin/rpc/model"
	"time"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictUpdateLogic {
	return &DictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictUpdateLogic) DictUpdate(in *sys.DictUpdateReq) (*sys.DictUpdateResp, error) {
	_, err := l.svcCtx.DictModel.Insert(model.SysDict{
		Id:             in.Id,
		Value:          in.Value,
		Label:          in.Label,
		Type:           in.Type,
		Description:    in.Description,
		Sort:           float64(in.Sort),
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Time{},
		Remarks:        in.Remarks,
		DelFlag:        0,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DictUpdateResp{}, nil
}
