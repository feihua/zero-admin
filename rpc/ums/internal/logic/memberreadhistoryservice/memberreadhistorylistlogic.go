package memberreadhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReadHistoryListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:35
*/
type MemberReadHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryListLogic {
	return &MemberReadHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReadHistoryList 浏览记录
func (l *MemberReadHistoryListLogic) MemberReadHistoryList(in *umsclient.MemberReadHistoryListReq) (*umsclient.MemberReadHistoryListResp, error) {
	q := query.UmsMemberReadHistory.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberLoginLog.MemberID.Eq(in.MemberId))
	}
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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
			ProductSubTitle: *item.ProductSubTitle,
			ProductPrice:    item.ProductPrice,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询会员浏览列表信息,参数：%+v,响应：%+v", in, list)

	return &umsclient.MemberReadHistoryListResp{
		Total: count,
		List:  list,
	}, nil
}
