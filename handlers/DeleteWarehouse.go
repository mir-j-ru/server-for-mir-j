package handlers

import(
    "strconv"
    "database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "ecosystem_mir_j/db"     // или "weather-go-api/db"
)

func DeleteWarehouse(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    idstr := r.URL.Query().Get("id")
    if idstr == ""{
        http.Error(w, "error, id no", http.StatusBadRequest)
        return
    }

    id,err:=strconv.ParseInt(idstr, 10,64)
    if err!=nil{
        http.Error(w, "error, id not int", http.StatusBadRequest)
        return
    }


    err = db.DeleteWarehouse(dbConn,id)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! warehouse deleted!")
}