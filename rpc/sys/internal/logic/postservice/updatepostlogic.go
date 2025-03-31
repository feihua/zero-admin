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
	"gorm.io/gorm"
	"time"

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
// 1.判断岗位信息是否存在
// 2.查询postName是否被占用,如果被占用,则直接返回
// 3.查询根据postCode是否被占用,如果被占用,则直接返回
// 4.岗位存在时,则直接更新岗位
func (l *UpdatePostLogic) UpdatePost(in *sysclient.UpdatePostReq) (*sysclient.UpdatePostResp, error) {
	q := query.SysPost.WithContext(l.ctx)

	// 1.判断岗位信息是否存在
	post, err := q.Where(query.SysPost.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "岗位信息不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("岗位信息不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询岗位信息异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询岗位信息异常")
	}

	// 2.查询postName是否被占用,如果被占用,则直接返回
	postName := in.PostName
	count, err := q.Where(query.SysPost.PostName.Eq(postName), query.SysPost.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位名称：%s,查询岗位信息,异常:%s", postName, err.Error())
		return nil, errors.New(fmt.Sprintf("更新岗位信息失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("更新岗位信息失败,岗位名称：%s,已存在", postName))
	}

	// 3.查询postCode是否被占用,如果被占用,则直接返回
	postCode := in.PostCode
	count, err = q.Where(query.SysPost.PostCode.Eq(postCode), query.SysPost.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "根据岗位编码：%s,查询岗位信息,异常:%s", postName, err.Error())
		return nil, errors.New(fmt.Sprintf("更新岗位信息失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("更新岗位信息失败，岗位编码：%s,已存在", postCode))
	}

	// 4.岗位存在时,则直接更新岗位
	now := time.Now()
	var job = &model.SysPost{
		ID:         in.Id,           // 岗位id
		PostCode:   in.PostCode,     // 岗位编码
		PostName:   in.PostName,     // 岗位名称
		Sort:       in.Sort,         // 显示顺序
		Status:     in.Status,       // 岗位状态（0：停用，1:正常）
		Remark:     in.Remark,       // 备注信息
		CreateBy:   post.CreateBy,   // 创建者
		CreateTime: post.CreateTime, // 创建时间
		UpdateBy:   in.UpdateBy,     // 更新者
		UpdateTime: &now,            // 更新时间
	}

	err = l.svcCtx.DB.Model(&model.SysPost{}).WithContext(l.ctx).Where(query.SysPost.ID.Eq(in.Id)).Save(job).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位信息失败,参数:%+v,异常:%s", job, err.Error())
		return nil, errors.New("更新岗位信息失败")
	}

	return &sysclient.UpdatePostResp{}, nil
}
