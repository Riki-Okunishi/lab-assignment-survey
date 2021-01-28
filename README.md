# lab-assignment-survey
研究室配属に関連するスクリプト

## 一覧
+ notfound.go
  +  未入力者のIDをファイルに出力する

## notfound.go

### 概要
配属希望調査サイトでの集計が終了した後に，未入力者がいないか確認するためのスクリプト．

サイトからダウンロードできる集計結果の`csv`ファイルのフォーマットが以下であることを前提としている．
```
ID,[教授1],[教授2],...,[教授n],進学
```
`ID`は学籍番号，`[教授k]`は各教授のポイント,`進学`は「大学院進学を希望する」が1．

この集計結果と以下のような学籍番号の一覧を比較して，未入力者のIDの一覧を出力する．
```
ID1
ID2
.
.
.
```

### 使い方
```bash
> go run notfound.go -result results.csv -input input_data.txt -output notfound.txt
```

### 設定
教授の数が変わったときは`NumOfProfessor`の数を変える，

```go:notfound.go

const (
	NumOfProfessor = 11
)

```
