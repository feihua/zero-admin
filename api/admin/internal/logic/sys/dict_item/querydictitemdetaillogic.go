package dict_item

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictItemDetailLogic 查询字典数据详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:13
*/
type QueryDictItemDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemDetailLogic {
	return &QueryDictItemDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictItemDetail 查询字典数据详情
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(req *types.QueryDictItemDetailReq) (resp *types.QueryDictItemDetailResp, err error) {
	detail, err := l.svcCtx.DictItemService.QueryDictItemDetail(l.ctx, &sysclient.QueryDictItemDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	dictItem := types.QueryDictItemDetailData{
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
	}

	return &types.QueryDictItemDetailResp{
		Code:    "000000",
		Message: "查询字典数据详情成功",
		Data:    dictItem,
	}, nil
}
