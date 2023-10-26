package main

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// 	"time"
// )
//
// func main() {
// 	secId := getSecId()
// 	getVideoUrl(secId)
// 	getUserInfo(secId)
// 	tryGetVideoUrl(secId)
// }
//
// func tryGetVideoUrl(secId string) {
// 	// 抖音已经在改接口中取消了视频链接，只能获取到封面的url
// 	cursor := time.Now().UnixMilli()
// 	url := fmt.Sprintf("https://m.douyin.com/web/api/v2/aweme/post/?reflow_source=reflow_page&sec_uid=%s&count=21&max_cursor=%d", secId, cursor)
// 	cli := http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")
// 	rsp, err := cli.Do(req)
// 	if err != nil {
// 		return
// 	}
// 	content, err := io.ReadAll(rsp.Body)
// 	fmt.Println(content)
// }
//
// func getUserInfo(secId string) {
// 	url := fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/user/info/?sec_uid=%s", secId)
// 	cli := http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")
// 	rsp, err := cli.Do(req)
// 	if err != nil {
// 		return
// 	}
// 	content, err := io.ReadAll(rsp.Body)
// 	fmt.Println(content)
// }
//
// func getVideoUrl(secId string) {
// 	cursor := time.Now().UnixMilli()
// 	url := fmt.Sprintf("https://m.douyin.com/web/api/v2/aweme/post/?reflow_source=reflow_page&sec_uid=%s&count=21&max_cursor=%d", secId, cursor)
// 	fmt.Println(url)
// 	cli := http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36")
// 	rsp, err := cli.Do(req)
// 	if err != nil {
// 		return
// 	}
// 	content, err := io.ReadAll(rsp.Body)
// 	if err != nil {
// 		return
// 	}
// 	awemeList := &AwemeList{}
// 	json.Unmarshal(content, &awemeList)
// 	fmt.Println(awemeList)
// }
//
// func getSecId() string {
// 	sec_id := ""
// 	// 抖音主页： https://v.douyin.com/idXMKVwn/
// 	userDomain := "https://v.douyin.com/idXMKVwn/"
// 	cli := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 		userDomain = req.URL.Path
// 		return nil
// 	},
// 	}
// 	req, _ := http.NewRequest("GET", userDomain, nil)
// 	_, err := cli.Do(req)
// 	if err != nil {
// 		return ""
// 	}
// 	sec_id = strings.Split(userDomain, "user/")[1]
// 	fmt.Println(sec_id)
// 	return sec_id
// }
//
// // AwemeList 接口中已经取消了视频链接，因此项目止步
// type AwemeList struct {
// 	AwemeList []struct {
// 		AwemeId string      `json:"aweme_id"`
// 		Desc    string      `json:"desc"` // 视频标题
// 		ChaList interface{} `json:"cha_list"`
// 		Video   struct {    // 视频封面
// 			Cover struct {
// 				Uri     string   `json:"uri"`
// 				UrlList []string `json:"url_list"`
// 			} `json:"cover"`
// 			BitRate interface{} `json:"bit_rate"`
// 		} `json:"video"`
// 	} `json:"aweme_list"`
// }
