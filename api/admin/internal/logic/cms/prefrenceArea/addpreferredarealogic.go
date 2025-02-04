package prefrenceArea

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddPreferredAreaLogic 添加优选专区
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type AddPreferredAreaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPreferredAreaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPreferredAreaLogic {
	return &AddPreferredAreaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddPreferredArea 添加优选专区
func (l *AddPreferredAreaLogic) AddPreferredArea(req *types.AddPreferredAreaReq) (resp *types.AddPreferredAreaResp, err error) {

	_, err = l.svcCtx.PreferredAreaService.AddPreferredArea(l.ctx, &cmsclient.AddPreferredAreaReq{
		Name:       req.Name,       //专区名称
		SubTitle:   req.SubTitle,   //子标题
		Pic:        req.Pic,        //展示图片
		Sort:       req.Sort,       //排序
		ShowStatus: req.ShowStatus, //显示状态：0->不显示；1->显示
		CreateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加优选专区失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddPreferredAreaResp{
		Code:    "000000",
		Message: "添加优选专区成功",
	}, nil
}
