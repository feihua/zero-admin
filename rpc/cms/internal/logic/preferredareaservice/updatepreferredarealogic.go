package preferredareaservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePreferredAreaLogic 更新优选专区
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdatePreferredAreaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePreferredAreaLogic {
	return &UpdatePreferredAreaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePreferredArea 更新优选专区
func (l *UpdatePreferredAreaLogic) UpdatePreferredArea(in *cmsclient.UpdatePreferredAreaReq) (*cmsclient.UpdatePreferredAreaResp, error) {
	q := query.CmsPreferredArea.WithContext(l.ctx)

	// 1.根据优选专区id查询优选专区是否已存在
	_, err := q.Where(query.CmsPreferredArea.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据优选专区id：%d,查询优选专区失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询优选专区失败"))
	}

	item := &model.CmsPreferredArea{
		ID:         in.Id,         // 主键ID
		Name:       in.Name,       // 专区名称
		SubTitle:   in.SubTitle,   // 子标题
		Pic:        in.Pic,        // 展示图片
		Sort:       in.Sort,       // 排序
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   in.UpdateBy,   // 更新者
	}

	// 2.优选专区存在时,则直接更新优选专区
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新优选专区失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新优选专区失败")
	}

	logc.Infof(l.ctx, "更新优选专区成功,参数：%+v", in)
	return &cmsclient.UpdatePreferredAreaResp{}, nil
}
