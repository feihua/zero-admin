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

// AddPostLogic 添加岗位
/*
Author: LiuFeiHua
Date: 2023/12/18 17:04
*/
type AddPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPost 添加岗位
// 1.根据岗位名称查询岗位是否已存在
// 2.如果岗位已存在,则直接返回
// 3.岗位不存在时,则直接添加岗位
func (l *AddPostLogic) AddPost(in *sysclient.AddPostReq) (*sysclient.AddPostResp, error) {

	q := query.SysPost.WithContext(l.ctx)

	//1.根据岗位名称查询岗位是否已存在
	postName := in.PostName
	count, err := q.Where(query.SysPost.PostName.Eq(postName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位名称：%s,查询岗位信息失败,异常:%s", postName, err.Error())
		return nil, errors.New(fmt.Sprintf("查询岗位信息失败"))
	}

	//2.如果岗位已存在,则直接返回
	if count > 0 {
		logc.Errorf(l.ctx, "岗位信息已存在：%+v", in)
		return nil, errors.New(fmt.Sprintf("岗位：%s,已存在", postName))
	}

	//3.岗位不存在时,则直接添加岗位
	job := &model.SysPost{
		PostName:   in.PostName,
		PostCode:   in.PostCode,
		PostStatus: in.PostStatus,
		PostSort:   in.PostSort,
		Remark:     in.Remark,
		IsDeleted:  0,
		CreateBy:   in.CreateBy,
	}

	err = q.Create(job)

	if err != nil {
		logc.Errorf(l.ctx, "添加岗位信息失败,参数:%+v,异常:%s", job, err.Error())
		return nil, errors.New("添加岗位信息失败")
	}

	return &sysclient.AddPostResp{}, nil
}
