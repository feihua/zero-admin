package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberMemberTagRelationListLogic) MemberMemberTagRelationList(in *ums.MemberMemberTagRelationListReq) (*ums.MemberMemberTagRelationListResp, error) {
	all, err := l.svcCtx.UmsMemberMemberTagRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberMemberTagRelationModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询用户和标签关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.MemberMemberTagRelationListData
	for _, item := range *all {

		list = append(list, &ums.MemberMemberTagRelationListData{
			Id:       item.Id,
			MemberId: item.MemberId,
			TagId:    item.TagId,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询用户和标签关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberMemberTagRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
