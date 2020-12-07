package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.PmsCommentModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

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

	fmt.Println(list)
	return &pms.CommentListResp{
		Total: 10,
		List:  list,
	}, nil
}
