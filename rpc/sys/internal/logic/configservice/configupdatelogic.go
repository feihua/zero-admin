package configservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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
