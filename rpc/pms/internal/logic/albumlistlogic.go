package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumListLogic {
	return &AlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumListLogic) AlbumList(in *pms.AlbumListReq) (*pms.AlbumListResp, error) {
	all, _ := l.svcCtx.PmsAlbumModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*pms.AlbumListData
	for _, item := range *all {

		list = append(list, &pms.AlbumListData{
			Id:          item.Id,
			Name:        item.Name,
			CoverPic:    item.CoverPic,
			PicCount:    item.PicCount,
			Sort:        item.Sort,
			Description: item.Description,
		})
	}

	fmt.Println(list)
	return &pms.AlbumListResp{
		Total: 10,
		List:  list,
	}, nil
}
