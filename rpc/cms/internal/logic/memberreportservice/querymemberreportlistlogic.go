package memberreportservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberReportListLogic 查询用户举报列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryMemberReportListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReportListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportListLogic {
	return &QueryMemberReportListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberReportList 查询用户举报列表
func (l *QueryMemberReportListLogic) QueryMemberReportList(in *cmsclient.QueryMemberReportListReq) (*cmsclient.QueryMemberReportListResp, error) {
	memberReport := query.CmsMemberReport
	q := memberReport.WithContext(l.ctx)
	if in.ReportType != 2 {
		q = q.Where(memberReport.ReportType.Eq(in.ReportType))
	}
	if len(in.ReportMemberName) > 0 {
		q = q.Where(memberReport.ReportMemberName.Like("%" + in.ReportMemberName + "%"))
	}
	if len(in.ReportObject) > 0 {
		q = q.Where(memberReport.ReportObject.Like("%" + in.ReportObject + "%"))
	}
	if in.ReportStatus != 2 {
		q = q.Where(memberReport.ReportStatus.Eq(in.ReportStatus))
	}
	if in.HandleStatus != 2 {
		q = q.Where(memberReport.HandleStatus.Eq(in.HandleStatus))
	}
	if len(in.Note) > 0 {
		q = q.Where(memberReport.Note.Like("%" + in.Note + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询用户举报列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户举报列表失败")
	}

	var list []*cmsclient.MemberReportListData

	for _, item := range result {
		list = append(list, &cmsclient.MemberReportListData{
			Id:               item.ID,                              // 编号
			ReportType:       item.ReportType,                      // 举报类型：0->商品评价；1->话题内容；2->用户评论
			ReportMemberName: item.ReportMemberName,                // 举报人
			ReportObject:     item.ReportObject,                    // 被举报对象
			ReportStatus:     item.ReportStatus,                    // 举报状态：0->未处理；1->已处理
			HandleStatus:     item.HandleStatus,                    // 处理结果：0->无效；1->有效；2->恶意
			Note:             item.Note,                            // 备注
			CreateTime:       time_util.TimeToStr(item.CreateTime), // 创建时间

		})
	}

	logc.Infof(l.ctx, "查询用户举报列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.QueryMemberReportListResp{
		Total: count,
		List:  list,
	}, nil
}
