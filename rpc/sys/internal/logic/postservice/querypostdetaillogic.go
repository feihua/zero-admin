package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

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
func (l *QueryPostDetailLogic) QueryPostDetail(in *sysclient.QueryPostDetailReq) (*sysclient.QueryPostDetailResp, error) {
	job, err := query.SysPost.WithContext(l.ctx).Where(query.SysPost.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位信息详情失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询岗位信息详情失败")
	}

	if job == nil {
		logc.Errorf(l.ctx, "查询岗位信息详情失败,参数：%+v,岗位信息不存在", in)
		return nil, errors.New("查询岗位信息详情失败,岗位信息不存在")
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
