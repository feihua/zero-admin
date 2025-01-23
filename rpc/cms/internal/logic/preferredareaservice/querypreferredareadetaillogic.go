package preferredareaservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryPreferredAreaDetailLogic 查询优选专区详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryPreferredAreaDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPreferredAreaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaDetailLogic {
	return &QueryPreferredAreaDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryPreferredAreaDetail 查询优选专区详情
func (l *QueryPreferredAreaDetailLogic) QueryPreferredAreaDetail(in *cmsclient.QueryPreferredAreaDetailReq) (*cmsclient.QueryPreferredAreaDetailResp, error) {
	item, err := query.CmsPreferredArea.WithContext(l.ctx).Where(query.CmsPreferredArea.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询优选专区详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优选专区详情失败")
	}

	data := &cmsclient.QueryPreferredAreaDetailResp{
		Id:         item.ID,                                 // 主键ID
		Name:       item.Name,                               // 专区名称
		SubTitle:   item.SubTitle,                           // 子标题
		Pic:        item.Pic,                                // 展示图片
		Sort:       item.Sort,                               // 排序
		ShowStatus: item.ShowStatus,                         // 显示状态：0->不显示；1->显示
		CreateBy:   item.CreateBy,                           // 创建者
		CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:   item.UpdateBy,                           // 更新者
		UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	logc.Infof(l.ctx, "查询优选专区详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
