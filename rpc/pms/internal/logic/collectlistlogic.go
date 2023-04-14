package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/rpc/pms/pms"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectListLogic {
	return &CollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectListLogic) CollectList(in *pmsclient.CollectListReq) (*pmsclient.CollectListResp, error) {
	collectList, err := l.svcCtx.PmsCollectModel.FindAll(in.MemberId)
	count, _ := l.svcCtx.PmsCollectModel.Count(in.MemberId)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询收藏列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var listData []*pms.CollectListData
	for _, item := range *collectList {
		product, _ := l.svcCtx.PmsProductModel.FindOne(l.ctx, in.MemberId)
		listData = append(listData, &pms.CollectListData{
			Id:          item.Id,
			Type:        item.CollectType,
			ValueID:     item.ValueId,
			Name:        product.Name,
			Brief:       product.Description,
			PicUrl:      product.Pic,
			RetailPrice: fmt.Sprintf("%.2f", product.Price),
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(listData)
	logx.WithContext(l.ctx).Infof("查询收藏列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.CollectListResp{
		Total: count,
		List:  listData,
	}, nil
}
