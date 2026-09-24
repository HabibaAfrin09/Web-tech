
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func checkURL(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "UNHEALTHY")
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println(url, "HEALTHY")
	} else {
		fmt.Println(url, "UNHEALTHY")
	}
}

func main() {

	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()

		if url == "" {
			continue
		}

		checkURL(url)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}