package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/cms/cms"
	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

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
	all, err := l.svcCtx.CmsSubjectModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.CmsSubjectModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询专题列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*cms.SubjectListData
	for _, item := range *all {

		list = append(list, &cms.SubjectListData{
			Id:              item.Id,
			CategoryId:      item.CategoryId.Int64,
			Title:           item.Title.String,
			Pic:             item.Pic.String,
			ProductCount:    item.ProductCount.Int64,
			RecommendStatus: item.RecommendStatus.Int64,
			CreateTime:      item.CreateTime.Time.Format("2006-01-02 15:04:05"),
			CollectCount:    item.CollectCount.Int64,
			ReadCount:       item.ReadCount.Int64,
			CommentCount:    item.CommentCount.Int64,
			AlbumPics:       item.AlbumPics.String,
			Description:     item.Description.String,
			ShowStatus:      item.ShowStatus.Int64,
			Content:         item.Content.String,
			ForwardCount:    item.ForwardCount.Int64,
			CategoryName:    item.CategoryName.String,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询专题列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &cmsclient.SubjectListResp{
		Total: count,
		List:  list,
	}, nil
}
