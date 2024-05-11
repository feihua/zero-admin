package memberpriceservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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

	q := query.PmsMemberPrice.WithContext(l.ctx).Where(query.PmsComment.ProductID.Eq(123))

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询会员价格列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.MemberPriceListData
	for _, item := range result {
		list = append(list, &pmsclient.MemberPriceListData{
			Id:              item.ID,
			ProductId:       item.ProductID,
			MemberLevelId:   item.MemberLevelID,
			MemberPrice:     item.MemberPrice,
			MemberLevelName: item.MemberLevelName,
		})
	}

	logc.Infof(l.ctx, "查询会员价格列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.MemberPriceListResp{
		Total: count,
		List:  list,
	}, nil
}
