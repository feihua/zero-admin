package membertaskrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTaskRelationListLogic 查询会员任务关联列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTaskRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTaskRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTaskRelationListLogic {
	return &QueryMemberTaskRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTaskRelationList 查询会员任务关联列表
func (l *QueryMemberTaskRelationListLogic) QueryMemberTaskRelationList(in *umsclient.QueryMemberTaskRelationListReq) (*umsclient.QueryMemberTaskRelationListResp, error) {
	memberTaskRelation := query.UmsMemberTaskRelation
	q := memberTaskRelation.WithContext(l.ctx)
	if in.MemberId != 2 {
		q = q.Where(memberTaskRelation.MemberID.Eq(in.MemberId))
	}
	if in.TaskId != 2 {
		q = q.Where(memberTaskRelation.TaskID.Eq(in.TaskId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员任务关联列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员任务关联列表失败")
	}

	var list []*umsclient.MemberTaskRelationListData

	for _, item := range result {
		list = append(list, &umsclient.MemberTaskRelationListData{
			Id:         item.ID,                              // 主键ID
			MemberId:   item.MemberID,                        // 会员ID
			TaskId:     item.TaskID,                          // 任务ID
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间

		})
	}

	return &umsclient.QueryMemberTaskRelationListResp{
		Total: count,
		List:  list,
	}, nil
}
