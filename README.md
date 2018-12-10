# 概要
go + realize + grpc-gateway をDocker上で動かしホットリロードを行うためのサンプルです。

# 確認方法

## 1. ファイルの設置
ここのファイル一式を以下に設置
```
$GOPATH/src/github.com/gosagawa/realize_sample
```

## 2. hostsの設定

local.realize_sample.netをhostsファイル(MacOSなら/etc/hosts)に登録する
```
127.0.0.1 local.realize_sample.net
```

## 3. ライブラリ取得

```
make bundle
```

## 4. コンテナ起動
```
docker-compose up
```

##  5. 動作確認
まず、以下が表示できる事を確認。
```
$curl http://local.realize_sample.net/user/1
{"id":1,"name":"John","age":18}
```

その後adapter/grpc/user.goを書き換えて保存後、APIの値が変われば成功です（自動でビルドされて、反映されます）。
```
curl http://local.realize_sample.net/user/1
{"id":1,"name":"John","age":18}
```

adapter/grpc/protoDifinition以下にproto定義を持っています。
protoは保存しても即時反映されず、protoのmakeした後に反映されます。
```
make proto
```

# 解説記事

このサンプルをGopher道場advent calendar 2018 12/13公開の記事で解説します。公開後にリンクします。
