package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPostDetailLogic 查询岗位管理详情
/*
Author: LiuFeiHua
Date: 2024/5/30 11:28
*/
type QueryPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostDetailLogic {
	return &QueryPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryPostDetail 查询岗位管理详情
// 1.判断岗位信息是否存在
func (l *QueryPostDetailLogic) QueryPostDetail(in *sysclient.QueryPostDetailReq) (*sysclient.QueryPostDetailResp, error) {
	job, err := query.SysPost.WithContext(l.ctx).Where(query.SysPost.ID.Eq(in.Id)).First()

	// 1.判断岗位信息是否存在
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "岗位信息不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("岗位信息不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询岗位信息异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询岗位信息异常")
	}

	data := &sysclient.QueryPostDetailResp{
		Id:         job.ID,                                 // 岗位id
		PostName:   job.PostName,                           // 岗位名称
		PostCode:   job.PostCode,                           // 岗位编码
		PostStatus: job.PostStatus,                         // 岗位状态
		PostSort:   job.PostSort,                           // 岗位排序
		Remark:     job.Remark,                             // 备注信息
		IsDeleted:  job.IsDeleted,                          // 是否删除  0：否  1：是
		CreateBy:   job.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(job.CreateTime),    // 创建时间
		UpdateBy:   job.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(job.UpdateTime), // 更新时间
	}

	return data, nil
}
