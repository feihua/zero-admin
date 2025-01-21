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

// AddDictItemLogic 添加字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 16:01
*/
type AddDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictItemLogic {
	return &AddDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddDictItem 添加字典数据
func (l *AddDictItemLogic) AddDictItem(req *types.AddDictItemReq) (resp *types.AddDictItemResp, err error) {
	_, err = l.svcCtx.DictItemService.AddDictItem(l.ctx, &sysclient.AddDictItemReq{
		DictType:   req.DictType,   // 字典类型
		DictLabel:  req.DictLabel,  // 字典标签
		DictValue:  req.DictValue,  // 字典键值
		DictStatus: req.DictStatus, // 字典状态
		DictSort:   req.DictSort,   // 排序
		Remark:     req.Remark,     // 备注信息
		IsDefault:  req.IsDefault,  // 是否默认  0：否  1：是
		CreateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加字典项信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddDictItemResp{
		Code:    "000000",
		Message: "添加字典项成功",
	}, nil
}
