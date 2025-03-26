package helpcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHelpCategoryLogic 删除帮助分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:23:59
*/
type DeleteHelpCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpCategoryLogic {
	return &DeleteHelpCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHelpCategory 删除帮助分类
func (l *DeleteHelpCategoryLogic) DeleteHelpCategory(in *cmsclient.DeleteHelpCategoryReq) (*cmsclient.DeleteHelpCategoryResp, error) {
	q := query.CmsHelpCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除帮助分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除帮助分类失败")
	}

	return &cmsclient.DeleteHelpCategoryResp{}, nil
}
