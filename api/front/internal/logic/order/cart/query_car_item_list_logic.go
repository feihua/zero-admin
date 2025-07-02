package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCarItemListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCarItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCarItemListLogic {
	return &QueryCarItemListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCarItemListLogic) QueryCarItemList() (resp *types.CartItemListResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	itemListResp, err := l.svcCtx.CartItemService.QueryCartItemList(l.ctx, &omsclient.QueryCartItemListReq{MemberId: memberId})

	if err != nil {
		logc.Errorf(l.ctx, "获取某个会员的购物车列表失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	list := make([]types.CartListData, 0)

	for _, detail := range itemListResp.List {
		list = append(list, types.CartListData{
			Id:                detail.Id,                // 主键ID
			MemberId:          detail.MemberId,          // 会员ID
			ProductId:         detail.ProductId,         // 商品ID
			ProductSkuId:      detail.ProductSkuId,      // 商品SKU ID
			Quantity:          detail.Quantity,          // 购买数量
			Price:             detail.Price,             // 添加到购物车时的价格
			Selected:          detail.Selected,          // 是否选中 0-未选中 1-选中
			ProductName:       detail.ProductName,       // 商品名称
			ProductSubTitle:   detail.ProductSubTitle,   // 商品副标题
			ProductPic:        detail.ProductPic,        // 商品主图URL
			ProductSkuCode:    detail.ProductSkuCode,    // 商品SKU编码
			ProductSn:         detail.ProductSn,         // 商品货号
			ProductBrand:      detail.ProductBrand,      // 商品品牌
			ProductCategoryId: detail.ProductCategoryId, // 商品分类ID
			ProductAttr:       detail.ProductAttr,       // 商品销售属性JSON
			MemberNickname:    detail.MemberNickname,    // 会员昵称
			Source:            detail.Source,            // 来源 1-PC 2-H5 3-小程序 4-APP
			DeleteStatus:      detail.DeleteStatus,      // 删除状态 0-正常 1-删除
			ExpireTime:        detail.ExpireTime,        // 过期时间
			CreateTime:        detail.CreateTime,        // 创建时间
			UpdateTime:        detail.UpdateTime,        // 更新时间

		})
	}

	return &types.CartItemListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
