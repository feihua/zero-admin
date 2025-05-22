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
)

// QueryMemberTagRelationListLogic 查询会员标签关联列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagRelationListLogic {
	return &QueryMemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTagRelationList 查询会员标签关联列表
func (l *QueryMemberTagRelationListLogic) QueryMemberTagRelationList(in *umsclient.QueryMemberTagRelationListReq) (*umsclient.QueryMemberTagRelationListResp, error) {
	memberTagRelation := query.UmsMemberTagRelation
	q := memberTagRelation.WithContext(l.ctx)
	if in.MemberId != 2 {
		q = q.Where(memberTagRelation.MemberID.Eq(in.MemberId))
	}
	if in.TagId != 2 {
		q = q.Where(memberTagRelation.TagID.Eq(in.TagId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员标签关联列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员标签关联列表失败")
	}

	var list []*umsclient.MemberTagRelationListData

	for _, item := range result {
		list = append(list, &umsclient.MemberTagRelationListData{
			Id:         item.ID,                              // 主键ID
			MemberId:   item.MemberID,                        // 会员ID
			TagId:      item.TagID,                           // 标签ID
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间

		})
	}

	return &umsclient.QueryMemberTagRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
