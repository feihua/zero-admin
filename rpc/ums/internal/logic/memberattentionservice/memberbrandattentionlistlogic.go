package memberattentionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberBrandAttentionListLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 16:42
*/
type MemberBrandAttentionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewMemberBrandAttentionListLogic 查询品牌关注列表
func NewMemberBrandAttentionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberBrandAttentionListLogic {
	return &MemberBrandAttentionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberBrandAttentionListLogic) MemberBrandAttentionList(in *umsclient.MemberBrandAttentionListReq) (*umsclient.MemberBrandAttentionListResp, error) {
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
			Id:             item.ID,
			MemberId:       item.MemberID,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			BrandId:        item.BrandID,
			BrandName:      item.BrandName,
			BrandLogo:      item.BrandLogo,
			BrandCity:      *item.BrandCity,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询商品类别列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberBrandAttentionListResp{
		Total: count,
		List:  list,
	}, nil

}
