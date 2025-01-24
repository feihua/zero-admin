package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductVertifyRecordStatusLogic 更新商品审核记录
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:07
*/
type UpdateProductVertifyRecordStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductVertifyRecordStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductVertifyRecordStatusLogic {
	return &UpdateProductVertifyRecordStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductVertifyRecordStatus 更新商品审核记录状态
func (l *UpdateProductVertifyRecordStatusLogic) UpdateProductVertifyRecordStatus(in *pmsclient.UpdateProductVertifyRecordStatusReq) (*pmsclient.UpdateProductVertifyRecordStatusResp, error) {
	// q := query.PmsProductVertifyRecord
	//
	// _, err := q.WithContext(l.ctx).Where(q.ID.In(in.Id...)).Update(q.Status, in.Status)
	//
	// if err != nil {
	// 	logc.Errorf(l.ctx, "更新商品审核记录状态失败,参数:%+v,异常:%s", in, err.Error())
	// 	return nil, errors.New("更新商品审核记录状态失败")
	// }
	//
	// logc.Infof(l.ctx, "更新商品审核记录状态成功,参数：%+v", in)
	return &pmsclient.UpdateProductVertifyRecordStatusResp{}, nil
}
