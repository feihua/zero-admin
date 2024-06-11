package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagLogic 更新用户标签表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:46
*/
type UpdateMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagLogic {
	return &UpdateMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTag 更新用户标签表
func (l *UpdateMemberTagLogic) UpdateMemberTag(in *umsclient.UpdateMemberTagReq) (*umsclient.UpdateMemberTagResp, error) {
	_, err := query.UmsMemberTag.WithContext(l.ctx).Updates(&model.UmsMemberTag{
		ID:                in.Id,
		TagName:           in.TagName,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: in.FinishOrderAmount,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberTagResp{}, nil
}
