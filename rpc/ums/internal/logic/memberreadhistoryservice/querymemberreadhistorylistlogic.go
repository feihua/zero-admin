package memberreadhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberReadHistoryListLogic 查询用户商品浏览历史记录列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:07
*/
type QueryMemberReadHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReadHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReadHistoryListLogic {
	return &QueryMemberReadHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberReadHistoryList 查询用户商品浏览历史记录列表
func (l *QueryMemberReadHistoryListLogic) QueryMemberReadHistoryList(in *umsclient.QueryMemberReadHistoryListReq) (*umsclient.QueryMemberReadHistoryListResp, error) {
	q := query.UmsMemberReadHistory.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberLoginLog.MemberID.Eq(in.MemberId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员浏览列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberReadHistoryListData
	for _, item := range result {

		list = append(list, &umsclient.MemberReadHistoryListData{
			Id:              item.ID,
			MemberId:        item.MemberID,
			MemberNickName:  item.MemberNickName,
			MemberIcon:      item.MemberIcon,
			ProductId:       item.ProductID,
			ProductName:     item.ProductName,
			ProductPic:      item.ProductPic,
			ProductSubTitle: item.ProductSubTitle,
			ProductPrice:    item.ProductPrice,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询会员浏览列表信息,参数：%+v,响应：%+v", in, list)

	return &umsclient.QueryMemberReadHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
