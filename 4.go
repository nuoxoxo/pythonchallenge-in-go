package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

func main(){

    URL := "http://www.pythonchallenge.com/pc/def/linkedlist.php?nothing="
    TAIL := "12345"
    re := regexp.MustCompile(`and the next nothing is (.*)`)
    count := 0
    for {
        response, _ := http.Get(URL + TAIL)
        body, _ := ioutil.ReadAll(response.Body)
        fmt.Println(count, "\b/", string(body))
        count++
        
        matches := re.FindAllStringSubmatch(string(body), -1)
        if matches == nil {
            if strings.Contains(string(body), "html") {
                return
            }
            tmp, _ := strconv.Atoi(TAIL)
            res := strconv.Itoa(tmp / 2)
            TAIL = res
        } else {
            TAIL = matches[0][1]
        }
    }
}

