package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	domain := strings.Join(r.URL.Query()["domain"], " ")

	fmt.Println("domain: ", domain)
	cmd := exec.Command("bash", "-c", fmt.Sprintf("dig %s", domain))
	cmdOutput, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", cmdOutput)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
