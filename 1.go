package main

import (
    "fmt"
)

func maketrans(s string) string {

    res := []rune{}
    for _, char := range s {
        nxt := int(char)
        if 'a' <= char && char <= 'z' {
            nxt += + 2
            if nxt >= 'z' {
                nxt = nxt - 'z' + 'a' - 1
            }
        }
        //fmt.Printf("%d/ %c %d - %c %d  \n", i, char, int(char), rune(nxt), int(nxt))
        res = append(res, rune(nxt))
    }
    return string (res)
}

func main() {
    s := "g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj."
    fmt.Println("str/", maketrans(s))
    fmt.Println("map/", maketrans("map"))

}
