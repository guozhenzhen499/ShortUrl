package app

import (
	"fmt"
	"net/http"
	"strings"
)

func LongToShort(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		_, err := fmt.Fprintln(w, "url参数错误")
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	shortUrl := GetShortUrl(url)
	if shortUrl == "" {
		fmt.Fprintln(w, "链接生成错误")
	}
	fmt.Fprintln(w, shortUrl)
}

func RedirectLongUrl(w http.ResponseWriter, r *http.Request) {

	shortPath := strings.TrimLeft(r.URL.Path, "/")
	if shortPath=="" {
		fmt.Fprintln(w,"短连接路由错误")
		return
	}
	longUrl,err:=GetLongUrl(shortPath)
	fmt.Println(longUrl)
	if err!= nil {
		fmt.Fprintln(w,"获取长链接出错")
		return
	}
	http.Redirect(w,r,longUrl,http.StatusTemporaryRedirect)
}
