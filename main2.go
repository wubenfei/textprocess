package main

//
// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"github.com/cilidm/toolbox/file"
// 	"github.com/dustin/go-humanize"
// 	"github.com/kirinlabs/HttpRequest"
// 	"github.com/tidwall/gjson"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )
//
// var year *bool
//
// func init() {
// 	// flag.BoolVar(year, "y", false, "是否按年归类")
// }
// func main() {
// 	flag.Parse()
// 	SpiderDY(year)
// }
//
// // --------------spider----------------
// var req *HttpRequest.Request
// var downloadDir = "download"
//
// func init() {
// 	req = HttpRequest.NewRequest()
// 	req.SetHeaders(map[string]string{
// 		"User-Agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36",
// 	})
// 	req.CheckRedirect(func(req *http.Request, via []*http.Request) error {
// 		return http.ErrUseLastResponse /* 不进入重定向 */
// 	})
// }
// func SpiderDY(year *bool) {
// 	// lines, err := ReadLine("url.txt")
// 	// if err != nil {
// 	// 	os.Create("url.txt")
// 	// 	log.Fatal("未找到url.txt文件，已自动创建，请在同目录下url.txt文件加入分享链接，每个分享链接一行")
// 	// }
// 	// for _, line := range lines {
// 	// 	reg := regexp.MustCompile(`[a-z]+://[\S]+`)
// 	// 	url := reg.FindAllString(line, -1)[0]
// 	// 	resp, err := req.Get(url)
// 	// 	defer resp.Close()
// 	// 	if err != nil {
// 	// 		fmt.Errorf(err.Error())
// 	// 		continue
// 	// 	}
// 	// 	if resp.StatusCode() != 302 {
// 	// 		continue
// 	// 	}
// 	// 	location := resp.Headers().Values("location")[0]
// 	// 	regNew := regexp.MustCompile(`(?:sec_uid=)[a-z,A-Z，0-9, _, -]+`)
// 	// 	sec_uid := strings.Replace(regNew.FindAllString(location, -1)[0], "sec_uid=", "", 1)
// 	sec_uid := "MS4wLjABAAAA_tn0xx5dNiKWmJtzPNGNmSbWI6c2-qqJfddiZ1yCJ-F_k7zgCnBYnnylPjjz8Lfo"
// 	respIes, err := req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/user/info/?sec_uid=%s", sec_uid))
// 	defer respIes.Close()
// 	if err != nil {
// 		fmt.Errorf(err.Error())
// 		return
// 	}
// 	body, err := respIes.Body()
// 	result := gjson.Get(string(body), "user_info.nickname").String()
// 	dirPath := fmt.Sprintf("%s/%s/", downloadDir, result)
// 	err = file.IsNotExistMkDir(dirPath)
// 	if err != nil {
// 		fmt.Errorf(err.Error())
// 		return
// 	}
// 	GetByMonth(sec_uid, dirPath, year)
// }
// func GetByMonth(sec_uid, dirPath string, year *bool) {
// 	y := 2018
// 	nowY, _ := strconv.Atoi(time.Now().Format("2006"))
// 	nowM, _ := strconv.Atoi(time.Now().Format("01"))
// 	for i := y; i <= nowY; i++ {
// 		for m := 1; m <= 12; m++ {
// 			var (
// 				begin int64
// 				end   int64
// 			)
// 			if i == nowY && m > nowM {
// 				break
// 			}
// 			begin = GetMonthStartAndEnd(strconv.Itoa(i), strconv.Itoa(m))
// 			if m == 12 {
// 				end = GetMonthStartAndEnd(strconv.Itoa(i+1), "1")
// 			} else {
// 				end = GetMonthStartAndEnd(strconv.Itoa(i), strconv.Itoa(m+1))
// 			}
// 			resp, err := req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&signature=PtCNCgAAXljWCq93QOKsFT7QjR",
// 				sec_uid, begin, end))
// 			defer resp.Close()
// 			if err != nil {
// 				fmt.Errorf(err.Error())
// 				continue
// 			}
// 			body, err := resp.Body()
// 			r := gjson.Get(string(body), "aweme_list").Array()
// 			if len(r) > 0 {
// 				if *year {
// 					dirPath = fmt.Sprintf("%s%d/", dirPath, i)
// 					err = file.IsNotExistMkDir(dirPath)
// 					if err != nil {
// 						fmt.Errorf(err.Error())
// 						continue
// 					}
// 				}
// 				for n := 0; n < len(r); n++ {
// 					videotitle := gjson.Get(string(body), fmt.Sprintf("aweme_list.%d.desc", n)).String()
// 					videourl := gjson.Get(string(body), fmt.Sprintf("aweme_list.%d.video.play_addr.url_list.0", n)).String()
// 					err = DownloadFile(dirPath+videotitle+".mp4", videourl)
// 					if err != nil {
// 						fmt.Errorf(err.Error())
// 						continue
// 					}
// 				}
// 			}
// 		}
// 	}
// }
//
// // -----------------util----------------------
// // 下载文件显示进度条
// type WriteCounter struct {
// 	Name  string
// 	Total uint64
// }
//
// func (wc *WriteCounter) Write(p []byte) (int, error) {
// 	n := len(p)
// 	wc.Total += uint64(n)
// 	wc.PrintProgress()
// 	return n, nil
// }
// func (wc WriteCounter) PrintProgress() {
// 	fmt.Printf("\r%s", strings.Repeat(" ", 35))
// 	fmt.Printf("\r 【%s】 Downloading... %s complete", wc.Name, humanize.Bytes(wc.Total))
// }
// func DownloadFile(fileName string, url string) error {
// 	req, err := http.NewRequest("GET", url, nil)
// 	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36")
// 	out, err := os.Create(fileName + ".tmp")
// 	if err != nil {
// 		return err
// 	}
// 	resp, err := (&http.Client{}).Do(req)
// 	if err != nil {
// 		out.Close()
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	counter := WriteCounter{}
// 	fileNameSp := strings.Split(fileName, "/")
// 	counter.Name = fileNameSp[len(fileNameSp)-1]
// 	if _, err = io.Copy(out, io.TeeReader(resp.Body, &counter)); err != nil {
// 		out.Close()
// 		return err
// 	}
// 	fmt.Print("\n")
// 	out.Close()
// 	if err = os.Rename(fileName+".tmp", fileName); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// // GetMonthStartAndEnd 获取月份的第一天和最后一天
// func GetMonthStartAndEnd(myYear string, myMonth string) int64 {
// 	// 数字月份必须前置补零
// 	if len(myMonth) == 1 {
// 		myMonth = "0" + myMonth
// 	}
// 	yInt, _ := strconv.Atoi(myYear)
// 	timeLayout := "2006-01-02 15:04:05"
// 	loc, _ := time.LoadLocation("Local")
// 	theTime, _ := time.ParseInLocation(timeLayout, myYear+"-"+myMonth+"-01 00:00:00", loc)
// 	newMonth := theTime.Month()
// 	t1 := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
// 	return t1
// }
//
// // 按行读取配置
// func ReadLine(fileName string) (lines []string, err error) {
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	buf := bufio.NewReader(f)
// 	for {
// 		line, err := buf.ReadString('\n')
// 		line = strings.TrimSpace(line)
// 		lines = append(lines, line)
// 		if err != nil {
// 			if err == io.EOF {
// 				return lines, nil
// 			}
// 			return nil, err
// 		}
// 	}
// }
