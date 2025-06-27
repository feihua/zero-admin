package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/pkg/es"
	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delete 根据ids删除商品
func (l *DeleteLogic) Delete(in *search.DeleteReq) (*search.DeleteResp, error) {

	var buf bytes.Buffer
	for _, id := range in.Ids {
		meta := map[string]map[string]string{
			"delete": {"_index": svc.IndexName, "_id": id},
		}
		line, _ := json.Marshal(meta)
		buf.Write(line)
		buf.WriteByte('\n')
	}
	_, err := es.GetESClient().Bulk(bytes.NewReader(buf.Bytes()), es.GetESClient().Bulk.WithContext(l.ctx))
	if err != nil {
		return nil, fmt.Errorf("bulk delete error: %+v", err)
	}

	return &search.DeleteResp{}, nil
}
