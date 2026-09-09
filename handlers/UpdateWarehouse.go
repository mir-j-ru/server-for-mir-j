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

func UpdateWarehouse(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    number := value.Get("id")
    quantitystr := value.Get("quantity")

    if number == ""{
        http.Error(w, "error create warehouse: not number", http.StatusBadRequest)
        return
    }

    if quantitystr == ""{
        http.Error(w, "error create warehouse: not quantity", http.StatusBadRequest)
        return
    }

    quantity, err := strconv.ParseInt(quantitystr, 10, 64)
    if err != nil{
        http.Error(w, "error create warehouse: quantity not int", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(number, 10, 64)
    if err != nil{
        http.Error(w, "error create warehouse: id not int", http.StatusBadRequest)
        return
    }

    err = db.UpdateWarehouse(dbConn, id, quantity)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! warehouse updated!")
}