package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompayAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompayAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CompayAddressDeleteLogic {
	return CompayAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompayAddressDeleteLogic) CompayAddressDelete(req types.DeleteCompayAddressReq) (*types.DeleteCompayAddressResp, error) {
	_, err := l.svcCtx.CompanyAddressService.CompanyAddressDelete(l.ctx, &omsclient.CompanyAddressDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除公司地址异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除公司地址失败")
	}

	return &types.DeleteCompayAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
