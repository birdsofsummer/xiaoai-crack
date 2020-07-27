package echo

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"errors"
	. "../http1"
)

type Echo struct {
	Args    Args    `json:"args"`
	Headers Headers `json:"headers"`
	Origin  string  `json:"origin"`
	URL     string  `json:"url"`
}
type Args map[string]interface{}
type Headers struct {
	Accept       string `json:"Accept"`
	Host         string `json:"Host"`
	UserAgent    string `json:"User-Agent"`
	XAmznTraceID string `json:"X-Amzn-Trace-Id"`
}
func Unmarshal(data []byte) (Echo, error) {
	var r Echo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Echo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func Whoami() (Echo,error){
	var echo Echo
	u:="https://httpbin.org/get"
	d:=map[string]interface{}{

	}
	r,_:=Get(u,d)
	fmt.Println(r.Status)
	if r.StatusCode != 200 {
		return echo, errors.New("network error")
	}
//	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ZZZZZZ")
		return echo,err
	}
	echo, err = Unmarshal(b)
	if err != nil {
		fmt.Println("ZZZZZZ")
		return echo,err
	}
	//fmt.Println(echo.Origin)
	return echo,nil
}

func GetLocalIP1()(string,error){
	ip:=""
	conn, err := net.Dial("udp", "www.google.com.hk:80")
	if err != nil {
		fmt.Println(err.Error())
		return ip,err
	}
	defer conn.Close()
	ip_port:=conn.LocalAddr().String()
	ip=strings.Split(ip_port, ":")[0]
	//fmt.Println(ip)
	return ip,nil
}

func GetLocalIP() (map[string]string){
	var result map[string]string
	result = make(map[string]string)
	ifaces, err := net.Interfaces()

	if err!=nil {
		fmt.Println(err)
	}
	//fmt.Println(ifaces)
	//{3 1500 wlp2s0 9a:ab:a3:2b:d5:77 up|broadcast|multicast}
	for _, iface := range ifaces {
		re,_:=regexp.Compile("docker|hassio|veth")
		is_v:=re.MatchString(iface.Name)

		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || is_v{
			 continue
		}
		addrs, err := iface.Addrs()
		if err!=nil {
			fmt.Println(err)
		}
		//fmt.Println("1111",addrs)
		for _,addr :=range addrs{
			addr1:=strings.SplitN(addr.String(), "/", 2)[0]
			//fmt.Println(iface.Name,addr1)
			result[iface.Name]=addr1
		}
	}
	return result
}


func Test(d []byte) {
  r, _ := Unmarshal(d)
  bytes, _ := r.Marshal()
  fmt.Println(bytes)
}


func Test1(){
	echo,err:=Whoami()
	if err!=nil {

	}
    ip:=echo.Origin
	fmt.Println(ip)

}

func Test2(){
	ip:=GetLocalIP() 
	fmt.Println(ip)

	ip1,_:=GetLocalIP1()
	fmt.Println(ip1)
}

func Main() {
	Test1()
	Test2()
}
