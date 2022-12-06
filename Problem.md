# Road to Mixi 課題

## 前提
Road to Mixi Mini SNS 00ではミニマルSNSのバックエンドAPIを実装していただきます。

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

