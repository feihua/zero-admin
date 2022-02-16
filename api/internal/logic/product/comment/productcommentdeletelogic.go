package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Pms.CommentDelete(l.ctx, &pmsclient.CommentDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品评价异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除商品评价失败")
	}

	return &types.DeleteProductCommentResp{
		Code:    "000000",
		Message: "",
	}, nil
}
