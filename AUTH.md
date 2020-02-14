# 用户身份验证

第三方应用请求腾讯文档开放平台用户身份验证时，按如下方式构造链接并跳转至此链接。
腾讯文档用户通过腾讯文档客户端、微信客户端、QQ客户端完成扫码登录，或者通过手机验证码、邮箱验证码登录。

## 请求登录验证

请求方式：GET

请求地址：https://docs.qq.com/oauth/v1/auth?appid={APP_ID}&state={STATE}&redirect_uri={REDIRECT_URI}&response_type={RESPONSE_TYPE}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| state | string | 否 | 回调URL时会附加此参数，第三方应用可以用来关联请求和回调
| redirect_uri | string | 是 | 验证完成后回调第三方应用的URL，需要URL Encode
| response_type | string | 是 | 第三方应用请求腾讯文档开放平台验证的临时票据类型，当前只支持填code

回调参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| code | string | 是 | 临时票据，第三方应用可以用来请求访问票据，时效5分钟，只能使用1次
| state | string | 否 | 回调URL时会附加此参数，第三方应用可以用来关联请求和回调

## 获取登录用户票据

请求方式：POST

请求地址：https://docs.qq.com/oauth/v1/token?appid={APP_ID}&appkey={APP_KEY}&grant_type={GRANT_TYPE}&code={CODE}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| appkey | string | 否 | 第三方应用在腾讯文档开放平台申请的应用密钥
| grant_type | string | 是 | 临时票据类型，目前只支持authorization_code
| code | string | 是 | 临时票据

返回参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| ret | int | 是 | 返回码，0表示成功，其他表示失败
| msg | string | 是 | 返回信息，成功为空串
| openid | string | 是 | 用户唯一标识
| access_token | string | 是 | 访问票据，用于获取用户资源
| access_expires_in | string | 是 | 访问票据的过期时间戳，默认为1月
| refresh_token | string | 是 | 刷新票据，用于刷新访问票据
| refresh_expires_in | string | 是 | 刷新票据的过期时间戳，默认为1年

## 刷新登录用户票据

请求方式：POST

请求地址：https://docs.qq.com/oauth/v1/refresh?appid={APP_ID}&openid={OPEN_ID}&refresh_token={REFRESH_TOKEN}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| openid | string | 是 | 用户唯一标识
| refresh_token | string | 是 | 刷新票据，用于刷新访问票据

返回参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| ret | int | 是 | 返回码，0表示成功，其他表示失败
| msg | string | 是 | 返回信息，成功为空串
| access_token | string | 是 | 访问票据，用于获取用户资源
| access_expires_in | string | 是 | 访问票据的过期时间戳，默认为1月

## 获取登录用户信息

请求方式：GET

请求地址 ：https://docs.qq.com/oauth/v1/userinfo?appid={APP_ID}&openid={OPEN_ID}&access_token={ACCESS_TOKEN}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| openid | string | 是 | 用户唯一标识
| access_token | string | 是 | 访问票据，用于获取用户资源

返回参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| ret | int | 是 | 返回码，0表示成功，其他表示失败
| msg | string | 是 | 返回信息，成功为空串
| nick_name | string | 是 | 用户昵称
| avatar_url | string | 是 | 用户头像URL
