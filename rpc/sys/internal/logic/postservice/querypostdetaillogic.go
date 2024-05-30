package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/logic/common"
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
		logc.Errorf(l.ctx, "查询岗位管理详情失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询岗位管理详情失败")
	}
	data := &sysclient.QueryPostDetailResp{
		CreateBy:   job.CreateBy,
		CreateTime: job.CreateTime.Format("2006-01-02 15:04:05"),
		Id:         job.ID,
		PostCode:   job.PostCode,
		PostName:   job.PostName,
		PostSort:   job.PostSort,
		PostStatus: job.PostStatus,
		Remark:     job.Remark,
		UpdateBy:   job.UpdateBy,
		UpdateTime: common.TimeToString(job.UpdateTime),
	}

	logc.Infof(l.ctx, "查询岗位管理详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
