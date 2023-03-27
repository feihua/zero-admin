package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectAddOrDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectAddOrDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectAddOrDeleteLogic {
	return &CollectAddOrDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectAddOrDeleteLogic) CollectAddOrDelete(in *pmsclient.CollectAddOrDeleteReq) (*pmsclient.CollectAddOrDeleteResp, error) {
	collect, _ := l.svcCtx.PmsCollectModel.FindOne(in.MemberId, in.ValueID)

	if collect == nil {
		//不存在，则添加
		return insertCollect(in, l)
	} else {
		//存在，则删除
		return deleteCollect(l, collect)
	}

}

// 删除收藏
func deleteCollect(l *CollectAddOrDeleteLogic, collect *pmsmodel.PmsCollect) (*pmsclient.CollectAddOrDeleteResp, error) {
	err := l.svcCtx.PmsCollectModel.Delete(int64(collect.Id))
	return &pmsclient.CollectAddOrDeleteResp{
		Pong: "pong",
	}, err
}

// 添加收藏
func insertCollect(in *pmsclient.CollectAddOrDeleteReq, l *CollectAddOrDeleteLogic) (*pmsclient.CollectAddOrDeleteResp, error) {
	_, err := l.svcCtx.PmsCollectModel.Insert(pmsmodel.PmsCollect{
		UserId:      in.MemberId,
		ValueId:     in.ValueID,
		CollectType: 0,
		AddTime:     time.Now(),
		Deleted:     0,
	})
	return &pmsclient.CollectAddOrDeleteResp{
		Pong: "pong",
	}, err
}
