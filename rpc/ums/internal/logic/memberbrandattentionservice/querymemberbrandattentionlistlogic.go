package memberbrandattentionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

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

	result, err := l.svcCtx.MemberBrandAttentionModel.FindPage(l.ctx, in.MemberId, in.PageNum, in.PageSize)

	if err != nil {
		logc.Errorf(l.ctx, "查询品牌关注列表失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询品牌关注列表失败")
	}

	var list []*umsclient.MemberBrandAttentionListData
	for _, item := range result {

		list = append(list, &umsclient.MemberBrandAttentionListData{
			Id:             item.ID.Hex(),                      //
			MemberId:       item.MemberId,                      // 会员id
			MemberNickName: item.MemberNickName,                // 会员姓名
			MemberIcon:     item.MemberIcon,                    // 会员头像
			BrandId:        item.BrandId,                       // 品牌id
			BrandName:      item.BrandName,                     // 品牌名称
			BrandLogo:      item.BrandLogo,                     // 品牌Logo
			BrandCity:      item.BrandCity,                     // 品牌所在城市
			CreateTime:     time_util.TimeToStr(item.CreateAt), // 关注时间
		})
	}

	return &umsclient.QueryMemberBrandAttentionListResp{
		List: list,
	}, nil

}
