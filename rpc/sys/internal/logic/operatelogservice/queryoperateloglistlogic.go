package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
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
	operateLog := query.SysOperateLog
	q := operateLog.WithContext(l.ctx)
	if len(in.Title) > 0 {
		q = q.Where(operateLog.Title.Like("%" + in.Title + "%"))
	}
	if in.BusinessType != 2 {
		q = q.Where(operateLog.BusinessType.Eq(in.BusinessType))
	}
	if len(in.Method) > 0 {
		q = q.Where(operateLog.Method.Like("%" + in.Method + "%"))
	}
	if len(in.RequestMethod) > 0 {
		q = q.Where(operateLog.RequestMethod.Like("%" + in.RequestMethod + "%"))
	}
	if in.OperatorType != 2 {
		q = q.Where(operateLog.OperatorType.Eq(in.OperatorType))
	}
	if len(in.OperateName) > 0 {
		q = q.Where(operateLog.OperateName.Like("%" + in.OperateName + "%"))
	}
	if len(in.DeptName) > 0 {
		q = q.Where(operateLog.DeptName.Like("%" + in.DeptName + "%"))
	}
	if len(in.OperateUrl) > 0 {
		q = q.Where(operateLog.OperateURL.Like("%" + in.OperateUrl + "%"))
	}
	if len(in.OperateIp) > 0 {
		q = q.Where(operateLog.OperateIP.Like("%" + in.OperateIp + "%"))
	}
	if len(in.OperateLocation) > 0 {
		q = q.Where(operateLog.OperateLocation.Like("%" + in.OperateLocation + "%"))
	}

	if len(in.Platform) > 0 {
		q = q.Where(operateLog.Platform.Like("%" + in.Platform + "%"))
	}
	if len(in.Browser) > 0 {
		q = q.Where(operateLog.Browser.Like("%" + in.Browser + "%"))
	}
	if len(in.Version) > 0 {
		q = q.Where(operateLog.Version.Like("%" + in.Version + "%"))
	}
	if len(in.Os) > 0 {
		q = q.Where(operateLog.Os.Like("%" + in.Os + "%"))
	}
	if len(in.Arch) > 0 {
		q = q.Where(operateLog.Arch.Like("%" + in.Arch + "%"))
	}
	if len(in.Engine) > 0 {
		q = q.Where(operateLog.Engine.Like("%" + in.Engine + "%"))
	}
	if len(in.EngineDetails) > 0 {
		q = q.Where(operateLog.EngineDetails.Like("%" + in.EngineDetails + "%"))
	}
	if len(in.Extra) > 0 {
		q = q.Where(operateLog.Extra.Like("%" + in.Extra + "%"))
	}
	if in.Status != 2 {
		q = q.Where(operateLog.Status.Eq(in.Status))
	}

	result, count, err := q.Order(operateLog.ID.Desc()).FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询操作日志列表失败")
	}
	var list = make([]*sysclient.OperateLogListData, 0, len(result))

	for _, item := range result {
		list = append(list, &sysclient.OperateLogListData{
			Id:              item.ID,                               // 操作日志id
			Title:           item.Title,                            // 模块标题
			BusinessType:    item.BusinessType,                     // 业务类型（0其它 1新增 2修改 3删除）
			Method:          item.Method,                           // 方法名称
			RequestMethod:   item.RequestMethod,                    // 请求方式
			OperatorType:    item.OperatorType,                     // 操作类别（0其它 1后台用户 2手机端用户）
			OperateName:     item.OperateName,                      // 操作人员
			DeptName:        item.DeptName,                         // 部门名称
			OperateUrl:      item.OperateURL,                       // 请求URL
			OperateIp:       item.OperateIP,                        // 主机地址
			OperateLocation: item.OperateLocation,                  // 操作地点
			OperateParam:    item.OperateParam,                     // 请求参数
			JsonResult:      item.JSONResult,                       // 返回参数
			Platform:        item.Platform,                         // 平台信息
			Browser:         item.Browser,                          // 浏览器类型
			Version:         item.Version,                          // 浏览器版本
			Os:              item.Os,                               // 操作系统
			Arch:            item.Arch,                             // 体系结构信息
			Engine:          item.Engine,                           // 渲染引擎信息
			EngineDetails:   item.EngineDetails,                    // 渲染引擎详细信息
			Extra:           item.Extra,                            // 其他信息（可选）
			Status:          item.Status,                           // 操作状态(0:异常,正常)
			ErrorMsg:        item.ErrorMsg,                         // 错误消息
			OperateTime:     time_util.TimeToStr(item.OperateTime), // 操作时间
			CostTime:        item.CostTime,                         // 消耗时间
		})
	}

	return &sysclient.QueryOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
