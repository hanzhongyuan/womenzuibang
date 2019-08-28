package main

import (
	"blackBatty/easyGo/HttpRequest"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Data []struct {
	Infostaticurl string `json:"infostaticurl"`
	Listtitle     string `json:"listtitle"`
	Releasedate   string `json:"releasedate"`
}

const URL = "http://www.chinalaw.gov.cn"
const FILEPATH = "C:\\Users\\liuliang\\Desktop\\law"

func getDoc(url, path string, i int) string {
	req := HttpRequest.NewRequest()
	res, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(res)))
	if err != nil {
		log.Fatal(err)
	}
	title := doc.Find("body > div.main_bg > div.main_con > div.con_bgtiao > div.con_bt")
	content := doc.Find("#content > span")
	fmt.Println(url)
	fmt.Println(title.Text())
	filename := filepath.Join(path, strconv.Itoa(i)+".txt")
	ds, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()
	ds.WriteString(content.Text())

	str := strconv.Itoa(i) + "---" + title.Text()
	return str
}

func solve(jsonUrl, path string, ch chan bool) {
	req := HttpRequest.NewRequest()
	res, err := req.Get(jsonUrl)
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	err = json.Unmarshal(res, &data)
	if err != nil {
		log.Fatal(err)
	}
	url := ""
	str := ""
	i := 1
	for _, vv := range data {
		url = URL + vv.Infostaticurl
		str += getDoc(url, path, i) + "\n"
		i++
	}

	dir := filepath.Join(path, "目录")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	ds, err := os.Create(filepath.Join(dir, "title.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()
	ds.WriteString(str)
	ch <- true
}

func main() {
	list := map[string]string{
		"法律":         "http://www.chinalaw.gov.cn/json/592_1.json",
		"司法协助相关法律法规": "http://www.chinalaw.gov.cn/json/357_1.json",
		"民商事司法协助条约":  "http://www.chinalaw.gov.cn/json/358_1.json",
		"刑事司法协助条约":   "http://www.chinalaw.gov.cn/json/359_1.json",
		"引渡条约":       "http://www.chinalaw.gov.cn/json/360_1.json",
		"被判刑人移管条约":   "http://www.chinalaw.gov.cn/json/361_1.json",
		"司法协助相关公约":   "http://www.chinalaw.gov.cn/json/362_1.json",
		"行政法规":       "http://www.chinalaw.gov.cn/json/593_1.json",
		"法规解读":       "http://www.chinalaw.gov.cn/json/596_1.json",
		"部门规章":       "http://www.chinalaw.gov.cn/json/594_1.json",
		"地方政府规章":     "http://www.chinalaw.gov.cn/json/595_1.json",
	}
	ch := make(chan bool)
	i := 0
	for k, v := range list {
		path := filepath.Join(FILEPATH, k)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		i++
		go solve(v, path, ch)
	}
	for j := 0; j < i; j++ {
		<-ch
	}
}
