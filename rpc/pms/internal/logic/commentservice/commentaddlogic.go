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

// CommentAddLogic 商品评价
/*
Author: LiuFeiHua
Date: 2024/5/8 10:42
*/
type CommentAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentAddLogic {
	return &CommentAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentAdd 添加商品评价
func (l *CommentAddLogic) CommentAdd(in *pmsclient.CommentAddReq) (*pmsclient.CommentAddResp, error) {
	err := query.PmsComment.WithContext(l.ctx).Create(&model.PmsComment{
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

	return &pmsclient.CommentAddResp{}, nil
}
