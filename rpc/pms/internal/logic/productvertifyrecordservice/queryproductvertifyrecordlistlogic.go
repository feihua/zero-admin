package productvertifyrecordservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
	result, err := l.svcCtx.ProductVertifyRecordModel.FindAll(l.ctx, in.ProductId)

	if err != nil {
		logc.Errorf(l.ctx, "查询商品审核记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品审核记录列表失败")
	}

	var list []*pmsclient.ProductVertifyRecordListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductVertifyRecordListData{
			Id:         item.ID.Hex(),                      //
			ProductId:  item.ProductId,                     // 商品id
			CreateTime: time_util.TimeToStr(item.CreateAt), // 创建时间
			ReviewMan:  item.ReviewMan,                     // 审核人
			Status:     item.Status,                        // 审核状态：0->未通过；1->通过
			Detail:     item.Detail,                        // 反馈详情
		})
	}

	return &pmsclient.QueryProductVertifyRecordListResp{
		Total: 0,
		List:  list,
	}, nil

}
