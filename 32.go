package main

import (
    "github.com/ODAEL/nonogo/holder"
    "github.com/ODAEL/nonogo/solver"
    "fmt"
    "net/http"
    "io/ioutil"
    "regexp"
    "strings"
    "strconv"
)

var lines []string

func main() {

    // should be a func
    /*
    */
    solve_nonogram(lines)

    // up.html

    data, _ := getbody("rock/up.html", "kohsamui", "thailand")
    fmt.Println(string(data))

    // access embedded url 2 - up.txt
    re := regexp.MustCompile(`(?s)<a href="(.*?)">`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "sub/" + Rest, sub)
    data, _ = getbody("rock/" + sub, "kohsamui", "thailand")
    fmt.Println(Cyan + "sub/data on newline" + Rest)
    fmt.Println(string(data))
    lines = strings.Split(string(data), "\n")

    // solve 2nd nonogram
    solve_nonogram(lines)

    // final result: python.html
}

func init() {

    // main page
    data, _ := getbody("rock/arecibo.html", "kohsamui", "thailand")
    fmt.Println(string(data))
    yell("body/ends")

    // access embedded url
    re := regexp.MustCompile(`(?s)Fill in the blanks <!-- for (.*?) -->`)
    sub := re.FindAllStringSubmatch(string(data), -1)[0][1]
    fmt.Println(Cyan + "sub/" + Rest, sub)
    data, _ = getbody("rock/" + sub, "kohsamui", "thailand")
    fmt.Println(Cyan + "sub/data on newline" + Rest)
    fmt.Println(string(data))
    lines = strings.Split(string(data), "\n")
}

func solve_nonogram(lines []string) {
    dime, hori, vert := false, false, false
    left, top := [][]int{}, [][]int{} // leftbox, topbox
    bound := []int{}
    for i, line := range lines {
        if len(line) < 1 {continue}
        if strings.Contains(line, "Dimensions") {
            dime = true
        } else if strings.Contains(line, "Horizontal") {
            hori = true
        } else if strings.Contains(line, "Vertical") {
            vert = true
        } else {
            nums := []int{}
            fields := strings.Fields(line)
            for _, num := range fields {
                n, _ := strconv.Atoi(num)
                nums = append(nums, n)
            }
            fmt.Println(i, nums)
            if vert {
                top = append(top, nums)
            } else if hori {
                left = append(left, nums)
            } else if dime && len(line) != 0 {
                bound = nums
            }
        }
    }

    fmt.Println("boundbox/on newline \n", top, "\n", left, "\n", bound)

    // play w/ nonogram
    topbox := holder.Box{Numbers: top} // define boundary box
    leftbox := holder.Box{Numbers: left}
    nonogram := holder.BuildEmptyNonogram(bound[0], bound[1]) // build nonogram
    nonogram.TopBox = topbox
    nonogram.LeftBox = leftbox
    solver.Solve( & nonogram) // solve 
    printer := holder.CmdNonogramPrinter{ Nonogram: nonogram }
    printer.PrintField() // print

}

func getbody(sub, u, p string) ( []uint8, error ) {
    URL := "http://www.pythonchallenge.com/pc/"
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL + sub, nil)
    req.SetBasicAuth(u, p)
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    data, _ := ioutil.ReadAll(resp.Body)
    return data, nil
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) { fmt.Println( Yell + s + Rest )}
func cyan(s string) { fmt.Println( Cyan + s + Rest )}

