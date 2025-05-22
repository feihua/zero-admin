package membergrowthlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberGrowthLogListLogic 查询会员成长值记录列表
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberGrowthLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberGrowthLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberGrowthLogListLogic {
	return &QueryMemberGrowthLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberGrowthLogList 查询会员成长值记录列表
func (l *QueryMemberGrowthLogListLogic) QueryMemberGrowthLogList(in *umsclient.QueryMemberGrowthLogListReq) (*umsclient.QueryMemberGrowthLogListResp, error) {
	memberGrowthLog := query.UmsMemberGrowthLog
	q := memberGrowthLog.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(memberGrowthLog.MemberID.Eq(in.MemberId))
	}
	if in.ChangeType != 0 {
		q = q.Where(memberGrowthLog.ChangeType.Eq(in.ChangeType))
	}

	if in.SourceType != 5 {
		q = q.Where(memberGrowthLog.SourceType.Eq(in.SourceType))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员成长值记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员成长值记录列表失败")
	}

	var list []*umsclient.MemberGrowthLogListData

	for _, item := range result {
		list = append(list, &umsclient.MemberGrowthLogListData{
			Id:           item.ID,                                       //
			MemberId:     item.MemberID,                                 // 会员ID
			ChangeType:   item.ChangeType,                               // 变更类型：1-添加成长值，2-减少成长值
			ChangeGrowth: item.ChangeGrowth,                             // 变更成长值
			SourceType:   item.SourceType,                               // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
			Description:  item.Description,                              // 描述
			OperateMan:   item.OperateMan,                               // 操作人员
			OperateNote:  item.OperateNote,                              // 操作备注
			CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间

		})
	}

	return &umsclient.QueryMemberGrowthLogListResp{
		Total: count,
		List:  list,
	}, nil
}
