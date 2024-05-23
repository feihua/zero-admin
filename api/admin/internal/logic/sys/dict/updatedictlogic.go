package dict

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateDictLogic
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
	_, err := l.svcCtx.DictService.DictUpdate(l.ctx, &sysclient.DictUpdateReq{
		Id:          req.Id,
		Value:       req.Value,
		Label:       req.Label,
		Type:        req.Type,
		Description: req.Description,
		Sort:        req.Sort,
		Remarks:     req.Remarks,
		UpdateBy:    l.ctx.Value("userName").(string),
		DelFlag:     req.DelFlag,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新字典失败")
	}

	return &types.UpdateDictResp{
		Code:    "000000",
		Message: "更新字典成功",
	}, nil
}
