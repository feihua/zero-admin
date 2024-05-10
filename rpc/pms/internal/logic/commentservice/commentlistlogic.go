package commentservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *pmsclient.CommentListReq) (*pmsclient.CommentListResp, error) {
	q := query.PmsComment.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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

	logc.Infof(l.ctx, "查询商品评价列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.CommentListResp{
		Total: count,
		List:  list,
	}, nil
}
