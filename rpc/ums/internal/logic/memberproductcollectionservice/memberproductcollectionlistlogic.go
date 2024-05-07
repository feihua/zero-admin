package memberproductcollectionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberProductCollectionListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:37
*/
type MemberProductCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionListLogic {
	return &MemberProductCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberProductCollectionList 收藏列表
func (l *MemberProductCollectionListLogic) MemberProductCollectionList(in *umsclient.MemberProductCollectionListReq) (*umsclient.MemberProductCollectionListResp, error) {
	q := query.UmsMemberProductCollection.WithContext(l.ctx)
	if in.MemberId != 0 {
		q.Where(query.UmsMemberProductCollection.MemberID.Eq(in.MemberId))
	}
	if in.ProductId != 0 {
		q.Where(query.UmsMemberProductCollection.ProductID.Eq(in.ProductId))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员收藏列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberProductCollectionListData
	for _, item := range result {

		list = append(list, &umsclient.MemberProductCollectionListData{
			Id:              item.ID,
			MemberId:        item.MemberID,
			MemberNickName:  item.MemberNickName,
			MemberIcon:      item.MemberIcon,
			ProductId:       in.ProductId,
			ProductName:     item.ProductName,
			ProductPic:      item.ProductPic,
			ProductSubTitle: *item.ProductSubTitle,
			ProductPrice:    item.ProductPrice,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询会员收藏列表信息,参数：%+v,响应：%+v", in, list)

	return &umsclient.MemberProductCollectionListResp{
		Total: count,
		List:  list,
	}, nil
}
