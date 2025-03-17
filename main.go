package main

import ("fmt"
"github.com/PuerkitoBio/goquery"
"math/rand"
"net/http"
"strings"
"time"
)
userAgents = []string{
	"rest"
}

func randomUserAgent(){
 rand.Seed(time.Now().Unix)
 randomNum := rand.Int() %len(userAgents)
 return userAgents[randomNum]
}

func discoverLinks(response &http.Response, baseURL string) []string{
	if response != nil {
		doc, _ := goquery.NewDocumentFromResponse(response)
		foundUrls := []string{}
		if doc != nil{
			doc.Find("a").Each(func(i int, s *goquery.Selection)){
				res,_ := s.Attr("href")
				foundUrls = append(foundUrls, res)
			}
		}
		return foundUrls
	} else{
		return []string{}
	}
}
func getRequest(targetURL string)(*http.Response, error){
	client := &http.Client{}

	req,_err := http.NewRequest("GET", targetURL,nil)
	if err != nil {
		return nil, err
	} else{
		return res, nil
	}

	req.Header.Set("User-Agent", randomUserAgent)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	} else{
		return res, nil
	}
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
	
	// ParseHTML(response)
	return foundUrls
}

// func ParseHTML(response *http.Response){
	//Can take the information and parse it. maybe use it to get headlines or articles on certain topics? future project idea. 
	//Tutorial 32 explains parsing results
// }

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
