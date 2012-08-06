package main

import (
    "fmt"
    "net/http"
    "os"
)

const (
    SSLerror = iota
    HSTSyes
    HSTSno
)

// check_site expects a host and prepends 'https://' onto it to create
// an HTTPS url. A GET request is performed on this URL, and the
// returned headers are checked for the HSTS header.
func Check (site string) int {
    res := HSTSno
    url := "https://" + site

    resp, err := http.Get(url)
    if err != nil {
        return SSLerror
    } else {
        defer resp.Body.Close()
    }

    _, ok := resp.Header["Strict-Transport-Security"]
    if ok {
        res = HSTSyes
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
