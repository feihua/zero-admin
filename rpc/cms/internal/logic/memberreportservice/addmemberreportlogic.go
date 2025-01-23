package memberreportservicelogic

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

// AddMemberReportLogic 添加用户举报
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberReportLogic {
	return &AddMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberReport 添加用户举报
func (l *AddMemberReportLogic) AddMemberReport(in *cmsclient.AddMemberReportReq) (*cmsclient.AddMemberReportResp, error) {
	q := query.CmsMemberReport

	item := &model.CmsMemberReport{
		ReportType:       in.ReportType,       // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName: in.ReportMemberName, // 举报人
		ReportObject:     in.ReportObject,     // 被举报对象
		ReportStatus:     in.ReportStatus,     // 举报状态：0->未处理；1->已处理
		HandleStatus:     in.HandleStatus,     // 处理结果：0->无效；1->有效；2->恶意
		Note:             in.Note,             // 备注
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加用户举报失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加用户举报失败")
	}

	logc.Infof(l.ctx, "添加用户举报成功,参数：%+v", in)
	return &cmsclient.AddMemberReportResp{}, nil
}
