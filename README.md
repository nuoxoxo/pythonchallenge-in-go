# notes

- _5/ -_ using `pickle` 🟡
- _7/ -_ read and analyze w/ `res.At(x,y)`
- _9/ -_ draw lines around a polygon
    - new image w/ `image.NewRGBA(image.Rect(sx,sy,ex,ey))`
    - draw w/ `res.Set(x,y,colour)`
- _11/ -_ new image w/ `image.NewRGBA( mypic.Bounds() )`
- _16/ -_ longest substring of repeats 👉 YES
    - not _move_ the pink segment but _Rotate_
- _18/ -_ now shows p3 correctly ~~common.png not written properly~~
    - strategy: imagine 2 byte slices, a and b
- _19/ -_  `Big endian` is we write the most significant byte first - `Little endian`, the least significant byte first.
```
create 3 new bytes slices, such that
first one contains what is in a, but not in b
second one contains what is in b and not in a
third one contains what is in both a and b
```
- GET from a URL
```go
var PAGE string

func main(){
    fmt.Println(PAGE, "\nbody ends/")
}

func init(){
    URL := ""
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL, nil)
    req.SetBasicAuth("", "")
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
19| _indian -_
18| _brightness - beurre:fly_
17| _violin - balloons_
16| _romance_
15| _mozart_
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
0 | ` `
