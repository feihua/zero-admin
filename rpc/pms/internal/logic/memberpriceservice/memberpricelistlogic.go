package memberpriceservicelogic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceListLogic {
	return &MemberPriceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceListLogic) MemberPriceList(in *pmsclient.MemberPriceListReq) (*pmsclient.MemberPriceListResp, error) {
	all, err := l.svcCtx.PmsMemberPriceModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsMemberPriceModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员价格列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pmsclient.MemberPriceListData
	for _, item := range *all {
		list = append(list, &pmsclient.MemberPriceListData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			MemberLevelId:   item.MemberLevelId,
			MemberPrice:     int64(item.MemberPrice),
			MemberLevelName: item.MemberLevelName,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员价格列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pmsclient.MemberPriceListResp{
		Total: count,
		List:  list,
	}, nil
}
