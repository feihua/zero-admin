package productcollectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCollectListLogic 查询收藏列表
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type QueryProductCollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCollectListLogic {
	return &QueryProductCollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCollectList 查询收藏列表
func (l *QueryProductCollectListLogic) QueryProductCollectList(in *pmsclient.QueryProductCollectListReq) (*pmsclient.QueryProductCollectListResp, error) {
	result, err := l.svcCtx.ProductCollectModel.FindPage(l.ctx, in.UserId, in.PageNum, in.PageSize)

	if err != nil {
		logc.Errorf(l.ctx, "查询收藏列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询收藏列表失败")
	}

	var list []*pmsclient.ProductCollectListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductCollectListData{
			Id:          item.ID.Hex(),                      //
			UserId:      item.MemberId,                      // 用户表的用户ID
			ValueId:     item.ValueId,                       // 如果type=0，则是商品ID；如果type=1，则是专题ID
			CollectType: item.CollectType,                   // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
			AddTime:     time_util.TimeToStr(item.CreateAt), // 创建时间
			Deleted:     item.Deleted,                       // 逻辑删除

		})
	}

	return &pmsclient.QueryProductCollectListResp{
		Total: 0,
		List:  list,
	}, nil
}
