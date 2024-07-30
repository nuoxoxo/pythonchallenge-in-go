- _5/ -_ using `pickle`
- _9/ -_ using `pillow` to draw polygon - done in Go
- in order to get somewhere
```go
var PAGE string

func main(){
    fmt.Println(PAGE, "\nbody ends/")
}

func init(){
    URL, UUU, PPP := "", "", ""
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL, nil)
    req.SetBasicAuth( UUU, PPP )
    resp, _ := conn.Do(req)

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body) 
    PAGE = string(body)
}
```
