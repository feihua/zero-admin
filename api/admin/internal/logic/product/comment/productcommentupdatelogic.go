package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	_, err := l.svcCtx.CommentService.CommentUpdate(l.ctx, &pmsclient.CommentUpdateReq{
		Id:               req.Id,
		ProductId:        req.ProductId,
		MemberNickName:   req.MemberNickName,
		ProductName:      req.ProductName,
		Star:             req.Star,
		MemberIp:         req.MemberIp,
		CreateTime:       req.CreateTime,
		ShowStatus:       req.ShowStatus,
		ProductAttribute: req.ProductAttribute,
		CollectCount:     req.CollectCount,
		ReadCount:        req.ReadCount,
		Content:          req.Content,
		Pics:             req.Pics,
		MemberIcon:       req.MemberIcon,
		ReplayCount:      req.ReplayCount,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品评价信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新商品评价失败")
	}

	return &types.UpdateProductCommentResp{
		Code:    "000000",
		Message: "更新商品评价成功",
	}, nil
}
