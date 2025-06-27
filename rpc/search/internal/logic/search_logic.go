package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/pkg/es"

	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Search 综合搜索、筛选、排序-根据关键字通过名称或副标题复合查询商品
func (l *SearchLogic) Search(in *search.SearchReq) (*search.SearchResp, error) {
	sortField := []interface{}{}
	switch in.Sort {
	case 1: // 按新品
		sortField = append(sortField, map[string]interface{}{"new_status_sort": map[string]string{"order": "desc"}})
	case 2: // 按销量
		sortField = append(sortField, map[string]interface{}{"sales": map[string]string{"order": "desc"}})
	case 3: // 价格从低到高
		sortField = append(sortField, map[string]interface{}{"price": map[string]string{"order": "asc"}})
	case 4: // 价格从高到低
		sortField = append(sortField, map[string]interface{}{"price": map[string]string{"order": "desc"}})
	default: // 相关度（默认）
		// 不加 sort，ES 默认按相关度排序
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":  in.Keyword,
							"fields": []string{"name", "brief", "keyword"},
						},
					},
				},
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{"category_id": in.CategoryId},
					},
					map[string]interface{}{
						"term": map[string]interface{}{"brand_id": in.BrandId},
					},
				},
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

	if len(sortField) > 0 {
		query["sort"] = sortField
	}

	data, _ := json.Marshal(query)
	res, err := es.GetESClient().Search(
		es.GetESClient().Search.WithContext(l.ctx),
		es.GetESClient().Search.WithIndex(svc.IndexName),
		es.GetESClient().Search.WithBody(bytes.NewReader(data)),
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
