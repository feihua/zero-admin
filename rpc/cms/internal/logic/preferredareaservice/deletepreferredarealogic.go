package preferredareaservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeletePreferredAreaLogic 删除优选专区
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeletePreferredAreaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePreferredAreaLogic {
	return &DeletePreferredAreaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePreferredArea 删除优选专区
func (l *DeletePreferredAreaLogic) DeletePreferredArea(in *cmsclient.DeletePreferredAreaReq) (*cmsclient.DeletePreferredAreaResp, error) {
	q := query.CmsPreferredArea

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除优选专区失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优选专区失败")
	}

	logc.Infof(l.ctx, "删除优选专区成功,参数：%+v", in)
	return &cmsclient.DeletePreferredAreaResp{}, nil
}
