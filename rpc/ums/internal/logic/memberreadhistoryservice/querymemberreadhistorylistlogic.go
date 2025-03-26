package memberreadhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
		q = q.Where(query.UmsMemberReadHistory.MemberID.Eq(in.MemberId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户商品浏览历史记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户商品浏览历史记录列表失败")
	}

	var list []*umsclient.MemberReadHistoryListData
	for _, item := range result {

		list = append(list, &umsclient.MemberReadHistoryListData{
			Id:              item.ID,                              //
			MemberId:        item.MemberID,                        // 会员id
			MemberNickName:  item.MemberNickName,                  // 会员姓名
			MemberIcon:      item.MemberIcon,                      // 会员头像
			ProductId:       item.ProductID,                       // 商品id
			ProductName:     item.ProductName,                     // 商品名称
			ProductPic:      item.ProductPic,                      // 商品图片
			ProductSubTitle: item.ProductSubTitle,                 // 商品标题
			ProductPrice:    item.ProductPrice,                    // 商品价格
			CreateTime:      time_util.TimeToStr(item.CreateTime), // 浏览时间
		})
	}

	return &umsclient.QueryMemberReadHistoryListResp{
		Total: count,
		List:  list,
	}, nil

}
