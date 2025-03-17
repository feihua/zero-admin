package prefrence_area

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePreferredAreaStatusLogic 更新优选专区状态状态
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type UpdatePreferredAreaStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePreferredAreaStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaStatusLogic {
	return &UpdatePreferredAreaStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePreferredAreaStatus 更新优选专区状态
func (l *UpdatePreferredAreaStatusLogic) UpdatePreferredAreaStatus(req *types.UpdatePreferredAreaStatusReq) (resp *types.UpdatePreferredAreaStatusResp, err error) {
	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.PreferredAreaService.UpdatePreferredAreaStatus(l.ctx, &cmsclient.UpdatePreferredAreaStatusReq{
		Ids:        req.Ids,        // 主键ID
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   updateBy,       // 更新者
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优选专区状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdatePreferredAreaStatusResp{
		Code:    "000000",
		Message: "更新优选专区状态成功",
	}, nil
}
