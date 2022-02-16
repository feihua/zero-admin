package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductDeleteLogic {
	return HomeNewProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductDeleteLogic) HomeNewProductDelete(req types.DeleteHomeNewProductReq) (*types.DeleteHomeNewProductResp, error) {
	_, err := l.svcCtx.Sms.HomeNewProductDelete(l.ctx, &smsclient.HomeNewProductDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除新鲜好物异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除新鲜好物失败")
	}

	return &types.DeleteHomeNewProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
