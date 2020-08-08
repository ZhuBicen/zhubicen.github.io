---
title: "MariaDB 中的 JSON 数据类型"
date: 2020-08-08T17:43:20+08:00
draft: false
tags: []
---

Maria DB 从版本 10.2.7 开始支持 JSON 数据类型，所以可以创建 JSON 类型的列

```
CREATE TABLE t (j JSON);

DESC t;
+-------+----------+------+-----+---------+-------+
| Field | Type     | Null | Key | Default | Extra |
+-------+----------+------+-----+---------+-------+
| j     | longtext | YES  |     | NULL    |       |
+-------+----------+------+-----+---------+-------+
```
查看该 table，可以看到 j 列的类型为 `longtext` ，在 MariaDB 的实现中， JSON 类型为 longtext 类型的别名。
```
MariaDB [test]> describe t;
+-------+----------+------+-----+---------+-------+
| Field | Type     | Null | Key | Default | Extra |
+-------+----------+------+-----+---------+-------+
| j     | longtext | YES  |     | NULL    |       |
+-------+----------+------+-----+---------+-------+
1 row in set (0.001 sec)
```

从版本 10.4.3 开始，JSON 类型的列默认使用 ` JSON_VALID` 的限制，所一插入的数据必须是正确的 JSON 格式

```
MariaDB [test]> insert into t values ('invalide');
ERROR 4025 (23000): CONSTRAINT `t.j` failed for `test`.`t`
```

在较低的版本上我们也可以使用 `JSON_VALID` 进行限制

## 创建含有  JSON  类型的表格

```SQL
CREATE TABLE products(id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 price DECIMAL(9,2) NOT NULL,
 stock INTEGER NOT NULL,
 attribute JSON,
 CHECK (JSON_VALID(attribute)));
```

## 插入 JSON 数据

```sql
INSERT INTO products VALUES(NULL, 'Blouse', 17, 15, '{"colour": "white"}');
```

如果插入错误格式（比如少了一个引号）的数据将会引起错误：

```
MariaDB [test]> INSERT INTO products VALUES(NULL, 'Blouse', 17, 15, '{"colour": "white}');
ERROR 4025 (23000): CONSTRAINT `products.attribute` failed for `test`.`products`
```
插入 JSON 数据后，查询结果如下：
```
MariaDB [test]> select * from products;
+----+--------+-------+-------+---------------------+
| id | name   | price | stock | attribute           |
+----+--------+-------+-------+---------------------+
|  1 | Blouse | 17.00 |    15 | {"colour": "white"} |
+----+--------+-------+-------+---------------------+
1 row in set (0.000 sec)
```

## 为  JSON 列建立索引

首先让我们再插入一条数据
```
MariaDB [test]> INSERT INTO products VALUES(NULL, 'Pingpong', 18, 15, '{"colour": "red"}');
Query OK, 1 row affected (0.013 sec)

MariaDB [test]> select * from products;
+----+----------+-------+-------+---------------------+
| id | name     | price | stock | attribute           |
+----+----------+-------+-------+---------------------+
|  1 | Blouse   | 17.00 |    15 | {"colour": "white"} |
|  3 | Pingpong | 18.00 |    15 | {"colour": "red"}   |
+----+----------+-------+-------+---------------------+
2 rows in set (0.000 sec)
```
我们可以使用 color 属性，建立索引

```sql
 alter table products add attr_color VARCHAR(32) as (JSON_VALUE(attribute, '$.colour')); 
 create index products_attr_color_ix on products(attr_color);
```
可以使用 attr_color 查询：

```
MariaDB [test]> SELECT * FROM products WHERE attr_color = 'red' or attr_color = 'white';
+----+----------+-------+-------+---------------------+------------+
| id | name     | price | stock | attribute           | attr_color |
+----+----------+-------+-------+---------------------+------------+
|  1 | Blouse   | 17.00 |    15 | {"colour": "white"} | white      |
|  3 | Pingpong | 18.00 |    15 | {"colour": "red"}   | red        |
+----+----------+-------+-------+---------------------+------------+
2 rows in set (0.001 sec)
```

## 更新 JSON 数据

```
MariaDB [test]> update products set attribute = JSON_REPLACE(attribute, '$.colour', 'green') where name = 'Pingpong';
Query OK, 1 row affected (0.019 sec)
Rows matched: 1  Changed: 1  Warnings: 0

MariaDB [test]> select * from products;
+----+----------+-------+-------+---------------------+------------+
| id | name     | price | stock | attribute           | attr_color |
+----+----------+-------+-------+---------------------+------------+
|  1 | Blouse   | 17.00 |    15 | {"colour": "white"} | white      |
|  3 | Pingpong | 18.00 |    15 | {"colour": "green"} | green      |
+----+----------+-------+-------+---------------------+------------+
2 rows in set (0.001 sec)
```



## JSON 函数

此外，除上面用到的 `JSON_REPLACE`  以外，MariaDB 还提供了20多个函数，可以用于操作 JSON 数据，

JSON_REMOVE： 用来删除 JSON 中的某个字段，或者数组中的某个元素

```
SELECT JSON_REMOVE('{"A": 1, "B": 2, "C": {"D": 3}}', '$.C');
+-------------------------------------------------------+
| JSON_REMOVE('{"A": 1, "B": 2, "C": {"D": 3}}', '$.C') |
+-------------------------------------------------------+
| {"A": 1, "B": 2}                                      |
+-------------------------------------------------------+

SELECT JSON_REMOVE('["A", "B", ["C", "D"], "E"]', '$[1]');
+----------------------------------------------------+
| JSON_REMOVE('["A", "B", ["C", "D"], "E"]', '$[1]') |
+----------------------------------------------------+
| ["A", ["C", "D"], "E"]                             |
+----------------------------------------------------+
```



JSON_SET： 用来设置 JSON 中的某个字段的值

```
MariaDB [test]> set @json= '{"name": "Bicen", "address": 20}';
Query OK, 0 rows affected (0.000 sec)

MariaDB [test]>
MariaDB [test]> select JSON_SET(@json, '$.height', 170);
+-------------------------------------------------+
| JSON_SET(@json, '$.height', 170)                |
+-------------------------------------------------+
| {"name": "Bicen", "address": 20, "height": 170} |
+-------------------------------------------------+
1 row in set (0.000 sec)

```



JSON_EXTRACT：用来提取某个字段的值

```
SET @json = '[1, 2, [3, 4]]';

SELECT JSON_EXTRACT(@json, '$[1]');
+-----------------------------+
| JSON_EXTRACT(@json, '$[1]') |
+-----------------------------+
| 2                           |
+-----------------------------+

set @json= '{"name": "Bicen", "address": 20}';
Query OK, 0 rows affected (0.000 sec)

MariaDB [test]> select JSON_EXTRACT(@json, '$.name');
+-------------------------------+
| JSON_EXTRACT(@json, '$.name') |
+-------------------------------+
| "Bicen"                       |
+-------------------------------+
1 row in set (0.000 sec)
```

更多的 JSON 函数可以查看官方文档：https://mariadb.com/kb/en/json-functions/

