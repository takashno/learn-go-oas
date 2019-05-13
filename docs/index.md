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

