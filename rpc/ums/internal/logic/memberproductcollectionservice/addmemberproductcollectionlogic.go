package memberproductcollectionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberProductCollectionLogic 添加用户收藏的商品
/*
Author: LiuFeiHua
Date: 2024/6/11 14:08
*/
type AddMemberProductCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberProductCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberProductCollectionLogic {
	return &AddMemberProductCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberProductCollection 添加用户收藏的商品
func (l *AddMemberProductCollectionLogic) AddMemberProductCollection(in *umsclient.AddMemberProductCollectionReq) (*umsclient.AddMemberProductCollectionResp, error) {
	err := query.UmsMemberProductCollection.WithContext(l.ctx).Create(&model.UmsMemberProductCollection{
		MemberID:        in.MemberId,        // 会员id
		MemberNickName:  in.MemberNickName,  // 会员姓名
		MemberIcon:      in.MemberIcon,      // 会员头像
		ProductID:       in.ProductId,       // 商品id
		ProductName:     in.ProductName,     // 商品名称
		ProductPic:      in.ProductPic,      // 商品图片
		ProductSubTitle: in.ProductSubTitle, // 商品标题
		ProductPrice:    in.ProductPrice,    // 商品价格
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加用户收藏的商品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加用户收藏的商品失败")
	}

	return &umsclient.AddMemberProductCollectionResp{}, nil
}
