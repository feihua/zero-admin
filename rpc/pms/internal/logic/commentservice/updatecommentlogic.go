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

// UpdateCommentLogic 更新商品评价表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:38
*/
type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateComment 更新商品评价表
func (l *UpdateCommentLogic) UpdateComment(in *pmsclient.UpdateCommentReq) (*pmsclient.UpdateCommentResp, error) {
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

	return &pmsclient.UpdateCommentResp{}, nil
}
