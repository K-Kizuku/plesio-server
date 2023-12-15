# ADR (Architecture Decision Record) for ristretto

## Goのインメモリキャッシュの技術選定

### 背景
セッション情報を保持するインメモリキャッシュの選定を行う．  
今回はパフォーマンスを最優先事項として技術選定を行う．

### 決定
[ristretto](https://github.com/dgraph-io/ristretto)を採用する

### 理由
今回の最優先事項がパフォーマンスであることから，最もパフォーマンスが良いという観点から，ristrettoを採用した．   
githubのスター数は他のパッケージと比べると少ないものの，リリース頻度が高く，直近のリリースも比較的近いため，今後も継続して開発が行われることが想定される．  
また，watchが多いことから，他のユーザーからの関心も大きいと判断し，ユーザー数(スター数)も今後増加すると考えられる．  

### 代替案
- [go-cache](https://github.com/patrickmn/go-cache)
  - かなりメジャーなパッケージ．代替案の中では最もgithubのスター数が多い
  - 最新リリースが6年前であるため，最近はあまりメンテナンスがされていないと考えられるため，採用を見送り
- [bigcache](https://github.com/allegro/bigcache)
  - 利用しやすそうなinterfaceとなっているため，開発効率は良さそう
  - item全体でTTLを設定できる
  - githubのスター数も多く安定して使うことができそうであるため，第２候補とする
- [golang-lru](https://github.com/hashicorp/golang-lru)
  - シンプルな機能となっており，扱いやすそう
  - シンプルすぎるが故に，カスタマイズが必要となった場合に苦労しそうなので採用を見送り

### 影響
うまくキャッシュできているかどうか，テストする必要あり   
キャッシュに使うメモリについても要検討

### 参考情報

https://github.com/dgraph-io/ristretto

https://github.com/patrickmn/go-cache

https://github.com/allegro/bigcache

https://github.com/hashicorp/golang-lru

執筆日：2023/12/11 23:16
