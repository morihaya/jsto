# jsto

[![Test](https://github.com/morihaya/jsto/actions/workflows/test.yml/badge.svg)](https://github.com/morihaya/jsto/actions/workflows/test.yml)

日本標準時(JST)から各国の時間へ変換して表示します

[English](README.md) | [日本語](README.ja.md)

```bash
Usage:
  jsto [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cst         show CST time (UTC+8, JST-1)
  edt         show EDT time (UTC-4, JST-13)
  gmt         show GMT/BST time (UTC+0/+1, JST-9/-8)
  help        Help about any command
  ist         show IST time (UTC+5:30, JST-3:30)
  pdt         show PDT time (UTC-7, JST-15)
  utc         show UTC time (UTC+0, JST-9)

Flags:
  -h, --help     help for jsto
  -t, --toggle   Help message for toggle

Use "jsto [command] --help" for more information about a command.
```

出力例
```
'UTC' The time is:
 2022/04/22 13:15:02
```

JSTの時間を指定して変換することも可能です:

```bash
jsto utc 12:30
```

出力例
```
'UTC' time for JST 12:30 is:
 2025/11/21 03:30:00
```

## テスト

単体テストを実行するには、以下のコマンドを実行してください:

```bash
go test -v ./...
```

## ビルド

### ローカルビルド

ローカルでアプリケーションをビルドするには:

```bash
go build -o jsto .
```

### Dockerビルド

Dockerイメージをビルドするには:

```bash
docker build -t jsto .
```
