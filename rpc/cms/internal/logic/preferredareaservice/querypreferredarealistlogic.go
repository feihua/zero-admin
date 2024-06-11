package preferredareaservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPreferredAreaListLogic 查询优选专区列表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:45
*/
type QueryPreferredAreaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPreferredAreaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaListLogic {
	return &QueryPreferredAreaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryPreferredAreaList 查询优选专区列表
func (l *QueryPreferredAreaListLogic) QueryPreferredAreaList(in *cmsclient.QueryPreferredAreaListReq) (*cmsclient.QueryPreferredAreaListResp, error) {
	q := query.CmsPreferredArea.WithContext(l.ctx)

	if len(in.Name) > 0 {
		q = q.Where(query.CmsPreferredArea.Name.Like("%" + in.Name + "%"))
	}
	if len(in.SubTitle) > 0 {
		q = q.Where(query.CmsPreferredArea.SubTitle.Like("%" + in.SubTitle + "%"))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.CmsPreferredArea.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品优选列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*cmsclient.PreferredAreaListData

	for _, item := range result {
		preferredArea := &cmsclient.PreferredAreaListData{}
		preferredArea.Id = item.ID
		preferredArea.Name = item.Name
		preferredArea.SubTitle = item.SubTitle
		preferredArea.Pic = item.Pic
		preferredArea.Sort = item.Sort
		preferredArea.ShowStatus = item.ShowStatus
		preferredArea.CreateBy = item.CreateBy
		preferredArea.CreateTime = item.CreateTime.Format("2006-01-02 15:04:05")

		if item.UpdateBy != nil {
			preferredArea.UpdateBy = *item.UpdateBy
		}
		if item.UpdateTime != nil {
			preferredArea.UpdateTime = item.UpdateTime.Format("2006-01-02 15:04:05")
		}

		list = append(list, preferredArea)
	}

	logc.Infof(l.ctx, "查询商品优选列表信息,参数：%+v,响应：%+v", in, list)
	return &cmsclient.QueryPreferredAreaListResp{
		Total: count,
		List:  list,
	}, nil

}
