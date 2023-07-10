package memberproductcollectionservicelogic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberProductCollectionListLogic) MemberProductCollectionList(in *umsclient.MemberProductCollectionListReq) (*umsclient.MemberProductCollectionListResp, error) {
	all, err := l.svcCtx.UmsMemberProductCollectionModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.UmsMemberProductCollectionModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员收藏列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberProductCollectionListData
	for _, item := range *all {

		list = append(list, &umsclient.MemberProductCollectionListData{
			Id:              item.Id,
			MemberId:        item.MemberId,
			MemberNickName:  item.MemberNickName,
			MemberIcon:      item.MemberIcon,
			ProductId:       in.ProductId,
			ProductName:     item.ProductName,
			ProductPic:      item.ProductPic,
			ProductSubTitle: item.ProductSubTitle.String,
			ProductPrice:    item.ProductPrice,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员收藏列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &umsclient.MemberProductCollectionListResp{
		Total: count,
		List:  list,
	}, nil
}
