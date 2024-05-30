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

// AddDictTypeLogic 添加字典类型信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type AddDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddDictTypeLogic {
	return AddDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddDictType 添加字典类型信息
func (l *AddDictTypeLogic) AddDictType(req *types.AddDictTypeReq) (*types.AddDictTypeResp, error) {
	_, err := l.svcCtx.DictTypeService.AddDictType(l.ctx, &sysclient.AddDictTypeReq{
		CreateBy:   l.ctx.Value("userName").(string),
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		IsSystem:   req.IsSystem,
		Remark:     req.Remark,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加字典类型信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddDictTypeResp{
		Code:    "000000",
		Message: "添加字典类型成功",
	}, nil
}
