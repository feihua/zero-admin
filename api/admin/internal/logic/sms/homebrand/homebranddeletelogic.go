package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandDeleteLogic {
	return HomeBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandDeleteLogic) HomeBrandDelete(req types.DeleteHomeBrandReq) (*types.DeleteHomeBrandResp, error) {
	_, err := l.svcCtx.HomeBrandService.HomeBrandDelete(l.ctx, &smsclient.HomeBrandDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除首页品牌异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除首页品牌失败")
	}

	return &types.DeleteHomeBrandResp{
		Code:    "000000",
		Message: "",
	}, nil
}
