package commentreplayservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCommentReplayLogic 添加产品评价回复
/*
Author: LiuFeiHua
Date: 2024/6/12 16:33
*/
type AddCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentReplayLogic {
	return &AddCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCommentReplay 添加产品评价回复
func (l *AddCommentReplayLogic) AddCommentReplay(in *pmsclient.AddCommentReplayReq) (*pmsclient.AddCommentReplayResp, error) {
	err := query.PmsCommentReplay.WithContext(l.ctx).Create(&model.PmsCommentReplay{
		CommentID:      in.CommentId,      // 评论id
		MemberNickName: in.MemberNickName, // 评论人员昵称
		MemberIcon:     in.MemberIcon,     // 评论人员头像
		Content:        in.Content,        // 内容
		Type:           in.Type,           // 评论人员类型；0->会员；1->管理员
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加产品评价回复失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加产品评价回复失败")
	}

	return &pmsclient.AddCommentReplayResp{}, nil
}
