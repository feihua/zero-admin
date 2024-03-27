package configservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/sysmodel"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConfigUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 16:58
*/
type ConfigUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigUpdateLogic {
	return &ConfigUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ConfigUpdate 更新配置信息
func (l *ConfigUpdateLogic) ConfigUpdate(in *sysclient.ConfigUpdateReq) (*sysclient.ConfigUpdateResp, error) {
	err := l.svcCtx.ConfigModel.Update(l.ctx, &sysmodel.SysConfig{
		Id:          in.Id,
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		UpdateBy:    sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime:  sql.NullTime{Time: time.Now()},
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.ConfigUpdateResp{}, nil
}
