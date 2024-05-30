package dict_type

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryDictTypeListLogic 查询字典类型列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type QueryDictTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryDictTypeListLogic {
	return QueryDictTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictTypeList 查询字典类型列表
func (l *QueryDictTypeListLogic) QueryDictTypeList(req *types.QueryDictTypeListReq) (*types.QueryDictTypeListResp, error) {
	var resp, err = l.svcCtx.DictTypeService.QueryDictTypeList(l.ctx, &sysclient.QueryDictTypeListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		DictName:   req.DictName,
		DictStatus: req.DictStatus,
		DictType:   req.DictType,
		IsSystem:   req.IsSystem,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDictTypeListData

	for _, dict := range resp.List {
		list = append(list, &types.QueryDictTypeListData{
			CreateBy:   dict.CreateBy,
			CreateTime: dict.CreateTime,
			DictName:   dict.DictType,
			DictStatus: dict.DictStatus,
			DictType:   dict.DictType,
			Id:         dict.Id,
			IsSystem:   dict.IsSystem,
			Remark:     dict.Remark,
			UpdateBy:   dict.UpdateBy,
			UpdateTime: dict.UpdateTime,
		})
	}

	return &types.QueryDictTypeListResp{
		Code:     "000000",
		Message:  "查询字典类型成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
