package memberbrandattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberBrandAttentionListLogic 查询会员关注品牌列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:20
*/
type QueryMemberBrandAttentionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberBrandAttentionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberBrandAttentionListLogic {
	return &QueryMemberBrandAttentionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberBrandAttentionList 查询会员关注品牌列表
func (l *QueryMemberBrandAttentionListLogic) QueryMemberBrandAttentionList(in *umsclient.QueryMemberBrandAttentionListReq) (*umsclient.QueryMemberBrandAttentionListResp, error) {
	q := query.UmsMemberBrandAttention.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberBrandAttention.MemberID.Eq(in.MemberId))
	}

	result, err := q.Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询品牌关注列表失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberBrandAttentionListData
	for _, item := range result {

		list = append(list, &umsclient.MemberBrandAttentionListData{
			Id:         item.ID,
			BrandId:    item.BrandID,
			BrandName:  item.BrandName,
			BrandLogo:  item.BrandLogo,
			BrandCity:  *item.BrandCity,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &umsclient.QueryMemberBrandAttentionListResp{
		Total: count,
		List:  list,
	}, nil

}
