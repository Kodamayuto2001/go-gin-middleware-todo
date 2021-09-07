### API仕様
```
POST:   /api/v1/users/add
リクエストヘッダ｛"name":"hoge","email":"hoge@hoge.com","password":"password"｝
```
```
$ curl -XPOST http://localhost:3000/api/v1/users/add \
-H "Content-Type: application/json" \
-d '{"name":"hoge","email":"hoge@hoge.com","password":"password"}'
```


### wsl:ubuntu-LTS mysql起動
```
$ sudo service mysql start
```

### mysqlログイン
```
$ sudo mysql -u USER -p
Enter password: ****************
```

### データベース一覧表示
```
mysql> show databases;
```

### データベースの選択
```
mysql> use DATABASE;
```

### テーブル一覧表示
```
mysql> show tables;
```

### レコード一覧表示
```
mysql> select * from TABLE;
```

### 会員登録用テーブルの作成例
```
mysql> create table `members` (
    `id`        int not null auto_increment primary key,
    `name`      varchar(100) not null,
    `email`     varchar(100) not null,
    `password`  varchar(100) not null
) engine=InnoDB default charset=utf8;
```

### テーブル設計の表示
```
mysql> desc TABLE_NAME;
```
