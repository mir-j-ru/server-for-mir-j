package handlers

import(
    "database/sql"
    //"encoding/json"
    "fmt"
    //"log"
    "net/http"
    //"os"
    "strconv"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "ecosystem_mir_j/db"     // или "weather-go-api/db"
)

func PutWarehoyse(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    category := value.Get("category")
    number := value.Get("id")

    if category == ""{
        http.Error(w, "error create warehouse: not category", http.StatusBadRequest)
        return
    }

    if number == ""{
        http.Error(w, "error create warehouse: not number", http.StatusBadRequest)
        return
    }

    id, err :=  strconv.ParseInt(number, 10, 64)
    if err != nil{
        http.Error(w, "error update warehouse: id not int")
        return
    }

    err = db.PutWarehoyse(dbConn, category, number)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! warehouse added!")
}