package logic

import (
	"go-zero-admin/rpc/us/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

//返回	Class:
//		Academy:
//		School:
//		Grade:
func GetRoleExtendInfoByRoleName(svcCtx *svc.ServiceContext, roleName string, userId int64) (map[string]string, error) {
	retMap := make(map[string]string, 0)
	switch roleName {
	case "teacher":
		if usTeacher, err := svcCtx.UsTeacherModel.FindOne(userId); err != nil {
			return retMap, err
		} else {
			retMap["academy"] = usTeacher.Academy.String
			retMap["School"] = usTeacher.School.String
			return retMap, nil
		}
	case "student":
		if usStudent, err := svcCtx.UsStudentModel.FindOne(userId); err != nil {
			return retMap, err
		} else {
			retMap["class"] = usStudent.Class.String
			retMap["grade"] = usStudent.Grade.String
			retMap["academy"] = usStudent.Academy.String
			retMap["School"] = usStudent.School.String
			return retMap, nil
		}
	default:
		logx.Errorf("role type error,RoleName:" + roleName)
		return retMap, errorRoleUnRegister
	}

	return retMap, nil
}
