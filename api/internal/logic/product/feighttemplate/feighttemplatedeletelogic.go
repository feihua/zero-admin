package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FeightTemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeightTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeightTemplateDeleteLogic {
	return FeightTemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeightTemplateDeleteLogic) FeightTemplateDelete(req types.DeleteFeightTemplateReq) (*types.DeleteFeightTemplateResp, error) {
	_, err := l.svcCtx.Pms.FeightTemplateDelete(l.ctx, &pmsclient.FeightTemplateDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除运费模板异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除运费模板失败")
	}

	return &types.DeleteFeightTemplateResp{
		Code:    "000000",
		Message: "",
	}, nil
}
