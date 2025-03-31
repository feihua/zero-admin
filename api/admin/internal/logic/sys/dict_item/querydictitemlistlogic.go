package dict_item

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictItemListLogic 字典数据列表
/*
Author: LiuFeiHua
Date: 2024/5/28 16:01
*/
type QueryDictItemListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemListLogic {
	return &QueryDictItemListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictItemList 查询字典数据列表
func (l *QueryDictItemListLogic) QueryDictItemList(req *types.QueryDictItemListReq) (resp *types.QueryDictItemListResp, err error) {
	result, err := l.svcCtx.DictItemService.QueryDictItemList(l.ctx, &sysclient.QueryDictItemListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		DictLabel: strings.TrimSpace(req.DictLabel), // 字典标签
		DictType:  req.DictType,                     // 字典类型
		Status:    req.Status,                       // 状态（0：停用，1:正常）
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDictItemListData

	for _, detail := range result.List {
		list = append(list, &types.QueryDictItemListData{
			Id:         detail.Id,         // 字典数据id
			DictSort:   detail.DictSort,   // 字典排序
			DictLabel:  detail.DictLabel,  // 字典标签
			DictValue:  detail.DictValue,  // 字典键值
			DictType:   detail.DictType,   // 字典类型
			CssClass:   detail.CssClass,   // 样式属性（其他样式扩展）
			ListClass:  detail.ListClass,  // 表格回显样式
			IsDefault:  detail.IsDefault,  // 是否默认（Y是 N否）
			Status:     detail.Status,     // 状态（0：停用，1:正常）
			Remark:     detail.Remark,     // 备注
			CreateBy:   detail.CreateBy,   // 创建者
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新者
			UpdateTime: detail.UpdateTime, // 更新时间
		})
	}

	return &types.QueryDictItemListResp{
		Code:     "000000",
		Message:  "查询字典数据列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
