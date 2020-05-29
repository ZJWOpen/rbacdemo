package common

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"just.for.test/rbacdemo/util"
)

const (
	// TokenKey 生产token所需要的key,支持长度为16，24，...其他的长度为非法的
	TokenKey = "casbin.casbin.cm"
	// SysTokenKey = "sys.casbin.casbin.com.cn"
)

// Token 生产所需的结构体
type Token struct {
	ID       int64
	UserName string
	Role     string
}

func TokenToString(o *Token) string {
	return tokenToString(o, TokenKey)
}

func tokenToString(o *Token, tokenKey string) string {
	data, err := json.Marshal(o)
	if err != nil {
		logrus.Errorln("GenTokenString Error:", err)
		return ""
	}
	data, err = util.Encrypt(data, []byte(tokenKey))
	if err != nil {
		logrus.Errorln("GenTokenString Encrypt Error:", err)
		return ""
	}
	return EncodeBase64(data)
}

// func SysTokenToString(o *Token) string {
// 	return tokenToString(o, SysTokenKey)
// }

func TokenFromString(content string) (*Token, error) {
	return tokenFromStringCore(content, TokenKey)
}

func tokenFromStringCore(content, tokenKey string) (*Token, error) {
	data, err := DecodeBase64(content)
	if err != nil {
		logrus.Errorln("Token ParseToken Base64Decode Error:", err)
		return nil, err
	}
	data, err = util.Decrypt(data, []byte(tokenKey))
	if err != nil {
		logrus.Errorln("Token ParseToken Decrypt Error:", err)
		return nil, err
	}
	token := new(Token)
	if err = json.Unmarshal(data, &token); err != nil {
		logrus.Errorln("Token ParseToken Unmarshal Error:", err)
		return nil, err
	}
	return token, nil

}

// func SysTokenFromString(content string) (*Token, error) {
// 	return tokenFromStringCore(content, SysTokenKey)
// }
