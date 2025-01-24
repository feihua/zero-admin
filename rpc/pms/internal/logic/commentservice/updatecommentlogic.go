package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
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
		ID:               in.Id,               //
		ProductID:        in.ProductId,        // 商品id
		MemberNickName:   in.MemberNickName,   // 评价者昵称
		ProductName:      in.ProductName,      // 商品名称
		Star:             in.Star,             // 评价星数：0->5
		MemberIP:         in.MemberIp,         // 评价的ip
		ShowStatus:       in.ShowStatus,       // 是否显示，0-不显示，1-显示
		ProductAttribute: in.ProductAttribute, // 购买时的商品属性
		CollectCount:     in.CollectCount,     // 点赞数
		ReadCount:        in.ReadCount,        // 阅读数
		Content:          in.Content,          // 内容
		Pics:             in.Pics,             // 上传图片地址，以逗号隔开
		MemberIcon:       in.MemberIcon,       // 评论用户头像
		ReplayCount:      in.ReplayCount,      // 回复数量
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateCommentResp{}, nil
}
