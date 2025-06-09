package memberproductcollectionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberProductCollectionListLogic 查询用户收藏的商品列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:09
*/
type QueryMemberProductCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberProductCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberProductCollectionListLogic {
	return &QueryMemberProductCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberProductCollectionList 查询用户收藏的商品列表
func (l *QueryMemberProductCollectionListLogic) QueryMemberProductCollectionList(in *umsclient.QueryMemberProductCollectionListReq) (*umsclient.QueryMemberProductCollectionListResp, error) {
	result, err := l.svcCtx.MemberProductCollectionModel.FindPage(l.ctx, in.MemberId, in.PageNum, in.PageSize)

	if err != nil {
		logc.Errorf(l.ctx, "查询用户收藏的商品列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户收藏的商品列表失败")
	}

	var list []*umsclient.MemberProductCollectionListData
	for _, item := range result {

		list = append(list, &umsclient.MemberProductCollectionListData{
			Id:              item.ID.Hex(),                      //
			MemberId:        item.MemberId,                      // 会员id
			MemberNickName:  item.MemberNickName,                // 会员姓名
			MemberIcon:      item.MemberIcon,                    // 会员头像
			ProductId:       item.ProductId,                     // 商品id
			ProductName:     item.ProductName,                   // 商品名称
			ProductPic:      item.ProductPic,                    // 商品图片
			ProductSubTitle: item.ProductSubTitle,               // 商品标题
			ProductPrice:    item.ProductPrice,                  // 商品价格
			CreateTime:      time_util.TimeToStr(item.CreateAt), // 收藏时间
		})
	}

	return &umsclient.QueryMemberProductCollectionListResp{
		List: list,
	}, nil
}
