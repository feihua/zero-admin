package subjectservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/pkg/time_util"
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

func (l *SubjectListByIdsLogic) SubjectListByIds(in *cmsclient.SubjectListByIdsReq) (*cmsclient.QuerySubjectListResp, error) {
	q := query.CmsSubject
	subjects, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Find()

	if err != nil {
		in, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询商品品牌列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*cmsclient.SubjectListData
	for _, item := range subjects {

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

	return &cmsclient.QuerySubjectListResp{
		Total: 0,
		List:  list,
	}, nil
}
