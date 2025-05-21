package memberinfoservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberPointsLogic 更新会员积分
/*
Author: LiuFeiHua
Date: 2025/5/21 14:57
*/
type UpdateMemberPointsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberPointsLogic {
	return &UpdateMemberPointsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberPoints 更新会员积分
func (l *UpdateMemberPointsLogic) UpdateMemberPoints(in *umsclient.UpdateMemberPointsReq) (*umsclient.UpdateMemberPointsResp, error) {
	q := query.UmsMemberInfo

	_, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Update(q.Points, in.Points)
	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员积分失败")
	}

	return &umsclient.UpdateMemberPointsResp{}, nil
}
