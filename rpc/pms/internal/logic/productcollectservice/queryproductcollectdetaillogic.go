package productcollectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCollectDetailLogic 查询收藏详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type QueryProductCollectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCollectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCollectDetailLogic {
	return &QueryProductCollectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCollectDetail 查询收藏详情
func (l *QueryProductCollectDetailLogic) QueryProductCollectDetail(in *pmsclient.QueryProductCollectDetailReq) (*pmsclient.QueryProductCollectDetailResp, error) {
	// item, err := query.PmsProductCollect.WithContext(l.ctx).Where(query.PmsProductCollect.ID.Eq(in.Id)).First()
	//
	// if err != nil {
	// 	logc.Errorf(l.ctx, "查询收藏详情失败,参数:%+v,异常:%s", in, err.Error())
	// 	return nil, errors.New("查询收藏详情失败")
	// }
	//
	// data := &pmsclient.QueryProductCollectDetailResp{
	// 	Id:          item.ID,                                    //
	// 	UserId:      item.UserID,                                // 用户表的用户ID
	// 	ValueId:     item.ValueID,                               // 如果type=0，则是商品ID；如果type=1，则是专题ID
	// 	CollectType: item.CollectType,                           // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
	// 	AddTime:     item.AddTime.Format("2006-01-02 15:04:05"), // 创建时间
	// 	Deleted:     item.Deleted,                               // 逻辑删除
	// }
	//
	// logc.Infof(l.ctx, "查询收藏详情,参数：%+v,响应：%+v", in, data)
	return nil, nil
}
