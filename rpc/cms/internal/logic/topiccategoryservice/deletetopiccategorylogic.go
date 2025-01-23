package topiccategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteTopicCategoryLogic 删除话题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeleteTopicCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicCategoryLogic {
	return &DeleteTopicCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteTopicCategory 删除话题分类
func (l *DeleteTopicCategoryLogic) DeleteTopicCategory(in *cmsclient.DeleteTopicCategoryReq) (*cmsclient.DeleteTopicCategoryResp, error) {
	q := query.CmsTopicCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除话题分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除话题分类失败")
	}

	logc.Infof(l.ctx, "删除话题分类成功,参数：%+v", in)
	return &cmsclient.DeleteTopicCategoryResp{}, nil
}
