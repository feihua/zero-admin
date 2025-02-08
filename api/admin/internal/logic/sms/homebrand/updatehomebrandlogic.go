package homebrand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeBrandSortLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type UpdateHomeBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeBrandLogic {
	return &UpdateHomeBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeBrandSort 修改推荐品牌排序
func (l *UpdateHomeBrandLogic) UpdateHomeBrand(req *types.UpdateHomeBrandReq) (resp *types.UpdateHomeBrandResp, err error) {
	_, err = l.svcCtx.HomeBrandService.UpdateHomeBrandSort(l.ctx, &smsclient.UpdateHomeBrandSortReq{
		Id:   req.Id,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新首页品牌排序失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateHomeBrandResp{
		Code:    "000000",
		Message: "更新首页品牌排序成功",
	}, nil
}
