package web

import(
    "net/http"
)

var SettingsHandle = &View{
    Get:getSettings,
}

func getSettings(w http.ResponseWriter, r *http.Request){
    w.Header().Add("content-type", "text/html")

    m := map[interface{}]interface{}{
        "Title":"Settings",
    }

    templates.ExecuteTemplate(w, "settings", m)
    
    return
}
