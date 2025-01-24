package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberMemberTagRelationListLogic 查询用户和标签关系表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:14
*/
type QueryMemberMemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberMemberTagRelationListLogic {
	return &QueryMemberMemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberMemberTagRelationList 查询用户和标签关系表列表
func (l *QueryMemberMemberTagRelationListLogic) QueryMemberMemberTagRelationList(in *umsclient.QueryMemberMemberTagRelationListReq) (*umsclient.QueryMemberMemberTagRelationListResp, error) {
	result, err := query.UmsMemberMemberTagRelation.WithContext(l.ctx).Where(query.UmsMemberProductCategoryRelation.MemberID.Eq(in.MemberId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户和标签关系列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberMemberTagRelationListData
	for _, item := range result {
		list = append(list, &umsclient.MemberMemberTagRelationListData{
			Id:       item.ID,       //
			MemberId: item.MemberID, // 会员id
			TagId:    item.TagID,    // 标签id
		})
	}

	logc.Infof(l.ctx, "查询用户和标签关系表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.QueryMemberMemberTagRelationListResp{
		Total: 0,
		List:  list,
	}, nil

}
