package subjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectListLogic 查询专题表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 16:46
*/
type QuerySubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectListLogic {
	return &QuerySubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubjectList 查询专题表列表
func (l *QuerySubjectListLogic) QuerySubjectList(in *cmsclient.QuerySubjectListReq) (*cmsclient.QuerySubjectListResp, error) {
	q := query.CmsSubject.WithContext(l.ctx)
	if len(in.Title) > 0 {
		q = q.Where(query.CmsSubject.Title.Like("%" + in.Title + "%"))
	}
	if in.RecommendStatus != 2 {
		q = q.Where(query.CmsSubject.RecommendStatus.Eq(in.RecommendStatus))
	}
	if in.ShowStatus != 2 {
		q = q.Where(query.CmsSubject.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询专题列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*cmsclient.SubjectListData
	for _, item := range result {

		subject := &cmsclient.SubjectListData{}
		subject.Id = item.ID
		subject.CategoryId = item.CategoryID
		subject.Id = item.ID
		subject.CategoryId = item.CategoryID
		subject.Title = item.Title
		subject.Pic = item.Pic
		subject.ProductCount = item.ProductCount
		subject.RecommendStatus = item.RecommendStatus
		subject.CollectCount = item.CollectCount
		subject.ReadCount = item.ReadCount
		subject.CommentCount = item.CommentCount
		subject.AlbumPics = item.AlbumPics
		subject.Description = *item.Description
		subject.ShowStatus = item.ShowStatus
		subject.Content = item.Content
		subject.ForwardCount = item.ForwardCount
		subject.CategoryName = item.CategoryName
		subject.CreateBy = item.CreateBy
		subject.CreateTime = item.CreateTime.Format("2006-01-02 15:04:05")

		if item.UpdateBy != nil {
			subject.UpdateBy = *item.UpdateBy
		}
		if item.UpdateTime != nil {
			subject.UpdateTime = item.UpdateTime.Format("2006-01-02 15:04:05")
		}

		list = append(list, subject)
	}

	logc.Infof(l.ctx, "查询专题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.QuerySubjectListResp{
		Total: count,
		List:  list,
	}, nil

}
