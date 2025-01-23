package preferredareaservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddPreferredAreaLogic 添加优选专区
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddPreferredAreaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreferredAreaLogic {
	return &AddPreferredAreaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPreferredArea 添加优选专区
func (l *AddPreferredAreaLogic) AddPreferredArea(in *cmsclient.AddPreferredAreaReq) (*cmsclient.AddPreferredAreaResp, error) {
	q := query.CmsPreferredArea

	item := &model.CmsPreferredArea{
		Name:       in.Name,       // 专区名称
		SubTitle:   in.SubTitle,   // 子标题
		Pic:        in.Pic,        // 展示图片
		Sort:       in.Sort,       // 排序
		ShowStatus: in.ShowStatus, // 显示状态：0->不显示；1->显示
		CreateBy:   in.CreateBy,   // 创建者
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加优选专区失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加优选专区失败")
	}

	logc.Infof(l.ctx, "添加优选专区成功,参数：%+v", in)
	return &cmsclient.AddPreferredAreaResp{}, nil
}
