package dictservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/sysmodel"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:02
*/
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

// DictAdd 添加字典信息
func (l *DictAddLogic) DictAdd(in *sysclient.DictAddReq) (*sysclient.DictAddResp, error) {
	dict := &sysmodel.SysDict{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    in.CreateBy,
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     in.DelFlag,
	}
	if _, err := l.svcCtx.DictModel.Insert(l.ctx, dict); err != nil {
		return nil, err
	}

	return &sysclient.DictAddResp{}, nil
}
