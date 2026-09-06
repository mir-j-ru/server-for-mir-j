package handlers

import(
    "database/sql"
    "encoding/json"
    //"fmt"
    //"log"
    "net/http"
    //"os"

    //"studia-of-biautiful-api/models" // или "weather-go-api/models"
    "ecosystem_mir_j/db"     // или "weather-go-api/db"
)

func GetWarehouse(w http.ResponseWriter, r *http.Request, dbConn *sql.DB){
    quotes, err := db.GetWarehouse(dbConn)
    if err!=nil{
        http.Error(w, "error db:", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quotes)
}
