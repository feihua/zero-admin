package history

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	member, err := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员信息失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	_, err = l.svcCtx.MemberReadHistoryService.AddMemberReadHistory(l.ctx, &umsclient.AddMemberReadHistoryReq{
		MemberId:        member.Id,
		MemberNickName:  member.Nickname,
		MemberIcon:      member.Avatar,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		ProductPic:      req.ProductPic,
		ProductSubTitle: req.ProductSubTitle,
		ProductPrice:    req.ProductPrice,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加浏览商品记录失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return &types.AddReadHistoryResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
