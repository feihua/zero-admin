package configservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"

	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigAddLogic {
	return &ConfigAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigAddLogic) ConfigAdd(in *sysclient.ConfigAddReq) (*sysclient.ConfigAddResp, error) {
	_, err := l.svcCtx.ConfigModel.Insert(l.ctx, &sysmodel.SysConfig{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    sql.NullString{String: in.CreateBy, Valid: true},
		CreateTime:  time.Now(),
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     0,
	})
	if err != nil {
		return nil, err
	}

	return &sysclient.ConfigAddResp{}, nil
}
