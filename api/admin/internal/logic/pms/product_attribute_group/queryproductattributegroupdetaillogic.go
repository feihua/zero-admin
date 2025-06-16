package product_attribute_group

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeGroupDetailLogic 查询商品属性分组详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeGroupDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductAttributeGroupDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeGroupDetailLogic {
	return &QueryProductAttributeGroupDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductAttributeGroupDetail 查询商品属性分组详情
func (l *QueryProductAttributeGroupDetailLogic) QueryProductAttributeGroupDetail(req *types.QueryProductAttributeGroupDetailReq) (resp *types.QueryProductAttributeGroupDetailResp, err error) {

	detail, err := l.svcCtx.ProductAttributeGroupService.QueryProductAttributeGroupDetail(l.ctx, &pmsclient.QueryProductAttributeGroupDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性分组详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryProductAttributeGroupDetailData{
		Id:         detail.Id,         // 主键id
		CategoryId: detail.CategoryId, // 分类ID
		Name:       detail.Name,       // 分组名称
		Sort:       detail.Sort,       // 排序
		Status:     detail.Status,     // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新人ID
		UpdateTime: detail.UpdateTime, // 更新时间

	}
	return &types.QueryProductAttributeGroupDetailResp{
		Code:    "000000",
		Message: "查询商品属性分组成功",
		Data:    data,
	}, nil
}
