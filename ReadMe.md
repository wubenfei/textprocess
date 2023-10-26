+ 复制versioninfo.json到main.go同级
+ 修改其中版本信息
+ 安装`go get -d github.com/josephspurrier/goversioninfo/cmd/goversioninfo`
+ 终端执行`go generate`
+ golangd中执行`env GOOS=windows GOARCH=amd64 go build -o text_process.exe -ldflags="-linkmode internal"`
+ 需要安装谷歌浏览器：https://www.google.com/chrome/?brand=YTUH&gclid=EAIaIQobChMIvae13d78gQMVOA17Bx2-ggRBEAAYASAAEgK6EPD_BwE&gclsrc=aw.ds
  + 安装时直接保持默认，会安装在C盘：C:\Program Files\Google\Chrome\Application\
+ 如遇到程序闪退，可查看log目录下日志记录，并将问题反馈到作者，经确认问题确实存在后，可补偿使用时长。