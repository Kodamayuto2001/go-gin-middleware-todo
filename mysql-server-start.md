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
