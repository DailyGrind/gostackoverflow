### Install dependencies

There are two third party dependecies

* [philgebhardt/mdcat](https://github.com/philgebhardt/mdcat) (fork of [samfoo/mdcat](https://github.com/samfoo/mdcat))
* [philgebhardt/Stack-on-Go](https://github.com/philgebhardt/Stack-on-Go/) (fork of [laktek/Stack-on-Go](https://github.com/laktek/Stack-on-Go))

```bash
go get
```

### Compile code

```bash
go build
```

### Run a search

```
./gostackoverflow "golang"

Setting HTTP headers in Golang (http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang)

Zen 2012-10-10 16:39:13 -0700 PDT

I'm trying to set a header in my Go web server. I'm using gorilla/mux and net/http packages.

I'd like to set Access-Control-Allow-Origin: * to allow cross domain AJAX.

Here's my Go code:


    func saveHandler(w http.ResponseWriter, r *http.Request) {
    // do some stuff with the request data
    }
    
    func main() {
        r := mux.NewRouter()
        r.HandleFunc("/save", saveHandler)
        http.Handle("/", r)
        http.ListenAndServe(":"+port, nil)
    }
    
The net/http package has documentation describing sending http request headers as if it were a client - I'm not exactly sure how to set response headers? 
```