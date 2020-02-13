# 用户身份验证

第三方应用请求腾讯文档开放平台用户身份验证时，按如下方式构造链接并跳转至此链接。
腾讯文档用户通过腾讯文档客户端、微信客户端、QQ客户端完成扫码登录，或者通过手机验证码、邮箱验证码登录。

## 请求登录验证

请求方式：GET

请求地址：https://docs.qq.com/oauth/v1/code/?appid={APP_ID}&state={STATE}&redirect_uri={REDIRECT_URI}

参数说明：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| state | string | 否 | 回调URL时会附加此参数，第三方应用可以用来关联请求和回调
| redirect_uri | string | 是 | 验证完成后回调第三方应用的URL，需要URL Encode

## 获取登录用户票据

请求方式：POST

请求地址：https://docs.qq.com/oauth/v1/token/?appid={APP_ID}&appkey={APP_KEY}&grant_type={GRANT_TYPE}&code={CODE}

## 获取登录用户信息

请求方式：GET

请求地址 ：https://docs.qq.com/oauth/v1/profile/?appid={APP_ID}&openid={OPEN_ID}&access_token={ACCESS_TOKEN}
