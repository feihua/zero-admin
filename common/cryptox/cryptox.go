package cryptox

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Crypto(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	return string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
}

func CheckCrypto(hashedPassword, password string) bool {
	// 密码验证
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) //验证（对比）
	if err != nil {
		return false
	} else {
		return true
	}
}
