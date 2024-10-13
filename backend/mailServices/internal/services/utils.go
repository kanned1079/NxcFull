package services

//// 渲染 HTML 模板，替换占位符
//func (v *VerifyEmailData) RenderEmailTemplate(templatePath string) (string, error) {
//	log.Println("renderEmailTemplate", v)
//	// 解析模板文件
//	tmpl, err := template.ParseFiles(templatePath)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse template file: %v", err)
//	}
//
//	// 渲染模板，替换占位符
//	var rendered bytes.Buffer
//	if err := tmpl.Execute(&rendered, v); err != nil {
//		return "", fmt.Errorf("failed to execute template: %v", err)
//	}
//
//	return rendered.String(), nil
//}
