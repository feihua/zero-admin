package member_tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagStatusLogic 更新会员标签状态
type UpdateMemberTagStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateMemberTagStatus 更新会员标签状态
func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(req *types.UpdateMemberTagStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberTagService.UpdateMemberTagStatus(l.ctx, &umsclient.UpdateMemberTagStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员标签状态信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新会员标签状态失败")
	}

	return res.Success()

}
