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
	"gorm.io/gorm"
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
	q := query.CmsPreferredArea
	item, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("优选专区不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优选专区异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优选专区异常")
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

	return data, nil
}
