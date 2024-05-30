package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogListLogic 查询操作日志列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:09
*/
type QueryOperateLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogListLogic {
	return &QueryOperateLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOperateLogList 查询操作日志列表
func (l *QueryOperateLogListLogic) QueryOperateLogList(in *sysclient.QueryOperateLogListReq) (*sysclient.QueryOperateLogListResp, error) {
	q := query.SysOperateLog.WithContext(l.ctx)
	if len(in.OperationIp) > 0 {
		q = q.Where(query.SysOperateLog.OperationIP.Like("%" + in.OperationIp + "%"))
	}
	if len(in.DeptName) > 0 {
		q = q.Where(query.SysOperateLog.DeptName.Like("%" + in.DeptName + "%"))
	}
	if len(in.OperationName) > 0 {
		q = q.Where(query.SysOperateLog.OperationName.Like("%" + in.OperationName + "%"))
	}
	if in.OperationStatus != 0 {
		q = q.Where(query.SysOperateLog.OperationStatus.Eq(in.OperationStatus))
	}
	if len(in.OperationType) > 0 {
		q = q.Where(query.SysOperateLog.OperationType.Like("%" + in.OperationType + "%"))
	}
	if len(in.OperationUrl) > 0 {
		q = q.Where(query.SysOperateLog.OperationURL.Like("%" + in.OperationUrl + "%"))
	}
	if len(in.Title) > 0 {
		q = q.Where(query.SysOperateLog.Title.Like("%" + in.Title + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询操作日志列表失败")
	}
	var list []*sysclient.OperateLogListData
	for _, log := range result {
		list = append(list, &sysclient.OperateLogListData{
			DeptName:          log.DeptName,
			Id:                log.ID,
			OperationIp:       log.OperationIP,
			OperationName:     log.OperationName,
			OperationParams:   log.OperationParams,
			OperationResponse: log.OperationResponse,
			OperationStatus:   log.OperationStatus,
			OperationTime:     log.OperationTime.Format("2006-01-02 15:04:05"),
			OperationType:     log.OperationType,
			OperationUrl:      log.OperationURL,
			RequestMethod:     log.RequestMethod,
			Title:             log.Title,
			UseTime:           log.UseTime,
		})
	}

	logc.Infof(l.ctx, "查询操作日志列表,参数：%+v,响应：%+v", in, list)
	return &sysclient.QueryOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
