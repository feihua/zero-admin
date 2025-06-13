package productvertifyrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductVertifyRecordLogic 添加商品审核记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:12
*/
type AddProductVertifyRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductVertifyRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductVertifyRecordLogic {
	return &AddProductVertifyRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddProductVertifyRecord 添加商品审核记录
func (l *AddProductVertifyRecordLogic) AddProductVertifyRecord(in *pmsclient.AddProductVertifyRecordReq) (*pmsclient.AddProductVertifyRecordResp, error) {
	err := l.svcCtx.ProductVertifyRecordModel.Insert(l.ctx, &model.ProductVertifyRecord{
		ProductId: in.ProductId, // 商品id
		ReviewMan: in.ReviewMan, // 审核人
		Status:    in.Status,    // 审核状态：0->未通过；1->通过
		Detail:    in.Detail,    // 反馈详情
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品审核记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加商品审核记录失败")
	}

	return &pmsclient.AddProductVertifyRecordResp{}, nil
}
