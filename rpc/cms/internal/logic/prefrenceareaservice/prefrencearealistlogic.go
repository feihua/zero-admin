package prefrenceareaservicelogic

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaListLogic {
	return &PrefrenceAreaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaListLogic) PrefrenceAreaList(in *cmsclient.PrefrenceAreaListReq) (*cmsclient.PrefrenceAreaListResp, error) {
	all, err := l.svcCtx.CmsPrefrenceAreaModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.CmsPrefrenceAreaModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品优选列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*cmsclient.PrefrenceAreaListData
	for _, item := range *all {

		list = append(list, &cmsclient.PrefrenceAreaListData{
			Id:         item.Id,
			Name:       item.Name,
			SubTitle:   item.SubTitle,
			Pic:        item.Pic,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询专题列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &cmsclient.PrefrenceAreaListResp{
		Total: count,
		List:  list,
	}, nil

}
