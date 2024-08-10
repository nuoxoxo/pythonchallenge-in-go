package main

import (
    "os"
    "fmt"
    "os/exec"
    "strings"
    "sync"
)

// md5/ bbb8b499a0eef99b52c7f13f4e78c24b
const yelo, cyan, rest string = "\033[33m", "\033[36m", "\033[0m"

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
        if err != nil { fmt.Println(yelo+"err/24"+rest, err) } // no panic
    }()
    wg.Wait()

    // unzip out.zip
    cmd := exec.Command("unzip", "out.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err := cmd.Run()
    if err != nil { fmt.Println(yelo+"err/out"+rest, err) }

    // unzip mybroken.zip
    cmd = exec.Command("unzip", "mybroken.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err = cmd.Run()
    if err != nil { fmt.Println(yelo+"err/mybroken"+rest, err) }

    // rmv zip
    cmd = exec.Command("rm", "maze.jpg", "out.zip", "mybroken.zip")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    err = cmd.Run()
    if err != nil { fmt.Println(yelo+"err/rm"+rest, err) }

    // correction/ TODO
}

