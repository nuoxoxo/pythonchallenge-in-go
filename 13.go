package main

import (
    "fmt"
    _"net/http"
    "github.com/kolo/xmlrpc"
    "reflect"
)

func main() {

    URL := "http://www.pythonchallenge.com/pc/phonebook.php"
    conn, err := xmlrpc.NewClient(URL, nil)
    if err != nil {
        fmt.Println("xmlrpc.NewClient/", err)
    }

    // get listmethods

    cmds := []string{}
    // note: system.listMethods is a RPC standard
    _ = conn.Call("system.listMethods", nil, & cmds)
    fmt.Println("\ncmds/", reflect.TypeOf(cmds))

    for _, cmd := range cmds {fmt.Println("cmd/", cmd)}
    /*
    cmd/ phone
    cmd/ system.listMethods
    cmd/ system.methodHelp
    cmd/ system.methodSignature
    cmd/ system.multicall
    cmd/ system.getCapabilities
    */

    res := []string{}
    // note: system.listMethods is a RPC standard
    _ = conn.Call(cmds[0], nil, & res)
    fmt.Println("phone/", res, "\nend/\n")

    // try: cmd/ system.methodHelp
    //      cmd/ system.methodSignature
    var _help interface{}
    err = conn.Call(cmds[2], "phone", & _help)
    fmt.Println(cmds[2], _help, err, "\nend/\n")

    /*
    _signature := [][]string{}
    _signature = append(_signature, []string{})
    err = conn.Call(cmds[3], "phone", & _signature)
    if err != nil { fmt.Println("err/", err) }
    fmt.Println(cmds[3], _signature, "\nend/\n")
    */

    // BUG
    // Bugfix: `system.methodSignature [[string string]]` is indeed the response
    //  param   : str
    //  return  : str

    /*
    var _signature2 interface{}
    err = conn.Call(cmds[3], "phone", & _signature2)
    if err != nil { fmt.Println("err/", err) }
    fmt.Println(cmds[3], _signature2, "\nend/\n")
    fmt.Println(reflect.TypeOf(_signature2))
    */

    i := 2
    for i < 4 {
        var ret interface{}
        err = conn.Call(cmds[i], "phone", & ret)
        if err != nil { fmt.Println("err/", err) }
        fmt.Println(cmds[i], ret, "\nend/\n")
        i++
    }

    var _phone interface{}
    err = conn.Call(cmds[0], "Bert", & _phone)
    if err != nil { fmt.Println("err/", err) }
    fmt.Println("phone/", _phone, "\nend/\n")
    fmt.Println("type/", reflect.TypeOf(_phone))

}

