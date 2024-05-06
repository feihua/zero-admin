package subjectservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListByIdsLogic {
	return &SubjectListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectListByIdsLogic) SubjectListByIds(in *cmsclient.SubjectListByIdsReq) (*cmsclient.SubjectListResp, error) {
	q := query.CmsSubject
	subjects, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Find()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品品牌列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*cmsclient.SubjectListData
	for _, item := range subjects {

		list = append(list, &cmsclient.SubjectListData{
			Id:              item.ID,
			CategoryId:      item.CategoryID,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     *item.Description,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
		})
	}

	logc.Infof(l.ctx, "查询专题列表信息,参数：%s,响应：%s", in, list)

	return &cmsclient.SubjectListResp{
		Total: 0,
		List:  list,
	}, nil
}
