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

// QueryMemberReportDetailLogic 查询用户举报详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QueryMemberReportDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberReportDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberReportDetailLogic {
	return &QueryMemberReportDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberReportDetail 查询用户举报详情
func (l *QueryMemberReportDetailLogic) QueryMemberReportDetail(in *cmsclient.QueryMemberReportDetailReq) (*cmsclient.QueryMemberReportDetailResp, error) {
	item, err := query.CmsMemberReport.WithContext(l.ctx).Where(query.CmsMemberReport.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询用户举报详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户举报详情失败")
	}

	data := &cmsclient.QueryMemberReportDetailResp{
		Id:               item.ID,                              // 编号
		ReportType:       item.ReportType,                      // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName: item.ReportMemberName,                // 举报人
		ReportObject:     item.ReportObject,                    // 被举报对象
		ReportStatus:     item.ReportStatus,                    // 举报状态：0->未处理；1->已处理
		HandleStatus:     item.HandleStatus,                    // 处理结果：0->无效；1->有效；2->恶意
		Note:             item.Note,                            // 备注
		CreateTime:       time_util.TimeToStr(item.CreateTime), // 创建时间
	}

	logc.Infof(l.ctx, "查询用户举报详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
