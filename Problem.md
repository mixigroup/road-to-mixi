# Road to MIXI 課題

# 課題1
## 前提
Road to MIXI Mini SNS 00ではミニマルSNSのバックエンドAPIを実装していただきます。

このSNSはユーザー間にフレンド関係だけが設定できるSNSで、ユーザは

* ID
* 名前
* フレンド関係

というデータだけを持つ。 テーブル構成を以下のように定義する。

```
CREATE TABLE ‘users‘ (
‘id‘ bigint(20) NOT NULL AUTO_INCREMENT,
‘user_id‘ int(11) NOT NULL,
‘name‘ varchar(64) DEFAULT ’’ NOT NULL,
PRIMARY KEY (‘id‘)
);

-- user1 user2がフレンド関係
CREATE TABLE ‘friend_link‘ (
‘id‘ bigint(20) NOT NULL AUTO_INCREMENT,
‘user1_id‘ int(11) NOT NULL,
‘user2_id‘ int(11) NOT NULL,
PRIMARY KEY (‘id‘)
);

-- user1はuser2をblockしている
CREATE TABLE ‘block_list‘ (
‘id‘ bigint(20) NOT NULL AUTO_INCREMENT,
‘user1_id‘ int(11) NOT NULL,
‘user2_id‘ int(11) NOT NULL,
PRIMARY KEY (‘id‘)
);
```

## 課題1.1

ユーザのフレンドのリスト(ID, 名前)を返す機能を実装せよ。

* エンドポイント
    * ‘/get_friend_list‘
* パラメータ
    * ID:ユーザのID

## 課題1.2

ユーザのフレンドのフレンド(2hop)のリスト(ID, 名前)を返す機能を実装せよ。

* エンドポイント
    * ‘/get_friend_of_friend_list‘
* パラメータ
    * ID:ユーザのID

## 課題1.3

以下の機能を全問で作成したエンドポイントに実装せよ。

* 前問の機能を満たしていること。
* フレンドのフレンドのリストに1hopのフレンドが含まれていた場合、それらを削除する機能。
* block関係であった場合は、フレンドのリストには含めない機能。

## 課題1.4

ページネーションを含めたユーザのフレンドのリストを返す機能を実装せよ。

エンドポイントとパラメータ以外に関しては、前問の機能を満たしていること。

* エンドポイント
    * ‘/get_friend_of_friend_list_paging‘
* パラメータ
    * ID:ユーザのID
    * limit: 1ページあたりの表示件数
    * page: ページの位置

## 課題1.5

このミニマムSNSに何か一つ新機能を実装せよ。

:bulb: テーブル定義の再設計やコードのリファクタリングなどをしても構わない。

# 課題2

課題2では ミニマルSNSのnginxの設定を行ってもらいます。


## 前提

とあるWebアプリケーションが

* L7ロードバランサー -> nginx -> Webアプリケーション

という構成で作成されている。

このWebアプリケーションにおいてnginxの設定・チューニングを行う。

このロードバランサーとnginx、Webアプリケーションは、172.31.0.0/16のプライベートネットワーク内にある。

nginxプロセスが稼働しているサーバ(インスタンス)のドキュメントルート配下は以下のようなファイルツリーとなっている。

```
├── file
│     ├── 404.html
│     ├── error.html
│     └── maintenance.html
├── img
│     ├── image1.png
│     └── image2.png
└── index.html
```

## 課題2.1

以下の挙動を行う設定ファイルを提出せよ。

* ‘/‘にアクセスすると ‘index.html‘を返すこと。
* 存在しないfileにアクセスした場合、‘file/404.html‘を返すこと。
* ‘img/image1.png‘、‘img/image2.png‘にアクセスした場合、返すこと。
* ‘/test‘にアクセスした場合、"this is a test"というデータを返すこと。

## 課題2.2

以下の挙動を行う設定ファイルを提出せよ。

* 前問の機能を満たしていること。
* リバースプロキシするwebアプリケーションとして、課題1で作成したアプリケーションを設定すること。
* メンテナンスモードというものを設定することができ、メンテナンスモード中は全リクエストに対して常に‘file/maintenance.html‘を返すこと。

## 課題2.2

以下の挙動を行う設定を行うこと。

* 前問の機能を満たしていること。
* メンテナンスモードでも特定のIPアドレス(任意)からのアクセスは通常通りにアクセスできること。
* メンテナンスモードでも‘/test‘は "this is test"を返すこと(無条件)。
* メンテナンスモードの切り替えは無停止でできるようにすること。
* Webアプリケーションからエラーが返った場合にerror.htmlを返すこと。

## 課題2.3

以下の挙動を行う設定を行うこと。

* 前問の機能を満たしていること。
* nginx側の裏側のappにおいて、リクエスト元(クライアント)のIPアドレスをWebアプリケーション側で識別できるようにすること。
* ‘image/image1.png‘と‘image/image2.png‘への有効期限24時間のブラウザキャッシュを行うよう設定すること。

