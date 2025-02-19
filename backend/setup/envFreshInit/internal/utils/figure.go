package utils

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func ShowFigure(str string, version string, auth string) {
	// 创建一个新的 figure 实例，使用标准字体
	figure.NewFigure(str, "slant", true).Print()
	lens, _ := fmt.Printf("Setup backend program v%s\tby %s\n", version, auth)
	for i := 0; i < lens+10; i++ {
		fmt.Print("=")
	}
	fmt.Println()

}
