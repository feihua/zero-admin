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
// 1.根据postName查询岗位是否已存在,如果岗位已存在,则直接返回
// 2.根据postCode查询岗位是否已存在,如果岗位已存在,则直接返回
// 3.岗位不存在时,则直接添加岗位
func (l *AddPostLogic) AddPost(in *sysclient.AddPostReq) (*sysclient.AddPostResp, error) {

	q := query.SysPost.WithContext(l.ctx)

	// 1.根据postName查询岗位是否已存在,如果岗位已存在,则直接返回
	postName := in.PostName
	count, err := q.Where(query.SysPost.PostName.Eq(postName)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位名称：%s,查询岗位信息,异常:%s", postName, err.Error())
		return nil, errors.New(fmt.Sprintf("添加岗位信息失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("添加岗位信息失败,岗位名称：%s,已存在", postName))
	}

	// 2.根据postCode查询岗位是否已存在,如果岗位已存在,则直接返回
	postCode := in.PostCode
	count, err = q.Where(query.SysPost.PostCode.Eq(postCode)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位编码：%s,查询岗位信息,异常:%s", postName, err.Error())
		return nil, errors.New(fmt.Sprintf("添加岗位信息失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("添加岗位信息失败，岗位编码：%s,已存在", postCode))
	}

	// 3.岗位不存在时,则直接添加岗位
	job := &model.SysPost{
		PostCode: in.PostCode, // 岗位编码
		PostName: in.PostName, // 岗位名称
		Sort:     in.Sort,     // 显示顺序
		Status:   in.Status,   // 岗位状态（0：停用，1:正常）
		Remark:   in.Remark,   // 备注
		CreateBy: in.CreateBy, // 创建者
	}

	err = q.Create(job)

	if err != nil {
		logc.Errorf(l.ctx, "添加岗位管理失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加岗位信息失败")
	}

	return &sysclient.AddPostResp{}, nil
}
