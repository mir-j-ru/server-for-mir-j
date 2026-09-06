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

func AddWarehouse(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    value := r.URL.Query()

    name := value.Get("name")
    category := value.Get("category")
    number := value.Get("number")
    quantitystr := value.Get("quantity")

    if name == ""{
        http.Error(w, "error create warehouse: not name", http.StatusBadRequest)
        return
    }

    if category == ""{
        http.Error(w, "error create warehouse: not category", http.StatusBadRequest)
        return
    }

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

    err = db.AddWarehouse(dbConn, name, quantity, category, number)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "success! warehouse added!")
}
