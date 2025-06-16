package product_attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductAttributeLogic 添加商品属性
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductAttributeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeLogic {
	return &AddProductAttributeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductAttribute 添加商品属性
func (l *AddProductAttributeLogic) AddProductAttribute(req *types.AddProductAttributeReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductAttributeService.AddProductAttribute(l.ctx, &pmsclient.AddProductAttributeReq{
		GroupId:      req.GroupId,      // 属性分组ID
		Name:         req.Name,         // 属性名称
		InputType:    req.InputType,    // 输入类型：1-手动输入，2-单选，3-多选
		ValueType:    req.ValueType,    // 值类型：1-文本，2-数字，3-日期
		InputList:    req.InputList,    // 可选值列表，用逗号分隔
		Unit:         req.Unit,         // 单位
		IsRequired:   req.IsRequired,   // 是否必填
		IsSearchable: req.IsSearchable, // 是否支持搜索
		IsShow:       req.IsShow,       // 是否显示
		Sort:         req.Sort,         // 排序
		Status:       req.Status,       // 状态：0->禁用；1->启用
		CreateBy:     userId,           // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加商品属性成功",
	}, nil
}
