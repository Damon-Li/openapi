# 文档管理

## 通用请求参数

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| appid | string | 是 | 第三方应用在腾讯文档开放平台申请的应用标识
| openid | string | 是 | 用户唯一标识

## 通用请求头部

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| Content-Type | string | 是 | application/json
| Authorization | string | 是 | Bearer <access_token>

## 通用返回参数

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| ret | int | 是 | 返回码，0表示成功，其他表示失败
| msg | string | 否 | 返回信息，成功为空串

## 创建文档

请求方式：POST

请求地址：https://docs.qq.com/oauth/v1/doc/?appid={APP_ID}&openid={OPEN_ID}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_name | string | 是 | 文档名称
| auto_rename | int | 否 | 是否自动根据第一行内容修改文档名称，默认为0，表示不自动修改，为1，则自动修改

返回参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_id | string | 是 | 文档ID
| doc_url | string | 否 | 文档URL

## 删除文档

请求方式：DELETE

请求地址：https://docs.qq.com/oauth/v1/doc/?appid={APP_ID}&openid={OPEN_ID}&doc_id={OPEN_ID}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_id | string | 是 | 文档ID

## 修改文档

请求方式：PUT

请求地址：https://docs.qq.com/oauth/v1/doc/?appid={APP_ID}&openid={OPEN_ID}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_id | string | 是 | 文档ID
| doc_name | string | 是 | 文档名称

## 查询文档

请求方式：GET

请求地址：https://docs.qq.com/oauth/v1/doc/?appid={APP_ID}&openid={OPEN_ID}&doc_id={OPEN_ID}

请求参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_id | string | 是 | 文档ID

返回参数：

| 参数 | 类型 | 必须 | 说明 |
| --- | --- | --- | --- |
| doc_id | string | 是 | 文档ID
| doc_url | string | 是 | 文档URL
| owner_open_id | string | 否 | 文档作者用户唯一标识
| owner_nick_name | string | 否 | 文档作者用户昵称
| owner_avatar_url | string | 否 | 文档作者头像URL
| doc_status | string | 否 | 0:状态正常 1:已删除 2:在回收站中 3:父文件夹在回收站中

