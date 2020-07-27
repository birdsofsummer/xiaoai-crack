package http1

import (
	//"crypto/hmac"
	//"crypto/sha1"
	//"crypto/sha256"
	//"strconv"
	//"encoding/base64"
//	iconv "github.com/djimenez/iconv-go"
//	"golang.org/x/text/encoding"
	//"sort"
	"math/rand"
	"strings"
	"time"
	//"os"
	"encoding/json"
	//"log"
	"net/url"
	"fmt"
	//"math"
	"bufio"
	"io"
	//"io/ioutil"
	"net/http"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

var (
  HEADERS=map[string]string{
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:76.0) Gecko/20100101 Firefox/76.0",
        "Accept": "application/json, text/plain, */*",
        "Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"DNT":	"1",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
        "Sec-Fetch-Dest":"empty",
        "Sec-Fetch-Mode":"cors",
        "Sec-Fetch-Site":"same-origin",
		//"Host":	"api.weibo.com",
        // "Referer": "https://api.weibo.com/chat/",
        //"Cookie":"",
		//"Accept-Encoding": "gzip, deflate, br",
	}
)



func GetHeaders(u string) map[string]string{
	var h=map[string]string{}
	for k,v:=range HEADERS{
		h[k]=v
	}
	h["Referer"]=u
	//h["Host"]=""
	//h["Cookie"]=config.Cookie
	//fmt.Println("zzzzzzzz",h)
	return h
}


func To_qs(u string,o map[string]string) string{
	 Url, _ := url.Parse(u)
	 p:=url.Values{}
	 for k, v := range o{
	 	p.Add(k,v)
	 }
	 Url.RawQuery = p.Encode()
	 return Url.String()
}

func To_qs1(o map[string]string) string{
	 p:=url.Values{}
	 for k, v := range o{
	 	p.Add(k,v)
	 }
	 return p.Encode()
}


func Random() string{
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Intn(254))
}

func Pick(s []string) string{
	rand.Seed(time.Now().UnixNano())
	l:=len(s)
	n:=rand.Intn(l)
	return s[n]
}

func Ip() string{
	z:=[]string{}
	for i:=0;i<4;i++{
		t:=Random()
		z=append(z,t)
	}
	return strings.Join(z , ".")
}


func Now() string{
    return fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
}


func To_json(d map[string]string) string{
	m,_:=json.Marshal(d)
    return string(m)
}


func detectContentCharset(body io.Reader) string {
        r := bufio.NewReader(body)
        if data, err := r.Peek(1024); err == nil {
                if _, name, ok := charset.DetermineEncoding(data, ""); ok {
                        return name
                }
        }
        return "utf-8"
}

func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
        if charset == "" {
                charset = detectContentCharset(body)
        }
        e, err := htmlindex.Get(charset)
        if err != nil {
                return nil, err
        }
        if name, _ := htmlindex.Name(e); name != "utf-8" {
                body = e.NewDecoder().Reader(body)
        }
        return body, nil
}





func Get (u string,d map[string]interface{}) (*http.Response,error){
     client := http.Client{}
	 p:=url.Values{}
	 for k, v := range d{
	 	p.Add(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	 }
	 qs:=p.Encode()
	 u1:=u+"?"+qs
	 fmt.Println("get",u1)
     req, _ := http.NewRequest("GET", u1,nil)
     h:=GetHeaders(u1)
     for k,v := range h{
 	    req.Header.Add(k, v)
     }
	 //fmt.Println("hhhhhhhhh",h)
	 return client.Do(req)
 }





