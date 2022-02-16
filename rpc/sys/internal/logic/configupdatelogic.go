package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *ConfigUpdateLogic) ConfigUpdate(in *sys.ConfigUpdateReq) (*sys.ConfigUpdateResp, error) {
	err := l.svcCtx.ConfigModel.Update(sysmodel.SysConfig{
		Id:             in.Id,
		Value:          in.Value,
		Label:          in.Label,
		Type:           in.Type,
		Description:    in.Description,
		Sort:           float64(in.Sort),
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
		Remarks:        in.Remarks,
	})
	if err != nil {
		return nil, err
	}

	return &sys.ConfigUpdateResp{}, nil
}
