package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/oms/omsclient"

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
	_, err := l.svcCtx.Oms.CompanyAddressDelete(l.ctx, &omsclient.CompanyAddressDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除公司地址异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除公司地址失败")
	}

	return &types.DeleteCompayAddressResp{
		Code:    "000000",
		Message: "",
	}, nil
}
