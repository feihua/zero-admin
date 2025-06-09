package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSortLogic {
	return &UpdateProductSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductSort 更新排序
func (l *UpdateProductSortLogic) UpdateProductSort(in *pmsclient.UpdateProductSortReq) (*pmsclient.UpdateProductResp, error) {
	q := query.PmsProduct

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.Sort.Value(in.Sort))

	if err != nil {
		logc.Errorf(l.ctx, "更新商排序失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品排序失败")
	}

	return &pmsclient.UpdateProductResp{}, nil
}
