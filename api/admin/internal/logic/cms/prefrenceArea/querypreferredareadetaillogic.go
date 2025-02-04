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

// QueryPreferredAreaDetailLogic 查询优选专区详情
/*
Author: 刘飞华
Date: 2025/02/04 14:56:41
*/
type QueryPreferredAreaDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPreferredAreaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaDetailLogic {
	return &QueryPreferredAreaDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryPreferredAreaDetail 查询优选专区详情
func (l *QueryPreferredAreaDetailLogic) QueryPreferredAreaDetail(req *types.QueryPreferredAreaDetailReq) (resp *types.QueryPreferredAreaDetailResp, err error) {

	detail, err := l.svcCtx.PreferredAreaService.QueryPreferredAreaDetail(l.ctx, &cmsclient.QueryPreferredAreaDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询优选专区详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryPreferredAreaDetailData{
		Id:         detail.Id,         //主键ID
		Name:       detail.Name,       //专区名称
		SubTitle:   detail.SubTitle,   //子标题
		Pic:        detail.Pic,        //展示图片
		Sort:       detail.Sort,       //排序
		ShowStatus: detail.ShowStatus, //显示状态：0->不显示；1->显示
		CreateBy:   detail.CreateBy,   //创建者
		CreateTime: detail.CreateTime, //创建时间
		UpdateBy:   detail.UpdateBy,   //更新者
		UpdateTime: detail.UpdateTime, //更新时间
	}
	return &types.QueryPreferredAreaDetailResp{
		Code:    "000000",
		Message: "查询优选专区成功",
		Data:    data,
	}, nil
}
