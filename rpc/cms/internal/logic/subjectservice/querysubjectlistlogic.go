package subjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
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
		list = append(list, &cmsclient.SubjectListData{
			Id:              item.ID,                                 // 专题id
			CategoryId:      item.CategoryID,                         // 专题分类id
			Title:           item.Title,                              // 专题标题
			Pic:             item.Pic,                                // 专题主图
			ProductCount:    item.ProductCount,                       // 关联产品数量
			RecommendStatus: item.RecommendStatus,                    // 推荐状态：0->不推荐；1->推荐
			CollectCount:    item.CollectCount,                       // 收藏数
			ReadCount:       item.ReadCount,                          // 阅读数
			CommentCount:    item.CommentCount,                       // 评论数
			AlbumPics:       item.AlbumPics,                          // 画册图片用逗号分割
			Description:     item.Description,                        // 专题内容
			ShowStatus:      item.ShowStatus,                         // 显示状态：0->不显示；1->显示
			Content:         item.Content,                            // 专题内容
			ForwardCount:    item.ForwardCount,                       // 转发数
			CategoryName:    item.CategoryName,                       // 专题分类名称
			CreateBy:        item.CreateBy,                           // 创建者
			CreateTime:      time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:        item.UpdateBy,                           // 更新者
			UpdateTime:      time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	logc.Infof(l.ctx, "查询专题列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.QuerySubjectListResp{
		Total: count,
		List:  list,
	}, nil

}
