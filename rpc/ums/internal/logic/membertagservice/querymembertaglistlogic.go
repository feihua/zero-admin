package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberTagListLogic 查询用户标签表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 13:47
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

// QueryMemberTagList 查询用户标签表列表
func (l *QueryMemberTagListLogic) QueryMemberTagList(in *umsclient.QueryMemberTagListReq) (*umsclient.QueryMemberTagListResp, error) {
	q := query.UmsMemberTag.WithContext(l.ctx)
	offset := (in.PageNum - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员标签列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberTagListData
	for _, item := range result {

		list = append(list, &umsclient.MemberTagListData{
			FinishOrderAmount: item.FinishOrderAmount,
			FinishOrderCount:  item.FinishOrderCount,
			Id:                item.ID,
			Status:            in.Status,
			TagName:           item.TagName,
		})
	}

	return &umsclient.QueryMemberTagListResp{
		Total: count,
		List:  list,
	}, nil

}
