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

// DeleteDictItemLogic 字典数据
/*
Author: LiuFeiHua
Date: 2024/5/28 16:01
*/
type DeleteDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictItemLogic {
	return &DeleteDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteDictItem 删除字典数据
func (l *DeleteDictItemLogic) DeleteDictItem(req *types.DeleteDictItemReq) (resp *types.DeleteDictItemResp, err error) {
	_, err = l.svcCtx.DictItemService.DeleteDictItem(l.ctx, &sysclient.DeleteDictItemReq{
		Ids: req.Ids, // 编号
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据dictItemId: %+v,删除字典项异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteDictItemResp{
		Code:    "000000",
		Message: "删除字典项成功",
	}, nil
}
