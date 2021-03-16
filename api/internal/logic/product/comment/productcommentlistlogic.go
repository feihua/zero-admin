package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCommentListLogic {
	return ProductCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommentListLogic) ProductCommentList(req types.ListProductCommentReq) (*types.ListProductCommentResp, error) {
	resp, err := l.svcCtx.Pms.CommentList(l.ctx, &pmsclient.CommentListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtProductCommentData

	for _, item := range resp.List {
		list = append(list, &types.ListtProductCommentData{
			Id:               item.Id,
			ProductId:        item.ProductId,
			MemberNickName:   item.MemberNickName,
			ProductName:      item.ProductName,
			Star:             item.Star,
			MemberIp:         item.MemberIp,
			CreateTime:       item.CreateTime,
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

	return &types.ListProductCommentResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "",
	}, nil
}
