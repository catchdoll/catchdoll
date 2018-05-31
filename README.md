# catchdoll
抓娃娃

#### API

#### 接口验证
* 接口采用jwt认证
* /login接口提供登录，首次登录post传入
    ```javascripts
    {"username":"xxx","password":"xxx"}
    ```
    格式的json
* 登录成功的话返回一个
    ```javascript
    {"code": 200,
    "expire": "2018-05-25T23:44:40+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjcyNjMwODAsImlkIjoiMSIsIm9yaWdfaWF0IjoxNTI3MjU5NDgwfQ.E3-URRY0XhoOr1xL5IHsGOa8nZ1DC-19MyvdTnHWHgI"}
    ```
    格式的json，用token作为每次请求接口的Bearer Token值
###### url路由
|路径|方式|返回值|功能|备注|
|---|---|---|----|----|
|/productOn|post||商品上架|已完成|
|/productOff|post||商品下架|已完成|
|/product/{id}|get||商品详情|
|/product/index|get||商品列表|
|/order|post||订单生成|
|/transaction|post||支付单生成|
||||支付操作|由多个子接口组成|
|/video/upload|post||上传视频|
|/video/{id}|get||获取单个视频信息(含评论)|已完成|
|/vidoe_top|get||获取置顶视频|已完成|
|/video/comment|post||评论视频|已完成|
|/video/comment|get||获取评论列表||
|/machine/:id/|get||获取单个娃娃机信息(含评论)|已完成|
|/machine|get||获取置顶娃娃机列表|已完成|
|/machine/map|get||获取娃娃机地图||
|/machine/comment|post||评论娃娃机|已完成|

###### 返回格式说明
|字段|含义|备注|
|---|---|---|
|status|请求数据状态|见下表|
|message|主要是错误信息|成功均为success|
|result|返回内容|错误请求无result字段|

###### status状态码说明
|状态码|含义|备注|
|---|---|---|
|600|请求正常|
|601|请求参数不合法|
|602|查询无结果|
|603|数据操作故障|



