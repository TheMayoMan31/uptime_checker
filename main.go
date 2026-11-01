package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {

	_, err := os.Stat("sites.txt")
	if errors.Is(err, os.ErrNotExist) {
		os.Create("sites.txt")
	}

	content, err := os.ReadFile("sites.txt")

	sites := []string{}


	sites = strings.Fields(string(content))

	for i:=0; i<len(sites);i++ {
		start := time.Now()
		resp, err := http.Get(sites[i])
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println(sites[i] , resp.Status, time.Since(start))
	}
}
