# mskkserv

モダンな SKK サーバ実装です。

## 特徴

- 単一のプロセスが全リクエストを処理するため高速である
- 複数のエンジンを登録し、候補がみつかるまで順に検索する機能を持つ
    - SKK 界隈で標準的に使われている CDB 形式のローカル辞書
    - [Google Input Tools](http://www.google.com/inputtools/) を用いた変換
- 特定のファイル配置に依存せず、すべてが設定可能である
- よくテストされている

## インストール

Go 言語で記述しているため、実行環境が必要。

```
go get github.com/tomykaira/mskkserv
```

ローカル辞書を利用する場合、CDB 形式の辞書を入手するか、自分で変換しておく。

## 利用法

```
mskkserv [-host=HOST] [-port=PORT] ENGINES
```

- HOST: バインドするホスト名 (デフォルトは `127.0.0.1`)
- PORT: listen するポート番号 (デフォルトは `1178`)
- ENGINES: 必須。次のなかから任意の組み合わせをスペース区切りで指定する。
  先に指定されたものから順に検索され、最初に結果がみつかったら後は検索しない。
    - `database=/path/to/SKK-JISYO.cdb`
    - `googletrans`

## License

Copyright 2014 tomykaira. All Rights Reserved.

This software is licensed under the MIT license.
See `LICENSE` for details.
