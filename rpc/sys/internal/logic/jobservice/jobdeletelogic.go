package jobservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobDeleteLogic 删除岗位
/*
Author: LiuFeiHua
Date: 2023/12/18 17:04
*/
type JobDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobDeleteLogic {
	return &JobDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// JobDelete 删除岗位
func (l *JobDeleteLogic) JobDelete(in *sysclient.JobDeleteReq) (*sysclient.JobDeleteResp, error) {
	q := query.SysJob
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除岗位失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除岗位失败")
	}

	return &sysclient.JobDeleteResp{}, nil
}
