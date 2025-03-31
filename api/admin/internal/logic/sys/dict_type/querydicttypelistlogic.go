package dict_type

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

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
		PageNum:  req.Current,
		PageSize: req.PageSize,
		DictName: strings.TrimSpace(req.DictName), // 字典名称
		DictType: req.DictType,                    // 字典类型
		Status:   req.Status,                      // 状态（0：停用，1:正常）
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDictTypeListData

	for _, detail := range resp.List {
		list = append(list, &types.QueryDictTypeListData{
			Id:         detail.Id,         // 字典id
			DictName:   detail.DictName,   // 字典名称
			DictType:   detail.DictType,   // 字典类型
			Status:     detail.Status,     // 状态（0：停用，1:正常）
			Remark:     detail.Remark,     // 备注
			CreateBy:   detail.CreateBy,   // 创建者
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新者
			UpdateTime: detail.UpdateTime, // 更新时间
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
