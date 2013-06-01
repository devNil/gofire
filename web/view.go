package web

import(
    "net/http"
)

type View struct{
    Post:func(http.ResponseWriter, *http.Request)
    Get:func(http.ResponseWriter, *http.Request)
}

func(v *View) ServeHTTP(w http.ResponseWriter, r *http.Request){
    m := r.Method

    if m == "GET" && v.Get != nil{
        v.Get(w, r)
        return
    }

    if m == "POST" && v.Post != nil{
        v.Post(w, r)
        return
    }

    http.NotFound(w, r)
    return
}
