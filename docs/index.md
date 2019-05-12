# GoLang 開発

## 概要
GoLangによるWEBシステム開発を行う場合における、
下流工程をメインとした開発手法およびプロセスを検討・整理する。
下流工程とは、製造工程（MK）以降を指す。

## 検討要素一覧
検討要素に対して関連する工程（一部作業も含む）をチェックした表となる。

|要素|仕様|MK|UT|SI|CI|CD (Dev)|CD (Pro)|
|:---|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|[IDE](./development/decisions/0002-ide.md)||x||||||
|[GoLangバージョン管理](./development/decisions/0003-golang-version-control.md)||x|||||
|[パッケージ管理（利用）](./development/decisions/0004-package-management-use.md)||x|x|x||||
|[パッケージ管理（内部管理）](./development/decisions/0005-package-inner-management.md)||x|x|x||||
|[ビルドシステム](./development/decisions/0006-build.md)||x|x|x||x|x|
|[WEBフレームワーク](./development/decisions/0007-web-framework.md)||x||||||
|[ORMフレームワーク](./development/decisions/0008-orm-framework.md)||x||||||
|[静的解析](./development/decisions/0009-static-code-analysis.md)|||x||x|||
|[テストフレームワーク](./development/decisions/0010-test-framework.md)|||x||x|||
|[Mockフレームワーク](./development/decisions/0011-mock-framework.md)|||x||x|||
|[カバレッジ](./development/decisions/0012-coverage.md)|||x||x|||
|[自動生成](./development/decisions/0013-auto-generator.md)|x|x|x||||
|[OAS3（ソースコード）](#OAS3（ソースコード）)||x||||||
|[OAS3（スタブ）](#OAS3（スタブ）)|||x|x||||
|[OAS3（ドキュメント）](#OAS3（ドキュメント）)|x|||||||
|[CIシステム](#CIシステム)||x|x||x|||

それぞれの内容については、[adr-tools](https://github.com/npryce/adr-tools/blob/master/INSTALL.md)で作成したMarkdownにてまとめる。

---

## 参考資料

|リンク|概要|
|:---|:---|
|[Go言語 文法まとめ](https://qiita.com/rock619/items/db44507d02814e490902)|文法がまとまっている|
|[実践Go言語](http://golang.jp/effective_go)|構文＋言語仕様説明|
|[コーディングから攻めるGO性能の話](https://persol-pt.github.io/posts/tech-workshop20180406/)|GOのパフォーマンスチューニング|

---

## パッケージ管理（利用）

GoLangに含有されている。  
ここ数年で大分揺れている。  
以下遷移も含めて一覧となる。  

|方法|Since|概要|
|:---|:---|:---|
|go get|?|?|
|vgo|?|?|
|Module|1.11|?|


---

## パッケージ管理（内部管理）

開発用の内部パッケージの管理方法。  
`Module` については `git` で `clone`するのが通例のようなので、
そういう意味だとGitLabがあればいいのか？  
あとは、`go mod` コマンドのリポジトリの向け先を変えることができればOK？

---

## ビルドシステム

GoLangに含有されている。  
Java（jar|war|ear|fatjar）やNode.js（bundle.js）のようなアーカイブは不要っぽい。  
パッケージ参照を行って、自前ビルドするだけ？

そのため、`Gradle`、`Maven`、といったビルドシステム的な位置付けのものは存在しないと思われる。  
要継続調査

---

## WEBフレームワーク

Webアプリケーションを構築するためのフレームワークに特化して検討する。  
また、優先すべきアプリ形態としてはREST-APIとなる。

|名称|既存|概要|ライセンス|開発状況|シェア|継続・保守性|
|:---|:---:|:---|:---:|:---|:---|:---|
|[gin-gonic/gin](https://github.com/gin-gonic/gin)|||MIT|||||
|[astaxie/beego](https://github.com/astaxie/beego)|||Apache2.0|||||
|[kataras/iris](https://github.com/kataras/iris)|||BSD|||||

3つほど、よく使われていそうなものを選択して調べた。  
なお、既存は上記のいずれでもない場合がある。

---

## 静的解析

静的解析は外部ツールを用いて実施する。  
コーディング時に、`Lint` 系ツールである `golint` を利用して検知する方法が一番効率が良い。  
コーディング中に、リアルタイムに警告を出してくれるため是正対応を行いやすい。

---

## Mockフレームワーク

GoLangのUnitテストからの延長線上になるが、  
Mockフレームワークについても調べる。


---

## カバレッジ

GoLangの `Testing` に `Coverage` を取得する仕組みが組み込まれている。  

### Testing - Coverage

#### Pros
- テストの実行オプションに追加するだけで簡単に取得が可能
- HTMLのレポートへ変換が可能（見た目的にはそこまでクォリティは高くない）
#### Cons
- `C0` のみ対応？レポートを見る限りそう見える…

#### TODO
- SonarQube等との連携方法

---

## 自動生成

GoLangによるWEBアプリケーション開発に導入可能な自動生成の仕組みについて調査する。  
ただし、フレームワークの選定に引きずられる部分が大きいところでもある。

---