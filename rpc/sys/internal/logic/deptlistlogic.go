package logic

import (
	"context"
	"encoding/json"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sys.DeptListReq) (*sys.DeptListResp, error) {
	count, _ := l.svcCtx.DeptModel.Count()
	all, err := l.svcCtx.DeptModel.FindAll(1, count)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询机构列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sys.DeptListData
	for _, dept := range *all {

		list = append(list, &sys.DeptListData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dept.LastUpdateBy,
			LastUpdateTime: dept.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dept.DelFlag,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询机构列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.DeptListResp{
		Total: count,
		List:  list,
	}, nil

}
