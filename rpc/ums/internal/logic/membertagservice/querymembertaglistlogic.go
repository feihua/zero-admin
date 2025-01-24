package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTagListLogic 查询用户标签列表
/*
Author: LiuFeiHua
Date: 2025/01/24 10:32:59
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

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户标签列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户标签列表失败")
	}

	var list []*umsclient.MemberTagListData

	for _, item := range result {
		list = append(list, &umsclient.MemberTagListData{
			Id:                item.ID,                //
			TagName:           item.TagName,           // 标签名称
			FinishOrderCount:  item.FinishOrderCount,  // 自动打标签完成订单数量
			Status:            item.Status,            // 状态：0->禁用；1->启用
			FinishOrderAmount: item.FinishOrderAmount, // 自动打标签完成订单金额

		})
	}

	logc.Infof(l.ctx, "查询用户标签列表信息,参数：%+v,响应：%+v", in, list)

	return &umsclient.QueryMemberTagListResp{
		Total: count,
		List:  list,
	}, nil
}
