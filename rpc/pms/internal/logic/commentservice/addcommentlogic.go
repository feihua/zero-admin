package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCommentLogic 添加商品评价表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:35
*/
type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddComment 添加商品评价表
func (l *AddCommentLogic) AddComment(in *pmsclient.AddCommentReq) (*pmsclient.AddCommentResp, error) {
	err := query.PmsComment.WithContext(l.ctx).Create(&model.PmsComment{
		ProductID:        in.ProductId,
		MemberNickName:   in.MemberNickName,
		ProductName:      in.ProductName,
		Star:             in.Star,
		MemberIP:         in.MemberIp,
		CreateTime:       time.Now(),
		ShowStatus:       in.ShowStatus,
		ProductAttribute: in.ProductAttribute,
		CollectCount:     in.CollectCount,
		ReadCount:        in.ReadCount,
		Content:          in.Content,
		Pics:             in.Pics,
		MemberIcon:       in.MemberIcon,
		ReplayCount:      in.ReplayCount,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AddCommentResp{}, nil
}
