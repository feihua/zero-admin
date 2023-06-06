package logic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictAddLogic) DictAdd(in *sys.DictAddReq) (*sys.DictAddResp, error) {
	_, err := l.svcCtx.DictModel.Insert(l.ctx, &sysmodel.SysDict{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    in.CreateBy,
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     in.DelFlag,
	})

	if err != nil {
		return nil, err
	}

	return &sys.DictAddResp{}, nil
}
