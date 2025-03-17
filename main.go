package main

import "fmt"

// userAgents = []string{
// 	"rest"
// }

func randomUserAgent(){

}

func discoverLinks(){

}
func getRequest(){

}

func resolveRelativeLinks(){

}

func Crawl(targetURL string, baseURL string) []string{
	fmt.Println(targetURL)
	response, _ := getRequest(targetURL)

	links := discoverLinks(response, baseURL)
	foundUrls := []string{}

	for _ , link := range links{
		ok, correctLink := resolveRelativeLinks(link, baseURL)
		if ok {
			if correctLink != ""{
				foundUrls = append(foundUrls, correctLink)
			}
		}
	}
	return foundUrls
}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	baseDomain := "https://www.theguardian.com"
	go func(){worklist <- []string{"https://www.theguardian.com"}}()

	seen := make(map[string]bool)

	for; n>0; n--{

	
	list := worklist

	for _, link := range list{
		if !seen[link]{
			seen[link] = true
			n++
			go func(link string, baseURL string){
				foundLinks := Crawl(link, baseDomain)
				if foundLinks != nil{
					worklist <- foundLinks
				}
			} 
		}
	}
}
}
