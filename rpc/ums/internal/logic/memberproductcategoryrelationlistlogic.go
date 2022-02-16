package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationListLogic {
	return &MemberProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationListLogic) MemberProductCategoryRelationList(in *ums.MemberProductCategoryRelationListReq) (*ums.MemberProductCategoryRelationListResp, error) {
	all, err := l.svcCtx.UmsMemberProductCategoryRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberProductCategoryRelationModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员与产品关糸列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.MemberProductCategoryRelationListData
	for _, item := range *all {

		list = append(list, &ums.MemberProductCategoryRelationListData{
			Id:                item.Id,
			MemberId:          item.MemberId,
			ProductCategoryId: item.ProductCategoryId,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员与产品关糸列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberProductCategoryRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
