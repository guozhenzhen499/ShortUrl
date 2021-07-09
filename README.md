# ShortUrl
go实现短链接
## 1.建表
```
    CREATE TABLE `link` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
      `url` varchar(200) NOT NULL DEFAULT '' COMMENT '链接',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='链接表'
```
