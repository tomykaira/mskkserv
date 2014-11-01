# mskkserv

モダンな SKK サーバ実装です。

## 特徴

- 単一のプロセスが全リクエストを処理するため高速である
- 複数のエンジンを登録し、候補がみつかるまで順に検索する機能を持つ
    - SKK 界隈で標準的に使われている CDB 形式のローカル辞書
    - [Google Input Tools](http://www.google.com/inputtools/) を用いた変換
- 特定のファイル配置に依存せず、すべてが設定可能である
- よくテストされている

## License

Copyright 2014 tomykaira. All Rights Reserved.

This software is licensed under the MIT license.
See `LICENSE` for details.
