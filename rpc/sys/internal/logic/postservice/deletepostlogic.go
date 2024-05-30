package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeletePostLogic 删除岗位
/*
Author: LiuFeiHua
Date: 2023/12/18 17:04
*/
type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePost 删除岗位
func (l *DeletePostLogic) DeletePost(in *sysclient.DeletePostReq) (*sysclient.DeletePostResp, error) {
	q := query.SysPost
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除岗位失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除岗位失败")
	}

	return &sysclient.DeletePostResp{}, nil
}
