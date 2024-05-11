package preferredareaservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreferredAreaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreferredAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreferredAreaListLogic {
	return &PreferredAreaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreferredAreaListLogic) PreferredAreaList(in *cmsclient.PreferredAreaListReq) (*cmsclient.PreferredAreaListResp, error) {
	q := query.CmsPreferredArea.WithContext(l.ctx)
	preferredAreas, err := q.Find()
	count, _ := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*cmsclient.PreferredAreaListData
	for _, item := range preferredAreas {

		list = append(list, &cmsclient.PreferredAreaListData{
			Id:         item.ID,
			Name:       item.Name,
			SubTitle:   item.SubTitle,
			Pic:        item.Pic,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
		})
	}

	logc.Infof(l.ctx, "查询专题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.PreferredAreaListResp{
		Total: count,
		List:  list,
	}, nil
}
