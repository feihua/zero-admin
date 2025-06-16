package product_attribute

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

// QueryProductAttributeDetailLogic 查询商品属性详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductAttributeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeDetailLogic {
	return &QueryProductAttributeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductAttributeDetail 查询商品属性详情
func (l *QueryProductAttributeDetailLogic) QueryProductAttributeDetail(req *types.QueryProductAttributeDetailReq) (resp *types.QueryProductAttributeDetailResp, err error) {

	detail, err := l.svcCtx.ProductAttributeService.QueryProductAttributeDetail(l.ctx, &pmsclient.QueryProductAttributeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品属性详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryProductAttributeDetailData{
		Id:           detail.Id,           // 主键id
		GroupId:      detail.GroupId,      // 属性分组ID
		Name:         detail.Name,         // 属性名称
		InputType:    detail.InputType,    // 输入类型：1-手动输入，2-单选，3-多选
		ValueType:    detail.ValueType,    // 值类型：1-文本，2-数字，3-日期
		InputList:    detail.InputList,    // 可选值列表，用逗号分隔
		Unit:         detail.Unit,         // 单位
		IsRequired:   detail.IsRequired,   // 是否必填
		IsSearchable: detail.IsSearchable, // 是否支持搜索
		IsShow:       detail.IsShow,       // 是否显示
		Sort:         detail.Sort,         // 排序
		Status:       detail.Status,       // 状态：0->禁用；1->启用
		CreateBy:     detail.CreateBy,     // 创建人ID
		CreateTime:   detail.CreateTime,   // 创建时间
		UpdateBy:     detail.UpdateBy,     // 更新人ID
		UpdateTime:   detail.UpdateTime,   // 更新时间

	}
	return &types.QueryProductAttributeDetailResp{
		Code:    "000000",
		Message: "查询商品属性成功",
		Data:    data,
	}, nil
}
