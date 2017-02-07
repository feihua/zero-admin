package collect

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddordeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddordeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddordeleteLogic {
	return AddordeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddordeleteLogic) Addordelete(req types.AddOrDeleteReq) (resp *types.AddOrDeleteResp, err error) {
	_, err = l.svcCtx.Pms.CollectAddOrDelete(l.ctx, &pmsclient.CollectAddOrDeleteReq{
		ValueID:  req.ValueID,
		MemberId: req.MemberId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加或者取消收藏失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddOrDeleteResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.AddOrDeleteResp{
		Errno:  0,
		Errmsg: "添加或者取消收藏成功",
	}, nil
}
