package dict_item

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

// UpdateDictItemLogic 字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 16:01
*/
type UpdateDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemLogic {
	return &UpdateDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDictItem 更新字典数据
func (l *UpdateDictItemLogic) UpdateDictItem(req *types.UpdateDictItemReq) (resp *types.UpdateDictItemResp, err error) {
	_, err = l.svcCtx.DictItemService.UpdateDictItem(l.ctx, &sysclient.UpdateDictItemReq{
		DictLabel:  req.DictLabel,
		DictSort:   req.DictSort,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		DictValue:  req.DictValue,
		Id:         req.Id,
		IsDefault:  req.IsDefault,
		Remark:     req.Remark,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典项失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictItemResp{
		Code:    "000000",
		Message: "更新字典项成功",
	}, nil
}
