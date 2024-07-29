package main

import (
    "fmt"
    "reflect"
    "os"
    "io/ioutil"
    "strconv"
)

func main(){
    content, _ := ioutil.ReadFile("files/evil2.gfx")
    fmt.Println(content[:42], reflect.TypeOf(content) ) // uint8
    N := len(content)
    i := 0
    for i < 5 {
        name := "files/0" + strconv.Itoa(i) + ".jpg"
        outfile, _ := os.Create( name )
        j := i
        for j < N {
            outfile.Write( content[j : j + 1] )
            j += 5
        }
        outfile.Close()
        i++
    }
}



