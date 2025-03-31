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

	if in.Status != 2 {
		q = q.Where(query.SysPost.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询岗位列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询岗位列表信息失败")
	}

	var list = make([]*sysclient.PostListData, 0, len(result))
	for _, post := range result {
		list = append(list, &sysclient.PostListData{
			Id:         post.ID,                                 // 岗位id
			PostCode:   post.PostCode,                           // 岗位编码
			PostName:   post.PostName,                           // 岗位名称
			Sort:       post.Sort,                               // 显示顺序
			Status:     post.Status,                             // 岗位状态（0：停用，1:正常）
			Remark:     post.Remark,                             // 备注
			CreateBy:   post.CreateBy,                           // 创建者
			CreateTime: time_util.TimeToStr(post.CreateTime),    // 创建时间
			UpdateBy:   post.UpdateBy,                           // 更新者
			UpdateTime: time_util.TimeToString(post.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryPostListResp{
		Total: count,
		List:  list,
	}, nil
}
