package memberreadhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberReadHistoryLogic 添加用户商品浏览历史记录
/*
Author: LiuFeiHua
Date: 2024/6/11 14:06
*/
type AddMemberReadHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberReadHistoryLogic {
	return &AddMemberReadHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberReadHistory 添加用户商品浏览历史记录
func (l *AddMemberReadHistoryLogic) AddMemberReadHistory(in *umsclient.AddMemberReadHistoryReq) (*umsclient.AddMemberReadHistoryResp, error) {
	err := query.UmsMemberReadHistory.WithContext(l.ctx).Create(&model.UmsMemberReadHistory{
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
		return nil, err
	}

	return &umsclient.AddMemberReadHistoryResp{}, nil
}
