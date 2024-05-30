package dict_type

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

// UpdateDictTypeLogic 更新字典类型信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:18
*/
type UpdateDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateDictTypeLogic {
	return UpdateDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDictType 更新字典类型信息
func (l *UpdateDictTypeLogic) UpdateDictType(req *types.UpdateDictTypeReq) (*types.UpdateDictTypeResp, error) {
	_, err := l.svcCtx.DictTypeService.UpdateDictType(l.ctx, &sysclient.UpdateDictTypeReq{
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		Id:         req.Id,
		IsSystem:   req.IsSystem,
		Remark:     req.Remark,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典类型信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictTypeResp{
		Code:    "000000",
		Message: "更新字典类型成功",
	}, nil
}
