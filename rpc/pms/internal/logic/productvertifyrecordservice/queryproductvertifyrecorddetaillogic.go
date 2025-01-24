package productvertifyrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductVertifyRecordDetailLogic 查询商品审核记录详情
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:07
*/
type QueryProductVertifyRecordDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductVertifyRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductVertifyRecordDetailLogic {
	return &QueryProductVertifyRecordDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductVertifyRecordDetail 查询商品审核记录详情
func (l *QueryProductVertifyRecordDetailLogic) QueryProductVertifyRecordDetail(in *pmsclient.QueryProductVertifyRecordDetailReq) (*pmsclient.QueryProductVertifyRecordDetailResp, error) {
	item, err := query.PmsProductVertifyRecord.WithContext(l.ctx).Where(query.PmsProductVertifyRecord.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品审核记录详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品审核记录详情失败")
	}

	data := &pmsclient.QueryProductVertifyRecordDetailResp{
		Id:         item.ID,                              //
		ProductId:  item.ProductID,                       // 商品id
		CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
		ReviewMan:  item.ReviewMan,                       // 审核人
		Status:     item.Status,                          // 审核状态：0->未通过；1->通过
		Detail:     item.Detail,                          // 反馈详情
	}

	logc.Infof(l.ctx, "查询商品审核记录详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
