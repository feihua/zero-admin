package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentAddLogic {
	return ProductCommentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentAddLogic) ProductCommentAdd(req types.AddProductCommentReq) (*types.AddProductCommentResp, error) {
	_, err := l.svcCtx.CommentService.CommentAdd(l.ctx, &pmsclient.CommentAddReq{
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
		logx.WithContext(l.ctx).Errorf("添加商品评价信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加商品评价失败")
	}

	return &types.AddProductCommentResp{
		Code:    "000000",
		Message: "添加商品评价成功",
	}, nil
}
