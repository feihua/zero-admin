package subjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectListLogic {
	return &SubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectListLogic) SubjectList(in *cmsclient.SubjectListReq) (*cmsclient.SubjectListResp, error) {
	q := query.CmsSubject.WithContext(l.ctx)
	subjects, err := q.Find()
	count, _ := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询专题列表信息失败,参数:%+v,异常:%s", in, err.Error())
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

	logc.Infof(l.ctx, "查询专题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.SubjectListResp{
		Total: count,
		List:  list,
	}, nil
}
