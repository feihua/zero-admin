package flashpromotionproductrelation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionProductRelationDeleteLogic 限时购和商品关系
/*
Author: LiuFeiHua
Date: 2024/5/14 15:30
*/
type FlashPromotionProductRelationDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationDeleteLogic {
	return &FlashPromotionProductRelationDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionProductRelationDelete 删除限时购和商品关系
func (l *FlashPromotionProductRelationDeleteLogic) FlashPromotionProductRelationDelete(req *types.DeleteFlashPromotionProductRelationReq) (resp *types.DeleteFlashPromotionProductRelationResp, err error) {
	_, err = l.svcCtx.FlashPromotionProductRelationService.FlashPromotionProductRelationDelete(l.ctx, &smsclient.FlashPromotionProductRelationDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %v,删除限时购和商品关系异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除限时购和商品关系失败")
	}
	return &types.DeleteFlashPromotionProductRelationResp{
		Code:    "000000",
		Message: "删除限时购和商品关系成功",
	}, nil
}
