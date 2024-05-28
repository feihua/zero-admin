package dictitemservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemLogic {
	return &QueryDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDictItem 根据条件查询单条字典项表记录
func (l *QueryDictItemLogic) QueryDictItem(in *sysclient.DictItemReq) (*sysclient.DictItemResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DictItemResp{}, nil
}
