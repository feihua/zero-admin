package memberproductcategoryrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCategoryRelationListLogic 查询会员与产品关糸列表信息
/*
Author: LiuFeiHua
Date: 2024/5/23 13:45
*/
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

// MemberProductCategoryRelationList 查询会员与产品关糸列表信息
func (l *MemberProductCategoryRelationListLogic) MemberProductCategoryRelationList(in *umsclient.MemberProductCategoryRelationListReq) (*umsclient.MemberProductCategoryRelationListResp, error) {
	q := query.UmsMemberProductCategoryRelation.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员与产品关糸列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberProductCategoryRelationListData
	for _, item := range result {

		list = append(list, &umsclient.MemberProductCategoryRelationListData{
			Id:                item.ID,
			MemberId:          item.MemberID,
			ProductCategoryId: item.ProductCategoryID,
		})
	}

	logc.Infof(l.ctx, "查询会员与产品关糸列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberProductCategoryRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
