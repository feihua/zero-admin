package noticeservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryNoticeListLogic 查询通知公告列表
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type QueryNoticeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNoticeListLogic {
	return &QueryNoticeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryNoticeList 查询通知公告列表
func (l *QueryNoticeListLogic) QueryNoticeList(in *sysclient.QueryNoticeListReq) (*sysclient.QueryNoticeListResp, error) {
	notice := query.SysNotice
	q := notice.WithContext(l.ctx)
	if len(in.NoticeTitle) > 0 {
		q = q.Where(notice.NoticeTitle.Like("%" + in.NoticeTitle + "%"))
	}
	if in.NoticeType != 0 {
		q = q.Where(notice.NoticeType.Eq(in.NoticeType))
	}

	if in.Status != 2 {
		q = q.Where(notice.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询通知公告列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询通知公告列表失败")
	}

	var list []*sysclient.QueryNoticeListData

	for _, item := range result {
		list = append(list, &sysclient.QueryNoticeListData{
			Id:            item.ID,                                 // 公告ID
			NoticeTitle:   item.NoticeTitle,                        // 公告标题
			NoticeType:    item.NoticeType,                         // 公告类型（1:通知,2:公告）
			NoticeContent: item.NoticeContent,                      // 公告内容
			Status:        item.Status,                             // 公告状态（0:关闭,1:正常 ）
			Remark:        item.Remark,                             // 备注
			CreateBy:      item.CreateBy,                           // 创建者
			CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:      item.UpdateBy,                           // 更新者
			UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间
		})
	}

	return &sysclient.QueryNoticeListResp{
		Total: count,
		List:  list,
	}, nil
}
