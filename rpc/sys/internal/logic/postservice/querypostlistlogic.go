package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPostListLogic 查询岗位列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:05
*/
type QueryPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostListLogic {
	return &QueryPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryPostList 查询岗位列表
func (l *QueryPostListLogic) QueryPostList(in *sysclient.QueryPostListReq) (*sysclient.QueryPostListResp, error) {
	q := query.SysPost.WithContext(l.ctx)
	if len(in.PostCode) > 0 {
		q = q.Where(query.SysPost.PostCode.Like("%" + in.PostCode + "%"))
	}
	if len(in.PostName) > 0 {
		q = q.Where(query.SysPost.PostName.Like("%" + in.PostName + "%"))
	}

	if in.PostStatus != 2 {
		q = q.Where(query.SysPost.PostStatus.Eq(in.PostStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询岗位列表信息失败")
	}

	var list []*sysclient.PostListData
	for _, job := range result {
		list = append(list, &sysclient.PostListData{
			Id:         job.ID,                                       // 岗位id
			PostName:   job.PostName,                                 // 岗位名称
			PostCode:   job.PostCode,                                 // 岗位编码
			PostStatus: job.PostStatus,                               // 岗位状态
			PostSort:   job.PostSort,                                 // 岗位排序
			Remark:     job.Remark,                                   // 备注信息
			IsDeleted:  job.IsDeleted,                                // 是否删除  0：否  1：是
			CreateBy:   job.CreateBy,                                 // 创建者
			CreateTime: job.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
			UpdateBy:   job.UpdateBy,                                 // 更新者
			UpdateTime: time_util.TimeToString(job.UpdateTime),       // 更新时间
		})
	}

	return &sysclient.QueryPostListResp{
		Total: count,
		List:  list,
	}, nil
}
