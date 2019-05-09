# learning-go-oas

## install goenv and go lang

goenvというのがあったので、入れてみる。  
あとでWindowsでできるかちょっと不安。
https://github.com/syndbg/goenv/blob/master/INSTALL.md

```bash
$ cd ~
$ git clone https://github.com/syndbg/goenv.git ~/.goenv
$ echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bash_profile
$ echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bash_profile
$ echo 'eval "$(goenv init -)"' >> ~/.bash_profile
$ echo 'export PATH="$GOROOT/bin:$PATH"' >> ~/.bash_profile
$ echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.bash_profile
$ goenv -v
goenv 1.23.3
$ goenv install -l
Available versions:
  1.2.2
  // omitted
  1.12.2
  1.12.3
  1.12.4
$ goenv install 1.12.4
Downloading go1.12.4.darwin-amd64.tar.gz...
-> https://dl.google.com/go/go1.12.4.darwin-amd64.tar.gz
Installing Go Darwin 64bit 1.12.4...
Installed Go Darwin 64bit 1.12.4 to /Users/takashimanozomu/.goenv/versions/1.12.4
$ go global 1.12.4
$ go version
go version go1.12.4 darwin/amd64
```

`1.12.4` を入れてみる。

## Editor

VSCodeが便利そうなので、プラグインを入れてみる。  
`Go` というプラグインを入れてみる。  
かなり利用されている模様。

とりあえず、何も考えず...  
コマンドパレットにて、`go:install update tools` で全部チェックつけてインストールする。


## Hello World (gs-normal-01)

まずはHello World

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("hello world")
}
```

`F5` で実行

```bash
API server listening at: 127.0.0.1:38540
hello world
```

起動は簡単。  
ブレークポイントを貼ることでデバッグもできた。


### Lint

VSCodeの`Go` プラグインで一通りツール類を入れれば、Lint系も一緒に付いてくるようで以下のようなコードを書くと警告が出てくれる。

```go
var test string = "aaa"
```

これは、リテラル的に、文字列だからいちいち型を定義する必要ないです。ということらしい。  
Javaに慣れているせいか、省略の美学はあまりわからない。  
シンタックスシュガーはプロジェクトでの決め事をしっかりしておかないと可読性やメンテナンス性を下げるから注意が必要だと `Scala` を使った時に思い知った過去を持つ。


## Package管理

npmのような依存するパッケージ管理の仕組みがGOにも存在する。  
いろいろ歴史があるようだが、2019/05現在は [Module](https://github.com/golang/go/wiki/Modules) を使うべきみたい。

`Module` を利用するためには、初期化処理を行う必要がある。  
以下コマンドで初期化する。
```bash
go mod init
```

go.modというファイルが作成される。  
ここに依存するパッケージ？ライブラリ？の定義が追加される模様。  


## REST-API (gs-restapi-01)


