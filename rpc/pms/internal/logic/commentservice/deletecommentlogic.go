package commentservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCommentLogic 删除商品评价
/*
Author: LiuFeiHua
Date: 2024/6/12 16:36
*/
type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteComment 删除商品评价
func (l *DeleteCommentLogic) DeleteComment(in *pmsclient.DeleteCommentReq) (*pmsclient.DeleteCommentResp, error) {
	_, err := l.svcCtx.ProductCommentModel.Delete(l.ctx, in.Id)

	if err != nil {
		logc.Errorf(l.ctx, "删除商品评价失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品评价失败")
	}

	return &pmsclient.DeleteCommentResp{}, nil
}
