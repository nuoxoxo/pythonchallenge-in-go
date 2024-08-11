package main

import (
    "os"
    "fmt"
    "os/exec"
    "strings"
    "sync"
    "io"
    "crypto/md5"
    "encoding/hex"
)

const testmd5 string = "bbb8b499a0eef99b52c7f13f4e78c24b"
const yelo, cyan, rest string = "\033[33m", "\033[36m", "\033[0m"
func yell(s string) string {return yelo + s + rest}

func main(){

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()

        // locate go
        cmd := exec.Command("which", "go")
        output, err := cmd.Output()
        if err != nil { panic(fmt.Sprintf("err/which %s", err)) }
        bingo := strings.TrimSpace(string(output))
        fmt.Println("bin/", bingo)

        // run 24
        cmd = exec.Command(bingo, "run", "24.go")
        cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
        err = cmd.Run()
        if err != nil { fmt.Println(yell("err/24"), err) } // no panic
    }()
    wg.Wait()

    // unzip out.zip
    cmd := exec.Command("unzip", "out.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err := cmd.Run()
    if err != nil { fmt.Println(yell("err/out"), err) }

    
    // bruteforce correction
    zipfile, err := os.Open("mybroken.zip")
    fmt.Println(yell("err/zip"), err)
    defer zipfile.Close()
    stat, err := zipfile.Stat()
    fmt.Println(yell("err/stat"), err)
    data := make([]byte, stat.Size())
    _, err = io.ReadFull(zipfile, data)
    fmt.Println(yell("err/read"), err)
    found := false
    for i, b := range data {
        repl := 0
        for repl < 256 {
            //modified := data[:i] + uint8(repl) + data[i + 1:]
            data[i] = uint8(repl)
            attempt := md5.Sum(data)
            newmd5 := hex.EncodeToString(attempt[:])
            if newmd5 == testmd5 {
                _ = os.WriteFile("fixed.zip", data, 0644)
                fmt.Println("i/bugbyte", i, b, "fix/", uint8(repl))
                found = true
            }
            repl++
        }
        if found { break }
        data[i] = b
    }

    // unzip mybroken.zip
    //  broken but the image still shows
    /*
    cmd = exec.Command("unzip", "mybroken.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err = cmd.Run()
    if err != nil { fmt.Println(yell("err/mybroken"), err) }
    */

    // unzip fixed.zip
    cmd = exec.Command("unzip", "fixed.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err = cmd.Run()
    if err != nil { fmt.Println(yell("err/fixed"), err) }

    // rmv zip
    cmd = exec.Command("rm", "maze.jpg", "out.zip", "mybroken.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err = cmd.Run()
    if err != nil { fmt.Println(yell("err/rm"), err) }
}

