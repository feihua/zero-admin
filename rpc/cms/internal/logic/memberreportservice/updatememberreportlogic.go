package memberreportservicelogic

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

// UpdateMemberReportLogic 更新用户举报
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReportLogic {
	return &UpdateMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberReport 更新用户举报
func (l *UpdateMemberReportLogic) UpdateMemberReport(in *cmsclient.UpdateMemberReportReq) (*cmsclient.UpdateMemberReportResp, error) {
	q := query.CmsMemberReport.WithContext(l.ctx)

	// 1.根据用户举报id查询用户举报是否已存在
	_, err := q.Where(query.CmsMemberReport.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "根据用户举报id：%d,查询用户举报失败,异常:%s", in.Id, err.Error())
		return nil, errors.New(fmt.Sprintf("查询用户举报失败"))
	}

	item := &model.CmsMemberReport{
		ID:               in.Id,               // 编号
		ReportType:       in.ReportType,       // 举报类型：0->商品评价；1->话题内容；2->用户评论
		ReportMemberName: in.ReportMemberName, // 举报人
		ReportObject:     in.ReportObject,     // 被举报对象
		ReportStatus:     in.ReportStatus,     // 举报状态：0->未处理；1->已处理
		HandleStatus:     in.HandleStatus,     // 处理结果：0->无效；1->有效；2->恶意
		Note:             in.Note,             // 备注
	}

	// 2.用户举报存在时,则直接更新用户举报
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户举报失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新用户举报失败")
	}

	logc.Infof(l.ctx, "更新用户举报成功,参数：%+v", in)
	return &cmsclient.UpdateMemberReportResp{}, nil
}
