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

// QueryMemberReportListLogic 查询用户举报列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryMemberReportListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberReportListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportListLogic {
	return &QueryMemberReportListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberReportList 查询用户举报列表
func (l *QueryMemberReportListLogic) QueryMemberReportList(req *types.QueryMemberReportListReq) (resp *types.QueryMemberReportListResp, err error) {
	result, err := l.svcCtx.MemberReportService.QueryMemberReportList(l.ctx, &cmsclient.QueryMemberReportListReq{
		PageNum:          int32(req.Current),
		PageSize:         int32(req.PageSize),
		ReportType:       req.ReportType,       // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName: req.ReportMemberName, // 举报人
		ReportObject:     req.ReportObject,     // 被举报对象
		ReportStatus:     req.ReportStatus,     // 举报状态：0->未处理；1->已处理
		HandleStatus:     req.HandleStatus,     // 处理结果：0->无效；1->有效；2->恶意

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字用户举报列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberReportListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberReportListData{
			Id:               item.Id,               // 编号
			ReportType:       item.ReportType,       // 举报类型：0->商品评价；1->话题内容；2->用户评论
			ReportMemberName: item.ReportMemberName, // 举报人
			ReportObject:     item.ReportObject,     // 被举报对象
			ReportStatus:     item.ReportStatus,     // 举报状态：0->未处理；1->已处理
			HandleStatus:     item.HandleStatus,     // 处理结果：0->无效；1->有效；2->恶意
			Note:             item.Note,             // 备注
			CreateTime:       item.CreateTime,       // 创建时间
		})
	}

	return &types.QueryMemberReportListResp{
		Code:     "000000",
		Message:  "查询用户举报列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
