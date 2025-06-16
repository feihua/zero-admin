package productspuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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

// UpdateVerifyStatus 修改审核状态
func (l *UpdateVerifyStatusLogic) UpdateVerifyStatus(in *pmsclient.UpdateProductSpuStatusReq) (*pmsclient.UpdateProductSpuStatusResp, error) {
	q := query.PmsProductSpu
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.VerifyStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "批量修改审核状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量修改审核状态失败")
	}

	// 修改完审核状态后插入审核记录
	for _, id := range in.Ids {
		err = l.svcCtx.ProductVertifyRecordModel.Insert(l.ctx, &model.ProductVertifyRecord{
			ProductId: id,           // 商品id
			ReviewMan: in.ReviewMan, // 审核人
			Status:    in.Status,    // 审核状态：0->未通过；1->通过
			Detail:    in.Detail,    // 反馈详情
		})

		if err != nil {
			logc.Errorf(l.ctx, "批量修改审核状态失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("批量修改审核状态失败")
		}
	}

	return &pmsclient.UpdateProductSpuStatusResp{}, nil
}
