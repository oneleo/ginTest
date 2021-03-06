package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// set debug mode
	gin.SetMode(gin.DebugMode)

	// set release mode
	//gin.SetMode(gin.ReleaseMode)

	// 禁用控制台顏色，寫入日誌文件時不需要添加顏色
	// gin.DisableConsoleColor()

	// 初始化引擎。
	// 用默認的中間件創建一個gin路由器:
	// logger and recovery (crash-free) middleware
	// 記錄 恢復(不崩潰) 中間件
	r := gin.Default()

	// 自定義模板渲染方式。
	html := template.Must(template.ParseFiles("./file1", "./file2"))
	r.SetHTMLTemplate(html)

	// 自定義分隔符。
	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("./templates/*")

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}
