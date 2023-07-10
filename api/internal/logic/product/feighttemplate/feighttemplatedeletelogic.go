package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.FeightTemplateService.FeightTemplateDelete(l.ctx, &pmsclient.FeightTemplateDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除运费模板异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除运费模板失败")
	}

	return &types.DeleteFeightTemplateResp{
		Code:    "000000",
		Message: "",
	}, nil
}
