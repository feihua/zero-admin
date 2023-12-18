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

// ConfigAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 16:47
*/
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

// ConfigAdd 添加配置信息
func (l *ConfigAddLogic) ConfigAdd(in *sysclient.ConfigAddReq) (*sysclient.ConfigAddResp, error) {
	config := &sysmodel.SysConfig{
		Value:       in.Value,
		Label:       in.Label,
		Type:        in.Type,
		Description: in.Description,
		Sort:        float64(in.Sort),
		CreateBy:    sql.NullString{String: in.CreateBy, Valid: true},
		CreateTime:  time.Now(),
		Remarks:     sql.NullString{String: in.Remarks, Valid: true},
		DelFlag:     0,
	}
	_, err := l.svcCtx.ConfigModel.Insert(l.ctx, config)

	if err != nil {
		return nil, err
	}

	return &sysclient.ConfigAddResp{}, nil
}
