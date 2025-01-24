package productcollectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductCollectLogic 删除收藏
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type DeleteProductCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCollectLogic {
	return &DeleteProductCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductCollect 删除收藏
func (l *DeleteProductCollectLogic) DeleteProductCollect(in *pmsclient.DeleteProductCollectReq) (*pmsclient.DeleteProductCollectResp, error) {
	// q := query.PmsProductCollect

	// _, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	//
	// if err != nil {
	// 	logc.Errorf(l.ctx, "删除收藏失败,参数:%+v,异常:%s", in, err.Error())
	// 	return nil, errors.New("删除收藏失败")
	// }
	//
	// logc.Infof(l.ctx, "删除收藏成功,参数：%+v", in)
	return &pmsclient.DeleteProductCollectResp{}, nil
}
