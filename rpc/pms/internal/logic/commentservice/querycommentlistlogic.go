package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCommentListLogic 查询商品评价表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:36
*/
type QueryCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCommentListLogic {
	return &QueryCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCommentList 查询商品评价表列表
func (l *QueryCommentListLogic) QueryCommentList(in *pmsclient.QueryCommentListReq) (*pmsclient.QueryCommentListResp, error) {
	q := query.PmsComment.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品评价列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.CommentListData
	for _, item := range result {

		list = append(list, &pmsclient.CommentListData{
			Id:               item.ID,
			ProductId:        item.ProductID,
			MemberNickName:   item.MemberNickName,
			ProductName:      item.ProductName,
			Star:             item.Star,
			MemberIp:         item.MemberIP,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			ShowStatus:       item.ShowStatus,
			ProductAttribute: item.ProductAttribute,
			CollectCount:     item.CollectCount,
			ReadCount:        item.ReadCount,
			Content:          item.Content,
			Pics:             item.Pics,
			MemberIcon:       item.MemberIcon,
			ReplayCount:      item.ReplayCount,
		})
	}

	return &pmsclient.QueryCommentListResp{
		Total: count,
		List:  list,
	}, nil

}
