package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentUpdateLogic {
	return ProductCommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentUpdateLogic) ProductCommentUpdate(req types.UpdateProductCommentReq) (*types.UpdateProductCommentResp, error) {
	_, err := l.svcCtx.Pms.CommentUpdate(l.ctx, &pmsclient.CommentUpdateReq{
		Id:               req.Id,
		ProductId:        req.ProductId,
		MemberNickName:   req.MemberNickName,
		ProductName:      req.ProductName,
		Star:             req.Star,
		MemberIp:         req.MemberIp,
		CreateTime:       req.CreateTime,
		ShowStatus:       req.ShowStatus,
		ProductAttribute: req.ProductAttribute,
		CollectCouont:    req.CollectCouont,
		ReadCount:        req.ReadCount,
		Content:          req.Content,
		Pics:             req.Pics,
		MemberIcon:       req.MemberIcon,
		ReplayCount:      req.ReplayCount,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新商品评价信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新商品评价失败")
	}

	return &types.UpdateProductCommentResp{
		Code:    "000000",
		Message: "更新商品评价成功",
	}, nil
}
