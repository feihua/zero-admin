package prefrence_area

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePreferredAreaLogic 更新优选专区
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type UpdatePreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaLogic {
	return &UpdatePreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePreferredArea 更新优选专区
func (l *UpdatePreferredAreaLogic) UpdatePreferredArea(req *types.UpdatePreferredAreaReq) (resp *types.BaseResp, err error) {

	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.PreferredAreaService.UpdatePreferredArea(l.ctx, &cmsclient.UpdatePreferredAreaReq{
		Id:         req.Id,         // 主键ID
		Name:       req.Name,       // 专区名称
		SubTitle:   req.SubTitle,   // 子标题
		Pic:        req.Pic,        // 展示图片
		Sort:       req.Sort,       // 排序
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   updateBy,       // 更新者
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新优选专区失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
