package prefrenceareaservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// PrefrenceAreaListLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 10:33
*/
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

// PrefrenceAreaList 查询商品优选列表
func (l *PrefrenceAreaListLogic) PrefrenceAreaList(in *cmsclient.PrefrenceAreaListReq) (*cmsclient.PrefrenceAreaListResp, error) {

	q := query.CmsPrefrenceArea.WithContext(l.ctx)
	prefrenceAreas, err := q.Find()
	count, _ := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*cmsclient.PrefrenceAreaListData
	for _, item := range prefrenceAreas {

		list = append(list, &cmsclient.PrefrenceAreaListData{
			Id:         item.ID,
			Name:       item.Name,
			SubTitle:   item.SubTitle,
			Pic:        item.Pic,
			Sort:       item.Sort,
			ShowStatus: item.ShowStatus,
		})
	}

	logc.Infof(l.ctx, "查询专题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.PrefrenceAreaListResp{
		Total: count,
		List:  list,
	}, nil

}
