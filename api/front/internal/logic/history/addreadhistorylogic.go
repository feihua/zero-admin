package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddReadHistoryLogic 添加浏览商品记录
/*
Author: LiuFeiHua
Date: 2024/5/16 10:43
*/
type AddReadHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReadHistoryLogic {
	return &AddReadHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddReadHistory 添加浏览商品记录
func (l *AddReadHistoryLogic) AddReadHistory(req *types.AddReadHistoryReq) (resp *types.AddReadHistoryResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	member, _ := l.svcCtx.MemberService.QueryMemberDetail(l.ctx, &umsclient.QueryMemberDetailReq{Id: memberId})

	_, err = l.svcCtx.MemberReadHistoryService.AddMemberReadHistory(l.ctx, &umsclient.AddMemberReadHistoryReq{
		MemberId:        member.Id,
		MemberNickName:  member.Nickname,
		MemberIcon:      member.Icon,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		ProductPic:      req.ProductPic,
		ProductSubTitle: req.ProductSubTitle,
		ProductPrice:    req.ProductPrice,
	})

	if err != nil {
		return nil, err
	}
	return &types.AddReadHistoryResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
