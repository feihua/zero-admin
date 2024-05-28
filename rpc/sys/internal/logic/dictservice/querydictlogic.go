package dictservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictLogic {
	return &QueryDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryDict 根据条件查询单条字典表记录
func (l *QueryDictLogic) QueryDict(in *sysclient.DictReq) (*sysclient.DictResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DictResp{}, nil
}
