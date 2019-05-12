# テストフレームワーク

[Back](../../index.md)

GoLangにテストを行うための仕組みが用意されている。  
詳細に関してはGoLangの[公式ドキュメント](https://golang.org/pkg/testing/)にTestingの利用方法が記載されている。

### 特徴/制約
- `testing` パッケージを利用（import）する
- `func TestXxx` で始まること（懐かしいJUnitっぽい制約）
- ファイル名は `_test.go` とすること
- `assert` が敢えて狙った上で用意されていないらしい

### サンプル

- 対象コード

```go
package main

import (
	"fmt"
)

func message() string {
	return "Hello World !!!"
}

func main() {
	fmt.Println(message())
}
```

- テストコード
```go
package main

import (
	"os"
	"testing"
)

// mainのテスト
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// messageのテスト
func TestMessage(t *testing.T) {
	if "Hello World !!!" != message() {
		t.Fatal("fail test")
	}
}
```

実行は、`go test` で単純実行はOK

- 正常時の結果
```bash
go test ./gs-normal-01/
ok      github.com/takashno/learning-go-oas/gs-normal-01        0.898s [no tests to run]
```

- エラー時の結果
```bash
go test ./gs-normal-01/
--- FAIL: TestMessage (0.00s)
    helloworld_test.go:16: fail test
FAIL
FAIL    github.com/takashno/learning-go-oas/gs-normal-01        0.891s
```
