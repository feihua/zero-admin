package membertagrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberTagRelationDetailLogic 查询会员标签关联详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagRelationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagRelationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagRelationDetailLogic {
	return &QueryMemberTagRelationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTagRelationDetail 查询会员标签关联详情
func (l *QueryMemberTagRelationDetailLogic) QueryMemberTagRelationDetail(in *umsclient.QueryMemberTagRelationDetailReq) (*umsclient.QueryMemberTagRelationDetailResp, error) {
	item, err := query.UmsMemberTagRelation.WithContext(l.ctx).Where(query.UmsMemberTagRelation.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员标签关联不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员标签关联不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员标签关联异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员标签关联异常")
	}

	data := &umsclient.QueryMemberTagRelationDetailResp{
		Id:         item.ID,                              // 主键ID
		MemberId:   item.MemberID,                        // 会员ID
		TagId:      item.TagID,                           // 标签ID
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
	}

	return data, nil
}
