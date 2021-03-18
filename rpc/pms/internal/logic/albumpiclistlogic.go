package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumPicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicListLogic {
	return &AlbumPicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicListLogic) AlbumPicList(in *pms.AlbumPicListReq) (*pms.AlbumPicListResp, error) {
	all, _ := l.svcCtx.PmsAlbumPicModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsAlbumPicModel.Count()

	var list []*pms.AlbumPicListData
	for _, item := range *all {

		list = append(list, &pms.AlbumPicListData{
			Id:      item.Id,
			AlbumId: item.AlbumId,
			Pic:     item.Pic,
		})
	}

	fmt.Println(list)
	return &pms.AlbumPicListResp{
		Total: count,
		List:  list,
	}, nil
}
