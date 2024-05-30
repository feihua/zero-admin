package postservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePostLogic 更新岗位信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:06
*/
type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePost 更新岗位信息
// 1.根据岗位id查询岗位是否已存在
// 3.岗位存在时,则直接更新岗位
func (l *UpdatePostLogic) UpdatePost(in *sysclient.UpdatePostReq) (*sysclient.UpdatePostResp, error) {
	q := query.SysPost.WithContext(l.ctx)

	//1.根据岗位id查询岗位是否已存在
	_, err := q.Where(query.SysPost.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位id：%d,查询岗位信息失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询岗位信息失败"))
	}

	// 3.岗位存在时,则直接更新岗位
	var job = &model.SysPost{
		ID:         in.Id,
		PostName:   in.PostName,
		PostCode:   in.PostCode,
		PostStatus: in.PostStatus,
		PostSort:   in.PostSort,
		Remark:     in.Remark,
		UpdateBy:   in.UpdateBy,
	}
	_, err = q.Updates(job)

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位信息失败,参数:%+v,异常:%s", job, err.Error())
		return nil, errors.New("更新岗位信息失败")
	}

	return &sysclient.UpdatePostResp{}, nil
}
