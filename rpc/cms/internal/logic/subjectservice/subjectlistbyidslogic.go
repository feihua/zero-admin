package subjectservicelogic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

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
	all, err := l.svcCtx.CmsSubjectModel.FindAllByIds(l.ctx, in.Ids)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品品牌列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*cmsclient.SubjectListData
	for _, item := range *all {

		list = append(list, &cmsclient.SubjectListData{
			Id:              item.Id,
			CategoryId:      item.CategoryId,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CreateTime:      item.CreateTime.Format("2006-01-02 15:04:05"),
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     item.Description.String,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询专题列表信息,参数：%s,响应：%s", reqStr, listStr)

	return &cmsclient.SubjectListResp{
		Total: 0,
		List:  list,
	}, nil
}
