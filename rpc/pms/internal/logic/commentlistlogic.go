package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *CommentListLogic) CommentList(in *pms.CommentListReq) (*pms.CommentListResp, error) {
	all, err := l.svcCtx.PmsCommentModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsCommentModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询商品评价列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*pms.CommentListData
	for _, item := range *all {

		list = append(list, &pms.CommentListData{
			Id:               item.Id,
			ProductId:        item.ProductId,
			MemberNickName:   item.MemberNickName,
			ProductName:      item.ProductName,
			Star:             item.Star,
			MemberIp:         item.MemberIp,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			ShowStatus:       item.ShowStatus,
			ProductAttribute: item.ProductAttribute,
			CollectCouont:    item.CollectCouont,
			ReadCount:        item.ReadCount,
			Content:          item.Content,
			Pics:             item.Pics,
			MemberIcon:       item.MemberIcon,
			ReplayCount:      item.ReplayCount,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询商品评价列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &pms.CommentListResp{
		Total: count,
		List:  list,
	}, nil
}
