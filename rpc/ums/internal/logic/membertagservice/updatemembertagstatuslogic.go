package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagStatusLogic 更新用户标签
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
*/
type UpdateMemberTagStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTagStatus 更新用户标签状态
func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(in *umsclient.UpdateMemberTagStatusReq) (*umsclient.UpdateMemberTagStatusResp, error) {
	q := query.UmsMemberTag

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户标签状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新用户标签状态失败")
	}

	logc.Infof(l.ctx, "更新用户标签状态成功,参数：%+v", in)
	return &umsclient.UpdateMemberTagStatusResp{}, nil
}
