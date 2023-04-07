package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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

func (l *MemberMemberTagRelationListLogic) MemberMemberTagRelationList(in *umsclient.MemberMemberTagRelationListReq) (*umsclient.MemberMemberTagRelationListResp, error) {
	// todo: add your logic here and delete this line
	// all, err := l.svcCtx.UmsMemberMemberTagRelationModel.FindAll(l.ctx, in.Current, in.PageSize)
	// count, _ := l.svcCtx.UmsMemberMemberTagRelationModel.Count(l.ctx)

	// if err != nil {
	// 	reqStr, _ := json.Marshal(in)
	// 	logx.WithContext(l.ctx).Errorf("查询用户和标签关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
	// 	return nil, err
	// }

	// var list []*ums.MemberMemberTagRelationListData
	// for _, item := range *all {

	// 	list = append(list, &ums.MemberMemberTagRelationListData{
	// 		Id:       item.Id,
	// 		MemberId: item.MemberId,
	// 		TagId:    item.TagId,
	// 	})
	// }

	// reqStr, _ := json.Marshal(in)
	// listStr, _ := json.Marshal(list)
	// logx.WithContext(l.ctx).Infof("查询用户和标签关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberMemberTagRelationListResp{
		// Total: count,
		// List:  list,
	}, nil
}
