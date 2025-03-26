package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateVerifyStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:12
*/
type UpdateVerifyStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVerifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVerifyStatusLogic {
	return &UpdateVerifyStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateVerifyStatus 批量修改审核状态
func (l *UpdateVerifyStatusLogic) UpdateVerifyStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.VerifyStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量修改审核状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量修改审核状态失败")
	}

	// 修改完审核状态后插入审核记录
	var list []*model.PmsProductVertifyRecord
	for _, id := range in.Ids {
		list = append(list, &model.PmsProductVertifyRecord{
			ProductID: id,            // 商品id
			ReviewMan: in.VertifyMan, // 审核人
			Status:    in.Status,     // 审核状态：0->未通过；1->通过
			Detail:    in.Detail,     // 反馈详情
		})
	}
	err = query.PmsProductVertifyRecord.WithContext(l.ctx).CreateInBatches(list, len(list))
	if err != nil {
		logc.Errorf(l.ctx, "批量修改审核状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量修改审核状态失败")
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}
