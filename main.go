package main

import (
	"flag"
	"fmt"
	"k8s/client"
	"path/filepath"
)

func main() {
	flag.Var(&client.Mates, "config", "逗号分隔的配置文件")
	flag.Parse()
	if len(client.Mates) < 1 {
		if home := client.HomeDir(); home != "" {
			client.Mates = client.Classmates([]string{filepath.Join(home, ".kube", "config")})
		}
	}
	fmt.Println("classmates ", client.Mates)
}
