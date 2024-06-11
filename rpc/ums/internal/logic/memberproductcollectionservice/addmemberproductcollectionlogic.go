package memberproductcollectionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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
		MemberID:        in.MemberId,
		MemberNickName:  in.MemberNickName,
		MemberIcon:      in.MemberIcon,
		ProductID:       in.ProductId,
		ProductName:     in.ProductName,
		ProductPic:      in.ProductPic,
		ProductSubTitle: &in.ProductSubTitle,
		ProductPrice:    in.ProductPrice,
		CreateTime:      time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberProductCollectionResp{}, nil
}
