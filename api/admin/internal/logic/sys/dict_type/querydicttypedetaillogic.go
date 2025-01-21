package dict_type

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

// QueryDictTypeDetailLogic 查询字典类型详情
/*
Author: LiuFeiHua
Date: 2024/5/29 17:01
*/
type QueryDictTypeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeDetailLogic {
	return &QueryDictTypeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDictTypeDetail 查询字典类型详情
func (l *QueryDictTypeDetailLogic) QueryDictTypeDetail(req *types.QueryDictTypeDetailReq) (resp *types.QueryDictTypeDetailResp, err error) {
	detail, err := l.svcCtx.DictTypeService.QueryDictTypeDetail(l.ctx, &sysclient.QueryDictTypeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字典类型列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryDictTypeDetailData{
		Id:         detail.Id,         // 编号
		DictName:   detail.DictName,   // 字典名称
		DictType:   detail.DictType,   // 字典类型
		DictStatus: detail.DictStatus, // 字典状态
		Remark:     detail.Remark,     // 备注信息
		IsSystem:   detail.IsSystem,   // 是否系统预留  0：否  1：是
		IsDeleted:  detail.IsDeleted,  // 是否删除  0：否  1：是
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间

	}

	return &types.QueryDictTypeDetailResp{
		Code:    "000000",
		Message: "查询字典类型成功",
		Data:    data,
	}, nil
}
