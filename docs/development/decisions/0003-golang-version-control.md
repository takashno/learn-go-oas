# GoLangバージョン管理

Date: 2019-05-11

## Status

調査中

## Context

固定とするならば、msi等のインストローラでも構わないように思える。  
それなりに言語としてのアップデートも激しい模様。  

## Decision

`goenv` が検索ヒットとして多いので試してみる。
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

それなりにスムーズにインストールが可能。  
ハマりポイントもなかったが、Windowsで試せていない。

## Consequences

判断中

