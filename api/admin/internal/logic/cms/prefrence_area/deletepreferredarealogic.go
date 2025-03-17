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

// DeletePreferredAreaLogic 删除优选专区
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type DeletePreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePreferredAreaLogic {
	return &DeletePreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeletePreferredArea 删除优选专区
func (l *DeletePreferredAreaLogic) DeletePreferredArea(req *types.DeletePreferredAreaReq) (resp *types.DeletePreferredAreaResp, err error) {
	_, err = l.svcCtx.PreferredAreaService.DeletePreferredArea(l.ctx, &cmsclient.DeletePreferredAreaReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除优选专区失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeletePreferredAreaResp{
		Code:    "000000",
		Message: "删除优选专区成功",
	}, nil
}
