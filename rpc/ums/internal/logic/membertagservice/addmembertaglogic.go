package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberTagLogic 添加用户标签表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:45
*/
type AddMemberTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberTagLogic {
	return &AddMemberTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberTag 添加用户标签表
func (l *AddMemberTagLogic) AddMemberTag(in *umsclient.AddMemberTagReq) (*umsclient.AddMemberTagResp, error) {
	err := query.UmsMemberTag.WithContext(l.ctx).Create(&model.UmsMemberTag{
		TagName:           in.TagName,
		FinishOrderCount:  in.FinishOrderCount,
		Status:            in.Status,
		FinishOrderAmount: in.FinishOrderAmount,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.AddMemberTagResp{}, nil
}
