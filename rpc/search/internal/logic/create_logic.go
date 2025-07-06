package logic

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bytedance/sonic"
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

	var deleteBuf bytes.Buffer
	for _, p := range in.Data {

		deleteMeta := map[string]map[string]string{
			"delete": {"_index": svc.IndexName, "_id": strconv.FormatInt(p.Id, 10)},
		}
		line, _ := sonic.Marshal(deleteMeta)
		deleteBuf.Write(line)
		deleteBuf.WriteByte('\n')

		meta := map[string]map[string]string{
			"index": {"_index": svc.IndexName, "_id": strconv.FormatInt(p.Id, 10)},
		}
		metaLine, _ := sonic.Marshal(meta)
		buf.Write(metaLine)
		buf.WriteByte('\n')
		data, _ := sonic.Marshal(p)
		buf.Write(data)
		buf.WriteByte('\n')
	}

	_, err := l.svcCtx.ESClient.Bulk(bytes.NewReader(deleteBuf.Bytes()), l.svcCtx.ESClient.Bulk.WithContext(l.ctx))
	if err != nil {
		logx.Infof("bulk delete error：%+v", err)
	} else {
		logx.Infof("删除商品es索引成功：%+v", deleteBuf.String())
	}

	_, err = l.svcCtx.ESClient.Bulk(bytes.NewReader(buf.Bytes()), l.svcCtx.ESClient.Bulk.WithContext(l.ctx))
	if err != nil {
		return nil, fmt.Errorf("bulk create error: %+v", err)
	} else {
		logx.Infof("添加商品es索引成功：%+v", buf.String())
	}

	return &search.CreateResp{}, nil
}
