package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlConfig struct {
	Host 		string
	Port        int
	Username    string
	Password    string
	Database    string
	maxOpen     int
	maxIdle     int
	maxLifetime time.Duration
}

var ShortUrlMysqlConfig = &MysqlConfig{
	"127.0.0.1",
	3306,
	"root",
	"123456",
	"shortUrl",
	100,
	10,
	60,
}

type Mysql struct {
	db     *sql.DB
	isInit bool
}

var SMysql = &Mysql{}

//init
func (mysql *Mysql) init() error {
	databaseConfig:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",ShortUrlMysqlConfig.Username,ShortUrlMysqlConfig.Password,
		ShortUrlMysqlConfig.Host,ShortUrlMysqlConfig.Port,ShortUrlMysqlConfig.Database)
	db,err:=sql.Open("mysql",databaseConfig)
	if err!= nil {
		return err
	}

	db.SetMaxOpenConns(ShortUrlMysqlConfig.maxOpen)//最大连接数
	db.SetMaxIdleConns(ShortUrlMysqlConfig.maxIdle)//限制连接数
	db.SetConnMaxLifetime(ShortUrlMysqlConfig.maxLifetime*time.Second)//最大连接周期
	db.Ping()
	mysql.db=db
	mysql.isInit=true

	return nil
}

func (mysql *Mysql) Query(querySql string) (*sql.Rows,error) {
	if !mysql.isInit {
		if err:=mysql.init();err!=nil {
			return nil,err
		}
	}

	res,err:=mysql.db.Query(querySql)
	if err!=nil {
		return nil,err
	}

	return res,nil
}

func (mysql *Mysql) Insert(insertSql string,lastInsertId *int64) error {
	if !mysql.isInit {
		if err:=mysql.init();err!=nil {
			return err
		}
	}
	res,err:=mysql.db.Exec(insertSql)
	if err!= nil {
		return err
	}

	id,err := res.LastInsertId()
	if err!= nil {
		return err
	}
	*lastInsertId=id
	return nil
}