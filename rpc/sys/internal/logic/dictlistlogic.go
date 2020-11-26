package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictListLogic) DictList(in *sys.DictListReq) (*sys.DictListResp, error) {
	all, err := l.svcCtx.DictModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	if err != nil {
		return nil, err
	}
	var list []*sys.DictListData
	for _, dict := range *all {
		fmt.Println(dict)
		list = append(list, &sys.DictListData{
			Id:             dict.Id,
			Value:          dict.Value,
			Label:          dict.Label,
			Type:           dict.Type,
			Description:    dict.Description,
			Sort:           int64(dict.Sort),
			Remarks:        dict.Remarks,
			CreateBy:       dict.CreateBy,
			CreateTime:     dict.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dict.LastUpdateBy,
			LastUpdateTime: dict.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dict.DelFlag,
		})
	}

	return &sys.DictListResp{
		Total: 10,
		List:  list,
	}, nil

}
