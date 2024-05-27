package productvertifyrecordservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductVertifyRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordListLogic {
	return &ProductVertifyRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordListLogic) ProductVertifyRecordList(in *pmsclient.ProductVertifyRecordListReq) (*pmsclient.ProductVertifyRecordListResp, error) {
	q := query.PmsProductVertifyRecord.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品审核列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductVertifyRecordListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductVertifyRecordListData{
			Id:         item.ID,
			ProductId:  item.ProductID,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			VertifyMan: item.VertifyMan,
			Status:     item.Status,
			Detail:     item.Detail,
		})
	}

	logc.Infof(l.ctx, "查询商品审核列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.ProductVertifyRecordListResp{
		Total: count,
		List:  list,
	}, nil
}
