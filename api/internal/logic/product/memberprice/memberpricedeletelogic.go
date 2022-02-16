package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceDeleteLogic {
	return MemberPriceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceDeleteLogic) MemberPriceDelete(req types.DeleteMemberPriceReq) (*types.DeleteMemberPriceResp, error) {
	_, err := l.svcCtx.Pms.MemberPriceDelete(l.ctx, &pmsclient.MemberPriceDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员价格异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员价格失败")
	}

	return &types.DeleteMemberPriceResp{
		Code:    "000000",
		Message: "",
	}, nil
}
