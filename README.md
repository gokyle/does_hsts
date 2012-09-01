## DOES_HSTS

This is a very simple Go command line utility that, given a list of sites
to check, will determine whether those sites support 
[HSTS](https://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security).

### Example usage
```bash
<monalisa: ~> $ does_hsts conformal.com google.com reddit.com
[+] checking whether conformal.com supports HSTS: ok
[+] checking whether google.com supports HSTS: not supported
[+] checking whether reddit.com supports HSTS: SSL failure!
```

### Installing
does_hsts requires the go compiler to be installed. See the Go language
website's [instructions](http://golang.org/doc/install) for installation. 
To build, run `make`; installation defaults to /usr/local/bin and can be 
installed with `make install`. The `PREFIX` environment variable can be 
used to change the installation directory: `PREFIX=${HOME} make install` 
would install to `${HOME}/bin`.
