- _5/ -_ using `pickle` 🟡
- _7/ -_ read and analyze w/ `res.At(x,y)`
- _9/ -_ draw lines around a polygon
    - new image w/ `image.NewRGBA(image.Rect(sx,sy,ex,ey))`
    - draw w/ `res.Set(x,y,colour)`
- _11/ -_ new image w/ `image.NewRGBA( mypic.Bounds() )`
- GET from a URL
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
