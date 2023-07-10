package collection

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCollectionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCollectionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCollectionDeleteLogic {
	return &ProductCollectionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCollectionDeleteLogic) ProductCollectionDelete(req *types.ProductCollectionDeleteReq) (resp *types.ProductCollectionDeleteResp, err error) {
	_, _ = l.svcCtx.MemberProductCollectionService.MemberProductCollectionDelete(l.ctx, &umsclient.MemberProductCollectionDeleteReq{
		Ids:      req.Ids,
		MemberId: l.ctx.Value("memberId").(int64),
	})

	return &types.ProductCollectionDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
