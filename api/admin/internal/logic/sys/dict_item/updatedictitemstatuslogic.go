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

// UpdateDictItemStatusLogic 更新字典数据状态
type UpdateDictItemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemStatusLogic {
	return &UpdateDictItemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDictItemStatus 更新字典数据状态
func (l *UpdateDictItemStatusLogic) UpdateDictItemStatus(req *types.UpdateDictItemStatusReq) (resp *types.UpdateDictItemStatusResp, err error) {
	_, err = l.svcCtx.DictItemService.UpdateDictItemStatus(l.ctx, &sysclient.UpdateDictItemStatusReq{
		Ids:      req.Ids,    // 字典数据id
		Status:   req.Status, // 状态（0：停用，1:正常）
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典数据状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictItemStatusResp{
		Code:    "000000",
		Message: "更新字典数据状态成功",
	}, nil
}
