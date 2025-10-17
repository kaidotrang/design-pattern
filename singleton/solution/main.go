package main

import (
	"design_pattern/singleton/config"
	"fmt"
)

func main() {
	cfg1 := config.GetConfig()
	fmt.Println("Before:", cfg1.AppName)

	// Không thể ghi đè instance vì không truy cập được biến `instance`
	// singleton.instance = &Config{AppName: "HackedApp"} // lỗi: không truy cập được

	cfg2 := config.GetConfig()
	fmt.Println("After:", cfg2.AppName) // vẫn là MyApp
}
