# notes

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

# keys

level | keyword
----- | -----------
21|
20|
19|
18|
17|
16|
15| ~~?~~
14| _cat - his/her name_
13| _italy_
12| _disproportional_
11| _evil_
10| _5808_
9 | _bull_
8 | ` `
7 | _integrity_
6 | _oxygen_
5 | _channel_
4 | _peak_
3 | _linkedlist_
2 | _eqaulity_
1 | _ocr_
0 | 2^38
