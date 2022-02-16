package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *ConfigAddLogic) ConfigAdd(in *sys.ConfigAddReq) (*sys.ConfigAddResp, error) {
	_, err := l.svcCtx.ConfigModel.Insert(sysmodel.SysConfig{
		Value:          in.Value,
		Label:          in.Label,
		Type:           in.Type,
		Description:    in.Description,
		Sort:           float64(in.Sort),
		CreateBy:       in.CreateBy,
		LastUpdateTime: time.Now(),
		LastUpdateBy:   in.CreateBy,
		Remarks:        in.Remarks,
		DelFlag:        0,
	})
	if err != nil {
		return nil, err
	}

	return &sys.ConfigAddResp{}, nil
}
