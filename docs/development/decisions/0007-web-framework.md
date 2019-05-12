# WEBフレームワーク

[Back](../../)

Webアプリケーションを構築するためのフレームワークに特化して検討する。  
また、優先すべきアプリ形態としてはREST-APIとなる。

|名称|既存|概要|ライセンス|開発状況|シェア|継続・保守性|
|:---|:---:|:---|:---:|:---|:---|:---|
|[gin-gonic/gin](https://github.com/gin-gonic/gin)|||MIT|かなり頻繁|?|問題なし（Issue多し）|
|[astaxie/beego](https://github.com/astaxie/beego)|||Apache2.0|||||
|[kataras/iris](https://github.com/kataras/iris)|||BSD|||||

3つほど、よく使われていそうなものを選択して調べた。  
なお、既存は上記のいずれでもない場合がある。

## gin-gonic/gin

実装サンプル  
https://github.com/takashno/learning-golang-gin

簡単に実装が出来た。  
複雑なことができるのかはわからないが、
基本的なWEBアプリを作成するのに圧倒的に不足しているようなものはない。
OpenApiのCodeGeneratorもginは対応している模様。