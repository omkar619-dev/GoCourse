package basics

import "fmt"
import "net/http"

func main() {
	fmt.Println("Starting.")

	resp,err:= http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
}