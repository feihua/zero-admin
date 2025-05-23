package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTagListLogic 查询用户标签列表
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagListLogic {
	return &QueryMemberTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTagList 查询用户标签列表
func (l *QueryMemberTagListLogic) QueryMemberTagList(in *umsclient.QueryMemberTagListReq) (*umsclient.QueryMemberTagListResp, error) {
	memberTag := query.UmsMemberTag
	q := memberTag.WithContext(l.ctx)
	if len(in.TagName) > 0 {
		q = q.Where(memberTag.TagName.Like("%" + in.TagName + "%"))
	}

	if in.Status != 2 {
		q = q.Where(memberTag.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户标签列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户标签列表失败")
	}

	var list []*umsclient.MemberTagListData

	for _, item := range result {
		list = append(list, &umsclient.MemberTagListData{
			Id:                item.ID,                                          // 主键ID
			TagName:           item.TagName,                                     // 标签名称
			Description:       item.Description,                                 // 标签描述
			FinishOrderCount:  item.FinishOrderCount,                            // 自动打标签完成订单数量
			FinishOrderAmount: float32(item.FinishOrderAmount),                  // 自动打标签完成订单金额
			Status:            item.Status,                                      // 状态：0-禁用，1-启用
			CreateBy:          item.CreateBy,                                    // 创建人ID
			CreateTime:        item.CreateTime.Format("2006-01-02 15:04:05"),    // 创建时间
			UpdateBy:          pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime:        time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &umsclient.QueryMemberTagListResp{
		Total: count,
		List:  list,
	}, nil
}
