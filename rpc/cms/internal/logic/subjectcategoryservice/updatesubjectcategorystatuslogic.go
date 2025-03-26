package subjectcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSubjectCategoryStatusLogic 更新专题分类
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateSubjectCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCategoryStatusLogic {
	return &UpdateSubjectCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSubjectCategoryStatus 更新专题分类状态
func (l *UpdateSubjectCategoryStatusLogic) UpdateSubjectCategoryStatus(in *cmsclient.UpdateSubjectCategoryStatusReq) (*cmsclient.UpdateSubjectCategoryStatusResp, error) {
	q := query.CmsSubjectCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题分类状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新专题分类状态失败")
	}

	return &cmsclient.UpdateSubjectCategoryStatusResp{}, nil
}
