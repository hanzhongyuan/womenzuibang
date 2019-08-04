package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/kirinlabs/HttpRequest"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 文档存取位置
var filePath = "/home/batty/桌面/Law_Doc"
var userAgent = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3771.80 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
	"Opera/8.0 (Windows NT 5.1; U; en)",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100922 Firefox/34.0",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 SE 2.X MetaSr 1.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/4.4.3.4000 Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.122 UBrowser/4.0.3214.0 Safari/537.36",
	// 移动端
	"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
	"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; HTC_Wildfire_A3333 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"MQQBrowser/26 Mozilla/5.0 (Linux; U; Android 2.3.7; zh-cn; MB200 Build/GRJ22; CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Opera/9.80 (Android 2.3.4; Linux; Opera Mobi/build-1107180945; U; en-GB) Presto/2.8.149 Version/11.10",
	"Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, like Gecko) Version/6.0.0.337 Mobile Safari/534.1+",
	"Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.0; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/233.70 Safari/534.6 TouchPad/1.0",
	"Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaN97-1/20.0.019; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) BrowserNG/7.1.18124",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Titan)",
	"UCWEB7.0.2.37/28/999",
	"NOKIA5700/ UCWEB7.0.2.37/28/999",
	"Openwave/ UCWEB7.0.2.37/28/999",
	"Mozilla/4.0 (compatible; MSIE 6.0; ) Opera/UCWEB7.0.2.37/28/999",
	// 部分pc端
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
	"Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
	"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
}

// len = 4
var mazilla = []string{
	"Mozilla/5.0 ",
	"Mozilla/4.0 ",
	"Opera/9.80 ",
	"Opera/8.0 ",
}

// len = 18
var version = []string{
	"(Windows NT 5.0) ",
	"(Windows NT 5.1) ",
	"(Windows NT 6.0) ",
	"(Windows NT 6.1) ",
	"(Windows NT 6.2) ",
	"(Windows NT 6.3) ",
	"(Windows NT 10.0) ",
	"(Windows NT 6.2; WOW64) ",
	"(Win64; x64) ",
	"(WOW64) ",
	"(X11; Linux i686) ",
	"(X11; Linux x86_64) ",
	"(X11; Linux i686 on x86_64) ",
	"(Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF" + strconv.Itoa(randNum(999999999999999999)) + ") ",
	"(Macintosh; Intel Mac OS X 10_" + strconv.Itoa(randNum(999999999999999999)) + "_0) ",
	"(Macintosh; PPC Mac OS X 10_" + strconv.Itoa(randNum(999999999999999999)) + "_0) ",
	"(X11; CrOS i686 2268." + strconv.Itoa(randNum(999999999999999999)) + ".0) ",
	"(Linux; U; Android 2.3.7; zh-cn; MB200 Build/GRJ" + strconv.Itoa(randNum(999999999999999999)) + "; CyanogenMod-7) ",
}
var kernel = "AppleWebKit/" + strconv.Itoa(randNum(999999999999999999)) + ".1.38 (KHTML, like Gecko) "
var browser = "Chrome/75.0." + strconv.Itoa(randNum(999999999999999999)) + "." + strconv.Itoa(randNum(999999999999999999)) +
	" Safari/" + strconv.Itoa(randNum(999999999999999999)) + "." + strconv.Itoa(randNum(999999999999999999))

func downloadDoc(url string, docPath string) {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		// 随机使用User-Agent
		//"User-Agent": userAgent[randNum(100)%58],
		"User-Agent": mazilla[randNum(100)%4] + version[randNum(500)%18] + kernel + browser,
	})
	//time.Sleep(100)
	res, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := res.Body()
	if err != nil {
		log.Fatal(err)
	}
	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}
	title := document.Find("body > div.w1200.ma.mb30 > div.clearfix.mt25 > div.fl.w830 > div.detail-page.pr.bg-f8.mt35 > h1")
	content := document.Find("body > div.w1200.ma.mb30 > div.clearfix.mt25 > div.fl.w830 > div.detail-page.pr.bg-f8.mt35 > div.detail-nr")
	fmt.Println("---------------------", title.Text(), "---------------------------------------------------")
	fmt.Println(url)
	fmt.Println(content.Text())
	if content.Text() == "" {
		downloadDoc(url, docPath)
		//input := bufio.NewReader(os.Stdin)
		//input.ReadLine()
	}
	fmt.Println("------------------------------------------------------------------------")
	err = ioutil.WriteFile(filepath.Join(docPath, title.Text()+".txt"), []byte(content.Text()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func getLaw(url, docPath string) {
	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		// 随机使用User-Agent
		//"User-Agent": userAgent[randNum(1000)%58],
		"User-Agent": mazilla[randNum(100)%4] + version[randNum(500)%18] + kernel + browser,
	})
	//time.Sleep(100)
	res, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := res.Body()
	if err != nil {
		log.Fatal(err)
	}
	//body > div.w1200.ma.mb50 > div.clearfix.mt25 > div.fl.w830 > div.clearfix.mt30 > div > div > em
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
	}
	em := doc.Find("body > div.w1200.ma.mb50 > div.clearfix.mt25 > div.fl.w830 > div.clearfix.mt30 > div > div > em")
	list := strings.Split(em.Text(), " / ")
	if len(list) != 2  || len(list) == 0 {
		getLaw(url, docPath)
	} else {
		fmt.Println("list = ", list, "url = ", url)
		document, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
		if err != nil {
			log.Fatal(err)
		}
		url_detail := document.Find(`body > div.w1200.ma.mb50 > div.clearfix.mt25 > div.fl.w830 > div.clearfix.mt30 > div > ul > li`)
		for i := 1; i <= url_detail.Size(); i++ {
			temp := url_detail.Find("body > div.w1200.ma.mb50 > div.clearfix.mt25 > div.fl.w830 > div.clearfix.mt30 > div > ul > li:nth-child(" +
				strconv.Itoa(i) + ") > h3 > a")
			val, exists := temp.Attr("href")
			if exists != false {
				downloadDoc("https://www.66law.cn"+val, docPath)
			}
		}
		page, _ := strconv.Atoi(list[0])
		total, _ := strconv.Atoi(list[1])
		if page < total {
			page++
			str := getBetweenStr(url, "https://www.66law.cn/tiaoli/", ".aspx")
			next := "https://www.66law.cn/tiaoli/" + str[0:strings.LastIndex(str, "p")] + "p" + strconv.Itoa(page) + ".aspx"
			//fmt.Println(next)
			getLaw(next, docPath)
		}
	}
}

func solve(urlPath, docPath string) {
	f, err := os.Open(urlPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
		getLaw(line, docPath)
		if err != nil || io.EOF == err {
			break
		}
	}
}

func main() {
	// 法律分类列表
	dirPath := "/home/batty/桌面/lawList"
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		err := os.MkdirAll(filepath.Join(filePath, strings.ReplaceAll(file.Name(), ".txt", "")), os.ModePerm)
		if err != nil {
			panic(err)
		}
		fmt.Println(file.Name())
		solve(filepath.Join(dirPath, file.Name()), filepath.Join(filePath, strings.ReplaceAll(file.Name(), ".txt", "")))
	}
}

// 备用功能
func getBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func randNum(max int) int {
	rad := rand.New(rand.NewSource(time.Now().Unix()))
	return rad.Intn(max)
}
