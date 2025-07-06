package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSimpleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSimpleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSimpleLogic {
	return &SearchSimpleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchSimple 简单搜索-根据关键字通过名称或副标题查询商品
func (l *SearchSimpleLogic) SearchSimple(in *search.SearchSimpleReq) (*search.SearchResp, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  in.Keyword,
				"fields": []string{"name", "brief"},
			},
		},
		"from": (in.PageNum - 1) * in.PageSize,
		"size": in.PageSize,
		"highlight": map[string]interface{}{
			"fields": map[string]interface{}{
				"name":  map[string]interface{}{},
				"brief": map[string]interface{}{},
			},
		},
	}

	data, _ := sonic.Marshal(query)
	res, err := l.svcCtx.ESClient.Search(
		l.svcCtx.ESClient.Search.WithContext(l.ctx),
		l.svcCtx.ESClient.Search.WithIndex(svc.IndexName),
		l.svcCtx.ESClient.Search.WithBody(bytes.NewReader(data)),
	)
	if err != nil {
		return nil, err
	}
	var result struct {
		Hits struct {
			Hits []struct {
				Source    search.ProductData  `json:"_source"`
				Highlight map[string][]string `json:"highlight"`
			} `json:"hits"`
		} `json:"hits"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	products := make([]*search.ProductData, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		p := hit.Source
		// 可处理高亮字段
		products = append(products, &p)
	}

	return &search.SearchResp{
		Data:  products,
		Total: 0,
	}, nil
}
