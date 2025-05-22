package memberpointslogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberPointsLogListLogic 查询会员积分记录列表
/*
Author: LiuFeiHua
Date: 2025/05/22 15:25:56
*/
type QueryMemberPointsLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberPointsLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberPointsLogListLogic {
	return &QueryMemberPointsLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberPointsLogList 查询会员积分记录列表
func (l *QueryMemberPointsLogListLogic) QueryMemberPointsLogList(in *umsclient.QueryMemberPointsLogListReq) (*umsclient.QueryMemberPointsLogListResp, error) {
	memberPointsLog := query.UmsMemberPointsLog
	q := memberPointsLog.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(memberPointsLog.MemberID.Eq(in.MemberId))
	}
	if in.ChangeType != 0 {
		q = q.Where(memberPointsLog.ChangeType.Eq(in.ChangeType))
	}

	if in.SourceType != 5 {
		q = q.Where(memberPointsLog.SourceType.Eq(in.SourceType))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员积分记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员积分记录列表失败")
	}

	var list []*umsclient.MemberPointsLogListData

	for _, item := range result {
		list = append(list, &umsclient.MemberPointsLogListData{
			Id:           item.ID,                                       //
			MemberId:     item.MemberID,                                 // 会员ID
			ChangeType:   item.ChangeType,                               // 变更类型：1-添加积分，2-减少积分
			ChangePoints: item.ChangePoints,                             // 变更积分
			SourceType:   item.SourceType,                               // 来源类型：0-其他，1-订单，2-活动，3-签到，4-管理员修改
			Description:  item.Description,                              // 描述
			OperateMan:   item.OperateMan,                               // 操作人员
			OperateNote:  item.OperateNote,                              // 操作备注
			CreateTime:   item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间

		})
	}

	return &umsclient.QueryMemberPointsLogListResp{
		Total: count,
		List:  list,
	}, nil
}
