package membermembertagrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberMemberTagRelationListLogic 用户和标签关系
/*
Author: LiuFeiHua
Date: 2024/5/7 11:06
*/
type MemberMemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationListLogic {
	return &MemberMemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberMemberTagRelationList 查询用户和标签关系
func (l *MemberMemberTagRelationListLogic) MemberMemberTagRelationList(in *umsclient.MemberMemberTagRelationListReq) (*umsclient.MemberMemberTagRelationListResp, error) {
	q := query.UmsMemberMemberTagRelation.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户和标签关系列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberMemberTagRelationListData
	for _, item := range result {
		list = append(list, &umsclient.MemberMemberTagRelationListData{
			Id:       item.ID,
			MemberId: item.MemberID,
			TagId:    item.TagID,
		})
	}

	logc.Infof(l.ctx, "查询用户和标签关系表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberMemberTagRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
