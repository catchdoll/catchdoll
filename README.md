# catchdoll
抓娃娃

#### API
###### url路由
|路径|方式|返回值|功能|备注|
|---|---|---|----|----|
|/productOn|post||商品上架|
|/productOff|post||商品下架|
|/product/{id}|get||商品详情|
|/product/index|get||商品列表|
|/order|post||订单生成|
|/transaction|post||支付单生成|
||||支付操作|由多个子接口组成|
|/video/upload|post||上传视频|
|/video/{id}|get||获取单个视频信息|
|/vidoe/index|get||获取所有视频信息|
|/video/comment|post||评论视频|
|/video/comment|get||获取评论列表|
|/video/comment/{id}|get||获取评论详情||
|/machine/:id/|get||获取单个娃娃机信息(含评论)|已完成|
|/machine|get||获取娃娃机列表||
|/machine/map|get||获取娃娃机地图||
|/machine/comment|post||评论娃娃机||
|/machine/comment/index|get||获取娃娃机评论列表||
|/machine/comment/{id}|get||获取单个娃娃机评论||

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



