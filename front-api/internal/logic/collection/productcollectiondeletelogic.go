package collection

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCollectionDeleteLogic 收藏相关
/*
Author: LiuFeiHua
Date: 2023/11/30 12:02
*/
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

// ProductCollectionDelete 删除收藏的商品
func (l *ProductCollectionDeleteLogic) ProductCollectionDelete(req *types.ProductCollectionDeleteReq) (resp *types.ProductCollectionDeleteResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberProductCollectionService.MemberProductCollectionDelete(l.ctx, &umsclient.MemberProductCollectionDeleteReq{
		Id:       req.Id,
		MemberId: memberId,
	})

	return &types.ProductCollectionDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
