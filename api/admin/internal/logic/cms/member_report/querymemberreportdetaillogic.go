package member_report

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

// QueryMemberReportDetailLogic 查询用户举报详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryMemberReportDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberReportDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportDetailLogic {
	return &QueryMemberReportDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberReportDetail 查询用户举报详情
func (l *QueryMemberReportDetailLogic) QueryMemberReportDetail(req *types.QueryMemberReportDetailReq) (resp *types.QueryMemberReportDetailResp, err error) {

	detail, err := l.svcCtx.MemberReportService.QueryMemberReportDetail(l.ctx, &cmsclient.QueryMemberReportDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户举报详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberReportDetailData{
		Id:               detail.Id,               // 编号
		ReportType:       detail.ReportType,       // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName: detail.ReportMemberName, // 举报人
		ReportObject:     detail.ReportObject,     // 被举报对象
		ReportStatus:     detail.ReportStatus,     // 举报状态：0->未处理；1->已处理
		HandleStatus:     detail.HandleStatus,     // 处理结果：0->无效；1->有效；2->恶意
		Note:             detail.Note,             // 备注
		CreateTime:       detail.CreateTime,       // 创建时间
	}
	return &types.QueryMemberReportDetailResp{
		Code:    "000000",
		Message: "查询用户举报成功",
		Data:    data,
	}, nil
}
