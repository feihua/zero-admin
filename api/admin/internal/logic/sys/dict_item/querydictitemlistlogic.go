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
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		DictLabel:  strings.TrimSpace(req.DictLabel),
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典数据列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDictItemListData

	for _, dict := range result.List {
		list = append(list, &types.QueryDictItemListData{
			CreateBy:   dict.CreateBy,
			CreateTime: dict.CreateTime,
			DictLabel:  dict.DictLabel,
			DictSort:   dict.DictSort,
			DictStatus: dict.DictStatus,
			DictType:   dict.DictType,
			DictValue:  dict.DictValue,
			Id:         dict.Id,
			IsDefault:  dict.IsDefault,
			Remark:     dict.Remark,
			UpdateBy:   dict.UpdateBy,
			UpdateTime: dict.UpdateTime,
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
