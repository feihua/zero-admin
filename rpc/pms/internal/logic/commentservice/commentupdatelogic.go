package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CommentUpdateLogic 商品评价
/*
Author: LiuFeiHua
Date: 2024/5/8 10:43
*/
type CommentUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentUpdateLogic {
	return &CommentUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentUpdate 更新商品评价
func (l *CommentUpdateLogic) CommentUpdate(in *pmsclient.CommentUpdateReq) (*pmsclient.CommentUpdateResp, error) {
	q := query.PmsComment
	_, err := q.WithContext(l.ctx).Updates(&model.PmsComment{
		ID:               in.Id,
		ProductID:        in.ProductId,
		MemberNickName:   in.MemberNickName,
		ProductName:      in.ProductName,
		Star:             in.Star,
		MemberIP:         in.MemberIp,
		CreateTime:       time.Now(),
		ShowStatus:       in.ShowStatus,
		ProductAttribute: in.ProductAttribute,
		CollectCouont:    in.CollectCouont,
		ReadCount:        in.ReadCount,
		Content:          in.Content,
		Pics:             in.Pics,
		MemberIcon:       in.MemberIcon,
		ReplayCount:      in.ReplayCount,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.CommentUpdateResp{}, nil
}
