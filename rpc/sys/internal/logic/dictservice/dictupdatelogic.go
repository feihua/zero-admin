package dictservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/sysmodel"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DictUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:03
*/
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

// DictUpdate 更新字典信息
func (l *DictUpdateLogic) DictUpdate(in *sysclient.DictUpdateReq) (*sysclient.DictUpdateResp, error) {
	//更新之前查询记录是否存在
	dict, err := l.svcCtx.DictModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.DictModel.Update(l.ctx, &sysmodel.SysDict{
		Id:          in.Id,
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    dict.CreateBy,
		CreateTime:  dict.CreateTime,
		UpdateBy:    sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime:  time.Now(),
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     in.DelFlag,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.DictUpdateResp{}, nil
}
