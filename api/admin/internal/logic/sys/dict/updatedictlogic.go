package dict

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictLogic 更新字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:18
*/
type UpdateDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDictLogic {
	return UpdateDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDict 更新字典信息
func (l *UpdateDictLogic) UpdateDict(req *types.UpdateDictReq) (*types.UpdateDictResp, error) {
	_, err := l.svcCtx.DictService.UpdateDict(l.ctx, &sysclient.DictUpdateReq{
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		Id:         req.Id,
		Remark:     req.Remark,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictResp{
		Code:    "000000",
		Message: "更新字典成功",
	}, nil
}
