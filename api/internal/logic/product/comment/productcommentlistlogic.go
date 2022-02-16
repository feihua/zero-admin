package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询商品评价列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询商品评价失败")
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
		Message:  "查询商品评价成功",
	}, nil
}
