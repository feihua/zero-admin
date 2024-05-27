package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AlbumListLogic) AlbumList(in *pmsclient.AlbumListReq) (*pmsclient.AlbumListResp, error) {
	q := query.PmsAlbum.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询相册列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.AlbumListData
	for _, item := range result {

		list = append(list, &pmsclient.AlbumListData{
			Id:          item.ID,
			Name:        item.Name,
			CoverPic:    item.CoverPic,
			PicCount:    item.PicCount,
			Sort:        item.Sort,
			Description: item.Description,
		})
	}

	logc.Infof(l.ctx, "查询相册列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.AlbumListResp{
		Total: count,
		List:  list,
	}, nil
}
