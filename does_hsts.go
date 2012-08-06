package main

import (
    "fmt"
    "net/http"
    "os"
)

func check_site (site string) int {
    res := 0
    url := "https://" + site

    resp, err := http.Get(url)
    if err != nil {
        return -1
    }

    defer resp.Body.Close()

    _, ok := resp.Header["Strict-Transport-Security"]
    if ok {
        res = 1
    }
    
    return res
}

func main () {
    for _, site := range os.Args[1:] {
        fmt.Printf("[+] checking whether %v supports HSTS: ", site)
    
        switch res := check_site(site) ; res {
        case -1:
            fmt.Println("SSL failure!")
        case 0:
            fmt.Println("not supported")
        case 1:
            fmt.Println("ok")
        default:
            fmt.Println("oh god there was so much blood")
        }
    }
}
