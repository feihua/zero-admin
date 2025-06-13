package commentreplayservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCommentReplayLogic 更新产品评价回复
/*
Author: LiuFeiHua
Date: 2024/6/12 16:35
*/
type UpdateCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReplayLogic {
	return &UpdateCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCommentReplay 更新产品评价回复
func (l *UpdateCommentReplayLogic) UpdateCommentReplay(in *pmsclient.UpdateCommentReplayReq) (*pmsclient.UpdateCommentReplayResp, error) {
	objectID, _ := primitive.ObjectIDFromHex(in.Id)
	_, err := l.svcCtx.ProductCommentReplayModel.Update(l.ctx, &model.ProductCommentReplay{
		ID:             objectID,          //
		CommentId:      in.CommentId,      // 评论id
		MemberNickName: in.MemberNickName, // 评论人员昵称
		MemberIcon:     in.MemberIcon,     // 评论人员头像
		Content:        in.Content,        // 内容
		Type:           in.Type,           // 评论人员类型；0->会员；1->管理员
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新产品评价回复失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新产品评价回复失败")
	}

	return &pmsclient.UpdateCommentReplayResp{}, nil
}
