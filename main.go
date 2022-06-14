package main

import (
	"fmt"
)

func main() {
	r := setupRouter()
	setting(r)

	// Runner
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
