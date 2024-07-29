package main

import (
    "fmt"
    "bytes"
    "compress/bzip2"
    "io/ioutil"
)

func main () {

    un := "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"
    pw := "BZh91AY&SY\x94$|\x0e\x00\x00\x00\x81\x00\x03$ \x00!\x9ah3M\x13<]\xc9\x14\xe1BBP\x91\xf08"
    User := []byte(un)
    readerUser := bzip2.NewReader(bytes.NewReader( User ))
    bytesUser, _ := ioutil.ReadAll(readerUser)
    stringUser := string( bytesUser )
    fmt.Println("usr/", stringUser)

    Pass := []byte(pw)
    readerPass := bzip2.NewReader(bytes.NewReader( Pass ))
    bytesPass, _ := ioutil.ReadAll(readerPass)
    stringPass := string( bytesPass )
    fmt.Println("pwd/", stringPass)
}

