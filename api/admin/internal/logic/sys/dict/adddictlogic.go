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

// AddDictLogic 添加字典信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type AddDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddDictLogic {
	return AddDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddDict 添加字典信息
func (l *AddDictLogic) AddDict(req *types.AddDictReq) (*types.AddDictResp, error) {
	_, err := l.svcCtx.DictService.AddDict(l.ctx, &sysclient.DictAddReq{
		CreateBy:   l.ctx.Value("userName").(string),
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		Remark:     req.Remark,
		IsSystem:   req.IsSystem,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加字典信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddDictResp{
		Code:    "000000",
		Message: "添加字典成功",
	}, nil
}
