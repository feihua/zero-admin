package seckillsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySeckillSessionListLogic 查询秒杀场次列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:29:58
*/
type QuerySeckillSessionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySeckillSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillSessionListLogic {
	return &QuerySeckillSessionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySeckillSessionList 查询秒杀场次列表
func (l *QuerySeckillSessionListLogic) QuerySeckillSessionList(in *smsclient.QuerySeckillSessionListReq) (*smsclient.QuerySeckillSessionListResp, error) {
	seckillSession := query.SmsSeckillSession
	q := seckillSession.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(seckillSession.Name.Like("%" + in.Name + "%"))
	}

	if in.Status != 2 {
		q = q.Where(seckillSession.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀场次列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询秒杀场次列表失败")
	}

	var list []*smsclient.SeckillSessionListData

	for _, item := range result {
		list = append(list, &smsclient.SeckillSessionListData{
			Id:         item.ID,                                          // 秒杀场次ID
			Name:       item.Name,                                        // 场次名称
			StartTime:  item.StartTime,                                   // 开始时间
			EndTime:    item.EndTime,                                     // 结束时间
			Status:     item.Status,                                      // 状态：0-禁用，1-启用
			Sort:       item.Sort,                                        // 排序
			CreateBy:   item.CreateBy,                                    // 创建人ID
			CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
			UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
			UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间

		})
	}

	return &smsclient.QuerySeckillSessionListResp{
		Total: count,
		List:  list,
	}, nil
}
