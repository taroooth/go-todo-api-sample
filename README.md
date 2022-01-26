# GO-TODO-API-SAMPLE

# 概要

- Golang勉強用のサンプルプロジェクト
- ToDoアプリのAPI

# 環境構築

## .envファイルを作成

```
$ cp .env.sample .env
```

## DB作成

MySQLを使用

```
$ mysql -u root -p
Enter password:

mysql>
```

`todo_sample`というDBを作成

```
mysql> create database todo_sample;
```

DB切り替え

```
mysql> use todo_sample;
Database changed
```

`todo`テーブル作成
```
mysql> source mysql/migrations/1_create_table_todo.up.sql;
Query OK, 0 rows affected (0.01 sec)
```

カラム確認
```
mysql> show columns from todo;
+---------+---------------------+------+-----+---------+----------------+
| Field   | Type                | Null | Key | Default | Extra          |
+---------+---------------------+------+-----+---------+----------------+
| id      | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| title   | varchar(255)        | YES  |     | NULL    |                |
| content | varchar(255)        | YES  |     | NULL    |                |
+---------+---------------------+------+-----+---------+----------------+
3 rows in set (0.00 sec)
```

## GoのPackageをインストール
```
$ go get .
```

# 動作確認

## プログラム実行
```
$ go run main.go
```

## curl
```
GET
$ curl http://localhost:3000/todo/v1/list
{"data":[{"Id":1,"Title":"プログラム","Content":"Golang env"}],"message":"ok"}%

POST
$ curl -X POST -H "Content-Type: application/json" -d '{"title": "プログラミング", "content": "Golang"}' http://localhost:3000/todo/v1/add
{"status":"created"}%

PUT
$ curl -X PUT -H "Content-Type: application/json" -d '{"title": "英語", "content": "Reading"}' http://localhost:3000/todo/v1/update/1
{"status":"ok"}%

DELETE
$ curl -X DELETE http://localhost:3000/todo/v1/delete/1
{"status":"ok"}%
```
