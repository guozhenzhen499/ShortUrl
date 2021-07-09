package app

import (
	"fmt"
	"log"
)

var shortHost = "localhost:8082"

func GetShortUrl(long string) string {
	var lastInsertId int64
	insertSql:=fmt.Sprintf("insert into `link` (`url`) value ('%s')",long)
	err:=SMysql.Insert(insertSql,&lastInsertId)
	if err!= nil {
		log.Printf("插入失败：%s",err.Error())
		return ""
	}
	b64:=DecToB64(int(lastInsertId))
	return fmt.Sprintf("%s/%s",shortHost,b64)
}

func GetLongUrl(short string) (string,error) {
	var url string
	dec:=B64ToDec(short)
	querySql:=fmt.Sprintf("select url from link where id=%d",dec)
	rows,err:=SMysql.Query(querySql)
	if err!=nil {
		return "",err
	}
	for rows.Next() {
		err:=rows.Scan(&url)
		if err!= nil {
			return "",err
		}
	}
	return url,nil
}