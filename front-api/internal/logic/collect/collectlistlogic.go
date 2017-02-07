package collect

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CollectListLogic {
	return CollectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectListLogic) CollectList(req types.CollectListReq) (resp *types.CollectListResp, err error) {
	collectList, err := l.svcCtx.Pms.CollectList(l.ctx, &pmsclient.CollectListReq{
		MemberId: req.MemberId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询用户收藏列表失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.CollectListResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	var data []types.CollectListData

	for _, item := range collectList.List {
		data = append(data, types.CollectListData{
			ID:          item.Id,
			Type:        item.Type,
			ValueID:     item.ValueID,
			Name:        item.Name,
			Brief:       item.Brief,
			PicUrl:      item.PicUrl,
			RetailPrice: item.RetailPrice,
		})
	}

	return &types.CollectListResp{
		Errno:  0,
		Data:   data,
		Errmsg: "查询用户收藏列表成功",
	}, nil
}
