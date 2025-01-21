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
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		DictName:   strings.TrimSpace(req.DictName), // 字典名称
		DictType:   req.DictType,                    // 字典类型
		DictStatus: req.DictStatus,                  // 字典状态
		IsSystem:   req.IsSystem,                    // 是否系统预留  0：否  1：是
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryDictTypeListData

	for _, dict := range resp.List {
		list = append(list, &types.QueryDictTypeListData{
			Id:         dict.Id,         // 编号
			DictName:   dict.DictName,   // 字典名称
			DictType:   dict.DictType,   // 字典类型
			DictStatus: dict.DictStatus, // 字典状态
			Remark:     dict.Remark,     // 备注信息
			IsSystem:   dict.IsSystem,   // 是否系统预留  0：否  1：是
			IsDeleted:  dict.IsDeleted,  // 是否删除  0：否  1：是
			CreateBy:   dict.CreateBy,   // 创建者
			CreateTime: dict.CreateTime, // 创建时间
			UpdateBy:   dict.UpdateBy,   // 更新者
			UpdateTime: dict.UpdateTime, // 更新时间
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
