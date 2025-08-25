package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// 定义日志模板
	tmplStr := `【{{ if eq .client_name "pc" }}客户端{{ else }}网页端{{ end }}】登录页：` +
		`【{{ join .select_login_options "、" }}】` +
		`{{ if .update_left_image }}，更新了左侧配图{{ end }}`
	// 自定义函数
	funcMap := template.FuncMap{
		"join": strings.Join,
	}
	// 创建并解析日志模板
	tmpl, err := template.New("login_page").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// 填充数据
	data := map[string]interface{}{
		"client_name": "pc",
		"select_login_options": []string{
			"手机",
			"邮箱",
			"微信",
			"SSO",
			"其他",
		},
		"update_left_image": true,
	}
	// 输出日志模板的结果
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
