package product

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/rpc/pms/client/productspuservice"
	"github.com/feihua/zero-admin/rpc/search/search_client"
	"github.com/zeromicro/go-zero/core/logc"
)

// DeleteProductFromEs 删除es中商品的索引
func DeleteProductFromEs(ctx context.Context, body []byte, Search search_client.Search, productSpuService productspuservice.ProductSpuService) {
	logc.Infof(ctx, "需要删除es中商品的索引信息: %s", body)
	var orderInfo map[string][]int64
	err := sonic.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}
	ids := orderInfo["ids"]

	_, err = Search.Delete(ctx, &search_client.DeleteReq{
		Ids: ids,
	})

	if err != nil {
		logc.Errorf(ctx, "删除es中商品的索引信息失败,请求参数：%s,错误信息：%+v", body, err)
		return
	}

	logc.Errorf(ctx, "删除es中商品的索引信息成功,请求参数：%s", body)
}
