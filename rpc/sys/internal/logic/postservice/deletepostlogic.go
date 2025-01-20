package postservicelogic

import (
	"context"
	"errors"
	"fmt"
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
	if len(in.Ids) == 0 {
		logc.Errorf(l.ctx, "删除岗位失败,参数:%+v,异常:%s", in, "岗位id不能为空")
		return nil, errors.New("岗位id不能为空")
	}

	for _, id := range in.Ids {
		count, err := q.WithContext(l.ctx).Where(q.ID.Eq(id)).Count()
		if err != nil {
			logc.Errorf(l.ctx, "查询岗位信息异常,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New(fmt.Sprintf("查询岗位信息异常"))
		}

		if count == 0 {
			return nil, errors.New(fmt.Sprintf("删除岗位失败,岗位信息不存在"))
		}

		count, err = query.SysUserPost.WithContext(l.ctx).Where(query.SysUserPost.PostID.Eq(id)).Count()
		if err != nil {
			logc.Errorf(l.ctx, "查询岗位使用数量异常,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New(fmt.Sprintf("查询岗位使用数量异常"))
		}

		if count > 0 {
			return nil, errors.New(fmt.Sprintf("删除岗位信息失败,已分配,不能删除"))
		}
	}

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除岗位失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除岗位失败")
	}

	return &sysclient.DeletePostResp{}, nil
}
