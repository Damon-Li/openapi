package auth

import "github.com/google/uuid"

func Authorization(code string) (openid string, err error) {
	// TODO: 根据code调用后台的账号服务进行登录验证，获取腾讯文档用户toid，根据toid生成openid，并建立mapping关系
	return uuid.New().String()[:8], nil
}