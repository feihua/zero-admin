package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create 创建商品
func (l *CreateLogic) Create(in *search.CreateReq) (*search.CreateResp, error) {
	var buf bytes.Buffer
	for _, p := range in.Data {
		meta := map[string]map[string]string{
			"index": {"_index": svc.IndexName, "_id": strconv.FormatInt(p.Id, 10)},
		}
		metaLine, _ := json.Marshal(meta)
		buf.Write(metaLine)
		buf.WriteByte('\n')
		data, _ := json.Marshal(p)
		buf.Write(data)
		buf.WriteByte('\n')
	}
	_, err := l.svcCtx.ESClient.Bulk(bytes.NewReader(buf.Bytes()), l.svcCtx.ESClient.Bulk.WithContext(l.ctx))
	if err != nil {
		return nil, fmt.Errorf("bulk create error: %+v", err)
	}

	return &search.CreateResp{}, nil
}
