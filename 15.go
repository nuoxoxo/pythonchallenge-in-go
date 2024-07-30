package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
    "strconv"
)

func main(){
    // get to know the lib
    REF := "2006-01-02"
    todo := "1986-07-11"
    date, _ := time.Parse(REF, todo)
    fmt.Println("date/", date)
    fmt.Println("wkdt/", date.Weekday())

    // find leapyear
    mid := 0
    end := 100
    leaps := []int{}
    for mid < end {
        nn := mid * 10 + 1006
        if nn % 4 == 0 {
            leaps = append(leaps, nn)
        }
        mid++
    }
    fmt.Println("\nleaps", leaps)

    // find the year where jan 26 is a monday
    monyears := []int{}
    for _, year := range leaps {
        todo = strconv.Itoa(year) + "-01-26"
        date, _ = time.Parse(REF, todo)
        wkdt := date.Weekday()
        if wkdt == time.Monday/*"Monday"*/ {
            monyears = append(monyears, year)
            fmt.Println("date/", date, "wkdt/", wkdt, "- Found!")
        }
    }
    fmt.Println("\nmonyears/", monyears, "len/", len(monyears))

    // 2nd youngest
    res := monyears[len(monyears) - 2]
    fmt.Println("2nd/", res)
    fmt.Println("\nfind out yourself what to do next")
}

// 

func init(){
    URL := "http://www.pythonchallenge.com/pc/return/uzi.html"
    ups, mid := "hugefile", 4
    conn := & http.Client{}
    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {fmt.Println("err/", err)}

    req.SetBasicAuth(ups[:mid], ups[mid:])
    resp, _ := conn.Do(req)
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body), "\nbody ends/\n\n")

}

