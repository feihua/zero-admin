package feighttemplateservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFeightTemplateLogic 删除运费模版
/*
Author: LiuFeiHua
Date: 2024/6/12 16:39
*/
type DeleteFeightTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFeightTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFeightTemplateLogic {
	return &DeleteFeightTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFeightTemplate 删除运费模版
func (l *DeleteFeightTemplateLogic) DeleteFeightTemplate(in *pmsclient.DeleteFeightTemplateReq) (*pmsclient.DeleteFeightTemplateResp, error) {
	q := query.PmsFeightTemplate
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除运费模版失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除运费模版失败")
	}

	return &pmsclient.DeleteFeightTemplateResp{}, nil
}
