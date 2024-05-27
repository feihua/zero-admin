package subjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SubjectListLogic 获取商品专题
/*
Author: LiuFeiHua
Date: 2024/5/13 9:48
*/
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

// SubjectList 获取商品专题
func (l *SubjectListLogic) SubjectList(in *cmsclient.SubjectListReq) (*cmsclient.SubjectListResp, error) {
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

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

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

	return &cmsclient.SubjectListResp{
		Total: count,
		List:  list,
	}, nil
}
