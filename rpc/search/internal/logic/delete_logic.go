package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"
	"strconv"
	"strings"

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

	if in.Ids == nil {
		query := `{"query": {"match_all": {}}}`
		_, err := l.svcCtx.ESClient.DeleteByQuery([]string{svc.IndexName}, strings.NewReader(query))
		if err != nil {
			return nil, fmt.Errorf("bulk delete error: %+v", err)
		}
		logx.Infof("删除商品所有索引成功")
		return &search.DeleteResp{}, nil
	}
	var buf bytes.Buffer
	for _, id := range in.Ids {
		meta := map[string]map[string]string{
			"delete": {"_index": svc.IndexName, "_id": strconv.FormatInt(id, 10)},
		}
		line, _ := json.Marshal(meta)
		buf.Write(line)
		buf.WriteByte('\n')
	}
	_, err := l.svcCtx.ESClient.Bulk(bytes.NewReader(buf.Bytes()), l.svcCtx.ESClient.Bulk.WithContext(l.ctx))
	if err != nil {
		return nil, fmt.Errorf("bulk delete error: %+v", err)
	}
	logx.Infof("删除商品es索引成功：%+v", buf.String())
	return &search.DeleteResp{}, nil
}
