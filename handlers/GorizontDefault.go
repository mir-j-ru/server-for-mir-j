package handlers

import(
    "database/sql"
    "strconv"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "ecosystem_mir_j/db"     // или "weather-go-api/db"
)

func GorizontDefault(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    wights := r.URL.Query().Get("wight")
    if wights == ""{
        http.Error(w, "error, wight no", http.StatusBadRequest)
        return
    }

    wight,err:=strconv.ParseInt(wights, 10,64)
    if err!=nil{
        http.Error(w, "error, id not int", http.StatusBadRequest)
        return
    }

    hs := r.URL.Query().Get("height")
    if hs == ""{
        http.Error(w, "error, height no", http.StatusBadRequest)
        return
    }

    h,err:=strconv.ParseInt(hs, 10,64)
    if err!=nil{
        http.Error(w, "error, id not int", http.StatusBadRequest)
        return
    }
    
    err = db.GorizontDefault(dbConn,h, wight)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! horizon default updated!")
}
