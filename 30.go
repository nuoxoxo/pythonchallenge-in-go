package main


import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func main(){
    data, _ := getbody("yankeedoodle.html")
    data, _ = getbody("yankeedoodle.csv")
    page := string(data)
    page = strings.ReplaceAll(string(data), "\n", "")
    fmt.Println(yell("len/page"), len(page))
    fmt.Println(yell("head/"), page[:100])
    fmt.Println(yell("tail/"), page[len(page) - 100:])
    var nums []string
    for _, line := range strings.Split(page, ",") {
        nums = append(nums, strings.TrimSpace(line))
    }

    fmt.Println(yell("head/"), nums[:10])
    fmt.Println(yell("tail/"), nums[len(nums) - 10:])
    fmt.Println(yell("len/nums"), len(nums))

}

func getbody(sub string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/ring/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth( "repeat", "switch" )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(data), yell("\bbody ends/"))
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string { return Yell + s + Rest }
func cyan(s string) string { return Cyan + s + Rest }


