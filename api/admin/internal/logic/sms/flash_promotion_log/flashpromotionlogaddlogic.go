package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionLogAddLogic {
	return FlashPromotionLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionLogAddLogic) FlashPromotionLogAdd(req types.AddFlashPromotionLogReq) (*types.AddFlashPromotionLogResp, error) {
	_, err := l.svcCtx.FlashPromotionLogService.AddFlashPromotionLog(l.ctx, &smsclient.AddFlashPromotionLogReq{
		MemberId:      req.MemberId,      // 会员id
		ProductId:     req.ProductId,     // 商品id
		MemberPhone:   req.MemberPhone,   // 会员电话
		ProductName:   req.ProductName,   // 商品名称
		SubscribeTime: req.SubscribeTime, // 会员订阅时间
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购通知记录信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddFlashPromotionLogResp{
		Code:    "000000",
		Message: "添加限时购通知记录成功",
	}, nil
}
