package productcollectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductCollectLogic 添加收藏
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type AddProductCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCollectLogic {
	return &AddProductCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductCollect 添加收藏
func (l *AddProductCollectLogic) AddProductCollect(in *pmsclient.AddProductCollectReq) (*pmsclient.AddProductCollectResp, error) {
	err := l.svcCtx.ProductCollectModel.Insert(l.ctx, &model.ProductCollect{
		MemberId:    in.UserId,      // 用户表的用户ID
		ValueId:     in.ValueId,     // 如果type=0，则是商品ID；如果type=1，则是专题ID
		CollectType: in.CollectType, // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
		Deleted:     in.Deleted,     // 逻辑删除
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加收藏失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加收藏失败")
	}

	return &pmsclient.AddProductCollectResp{}, nil
}
