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

// MemberProductCollectionAddLogic 收藏相关
/*
Author: LiuFeiHua
Date: 2023/11/30 11:52
*/
type MemberProductCollectionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionAddLogic {
	return &MemberProductCollectionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCollectionAdd 添加收藏商品
func (l *MemberProductCollectionAddLogic) MemberProductCollectionAdd(in *umsclient.MemberProductCollectionAddReq) (*umsclient.MemberProductCollectionAddResp, error) {
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

	return &umsclient.MemberProductCollectionAddResp{}, nil
}
