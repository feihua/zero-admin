package productvertifyrecordservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductVertifyRecordLogic 更新商品审核记录
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:07
*/
type UpdateProductVertifyRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductVertifyRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductVertifyRecordLogic {
	return &UpdateProductVertifyRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductVertifyRecord 更新商品审核记录
func (l *UpdateProductVertifyRecordLogic) UpdateProductVertifyRecord(in *pmsclient.UpdateProductVertifyRecordReq) (*pmsclient.UpdateProductVertifyRecordResp, error) {
	q := query.PmsProductVertifyRecord.WithContext(l.ctx)

	// 1.根据商品审核记录id查询商品审核记录是否已存在
	_, err := q.Where(query.PmsProductVertifyRecord.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据商品审核记录id：%d,查询商品审核记录失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询商品审核记录失败"))
	}

	item := &model.PmsProductVertifyRecord{
		ID:        in.Id,        //
		ProductID: in.ProductId, // 商品id
		ReviewMan: in.ReviewMan, // 审核人
		Status:    in.Status,    // 审核状态：0->未通过；1->通过
		Detail:    in.Detail,    // 反馈详情
	}

	// 2.商品审核记录存在时,则直接更新商品审核记录
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品审核记录失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品审核记录失败")
	}

	return &pmsclient.UpdateProductVertifyRecordResp{}, nil
}
