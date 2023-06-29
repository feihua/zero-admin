package logic

import (
	"context"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"
	"zero-admin/rpc/model/cmsmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationAddLogic {
	return &PrefrenceAreaProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PrefrenceAreaProductRelationAdd 优选商品关联
func (l *PrefrenceAreaProductRelationAddLogic) PrefrenceAreaProductRelationAdd(in *cmsclient.PrefrenceAreaProductRelationAddReq) (*cmsclient.PrefrenceAreaProductRelationAddResp, error) {
	_, err := l.svcCtx.CmsPrefrenceAreaProductRelationModel.Insert(l.ctx, &cmsmodel.CmsPrefrenceAreaProductRelation{
		PrefrenceAreaId: in.PrefrenceAreaId,
		ProductId:       in.ProductId,
	})
	if err != nil {
		return nil, err
	}

	return &cmsclient.PrefrenceAreaProductRelationAddResp{}, nil
}
