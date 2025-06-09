package memberproductcategoryrelationservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberProductCategoryRelationListLogic 查询会员与产品分类关系（用户喜欢的分类）列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:11
*/
type QueryMemberProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberProductCategoryRelationListLogic {
	return &QueryMemberProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberProductCategoryRelationList 查询会员与产品分类关系（用户喜欢的分类）列表
func (l *QueryMemberProductCategoryRelationListLogic) QueryMemberProductCategoryRelationList(in *umsclient.QueryMemberProductCategoryRelationListReq) (*umsclient.QueryMemberProductCategoryRelationListResp, error) {
	result, err := l.svcCtx.MemberProductCategoryRelationModel.FindPage(l.ctx, in.MemberId)

	if err != nil {
		logc.Errorf(l.ctx, "查询会员与产品分类关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员与产品分类关系失败")
	}

	var list []*umsclient.MemberProductCategoryRelationListData
	for _, item := range result {

		list = append(list, &umsclient.MemberProductCategoryRelationListData{
			Id:                item.ID.Hex(),          //
			MemberId:          item.MemberId,          // 会员id
			ProductCategoryId: item.ProductCategoryId, // 商品分类id
		})
	}

	return &umsclient.QueryMemberProductCategoryRelationListResp{
		List: list,
	}, nil

}
