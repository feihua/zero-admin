package memberreadhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"time"

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

	return &umsclient.AddMemberReadHistoryResp{}, nil
}
