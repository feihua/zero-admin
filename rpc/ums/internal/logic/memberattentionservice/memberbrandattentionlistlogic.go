package memberattentionservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	attentionList, count, err := l.svcCtx.UmsMemberBrandAttentionModel.FindAll(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询品牌关注列表失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberBrandAttentionListData
	for _, item := range *attentionList {

		list = append(list, &umsclient.MemberBrandAttentionListData{
			Id:             item.Id,
			MemberId:       item.MemberId,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			BrandId:        item.BrandId,
			BrandName:      item.BrandName,
			BrandLogo:      item.BrandLogo,
			BrandCity:      item.BrandCity.String,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logc.Infof(l.ctx, "查询商品类别列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberBrandAttentionListResp{
		Total: count,
		List:  list,
	}, nil

}
