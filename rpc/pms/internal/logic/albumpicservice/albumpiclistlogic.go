package albumpicservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AlbumPicListLogic) AlbumPicList(in *pmsclient.AlbumPicListReq) (*pmsclient.AlbumPicListResp, error) {
	q := query.PmsAlbumPic.WithContext(l.ctx)

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询相册图片列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.AlbumPicListData
	for _, item := range result {

		list = append(list, &pmsclient.AlbumPicListData{
			Id:      item.ID,
			AlbumId: item.AlbumID,
			Pic:     item.Pic,
		})
	}

	logc.Infof(l.ctx, "查询相册图片列表信息,参数：%+v,响应：%+v", in, list)
	return &pmsclient.AlbumPicListResp{
		Total: count,
		List:  list,
	}, nil
}
