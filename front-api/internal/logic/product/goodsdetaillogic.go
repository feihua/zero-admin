package product

import (
	"context"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodsDetailLogic {
	return GoodsDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsDetailLogic) GoodsDetail(req types.GoodsDetailReq) (resp *types.GoodsDetailResp, err error) {
	// productDetail, err := l.svcCtx.Pms.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{
	// 	Id: req.Id,
	// })
	// if err != nil {
	// 	return nil, errors.Wrapf(xerr.NewErrMsg("参数有误"), "根据产品ID %d 获取产品详情有误", req.Id)
	// }
	// data := types.GoodsDetailData{
	// 	SpecificationList: nil,
	// 	Issue:             nil,
	// 	UserHasCollect:    0,
	// 	ShareImage:        "",
	// 	Share:             false,
	// 	Attribute:         nil,
	// 	ProductList:       nil,
	// 	Info:              types.Info{},
	// }

	// return &types.GoodsDetailResp{
	// 	Errno:  0,
	// 	Data:   data,
	// 	Errmsg: "",
	// }, nil
	return &types.GoodsDetailResp{}, nil
}
