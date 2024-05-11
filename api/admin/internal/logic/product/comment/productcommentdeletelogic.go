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

type ProductCommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentDeleteLogic {
	return ProductCommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentDeleteLogic) ProductCommentDelete(req types.DeleteProductCommentReq) (*types.DeleteProductCommentResp, error) {
	_, err := l.svcCtx.CommentService.CommentDelete(l.ctx, &pmsclient.CommentDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除商品评价异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除商品评价失败")
	}

	return &types.DeleteProductCommentResp{
		Code:    "000000",
		Message: "",
	}, nil
}
