package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductVertifyRecordListLogic 查询商品审核记录列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:13
*/
type QueryProductVertifyRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductVertifyRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductVertifyRecordListLogic {
	return &QueryProductVertifyRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductVertifyRecordList 查询商品审核记录列表
func (l *QueryProductVertifyRecordListLogic) QueryProductVertifyRecordList(in *pmsclient.QueryProductVertifyRecordListReq) (*pmsclient.QueryProductVertifyRecordListResp, error) {
	q := query.PmsProductVertifyRecord.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品审核列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductVertifyRecordListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductVertifyRecordListData{
			Id:         item.ID,                              //
			ProductId:  item.ProductID,                       // 商品id
			CreateTime: time_util.TimeToStr(item.CreateTime), // 创建时间
			ReviewMan:  item.ReviewMan,                       // 审核人
			Status:     item.Status,                          // 审核状态：0->未通过；1->通过
			Detail:     item.Detail,                          // 反馈详情
		})
	}

	return &pmsclient.QueryProductVertifyRecordListResp{
		Total: count,
		List:  list,
	}, nil

}
