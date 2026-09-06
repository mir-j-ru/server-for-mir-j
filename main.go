package main

import(
    "database/sql"
    "fmt"
    "log"
    "net/http"      // ✅ Правильно
    "os"

    "ecosystem_mir_j/db"
    "ecosystem_mir_j/handlers"

    "github.com/joho/godotenv"
)

var dbConn *sql.DB

func main(){
    //загружаем окружение
    err := godotenv.Load()
    if err != nil{
        log.Fatal("error load godotenv: ", err)
    }

    dbConn := db.InitDB()
    defer dbConn.Close()

    http.HandleFunc("/get_clients", func(w http.ResponseWriter, r *http.Request){
        handlers.GetClients(w, r, dbConn)
    })//получение всех клиентов

    http.HandleFunc("/get_client", func(w http.ResponseWriter, r *http.Request){
        handlers.GetClient(w, r, dbConn)
    })//получение одного клиента

    http.HandleFunc("/add_client", func(w http.ResponseWriter, r *http.Request){
        handlers.AddClient(w, r, dbConn)
    })//добавлние клиента

    //http.HandleFunc("/put_client", func(w http.ResponseWriter, r *http.Request){
        //handlers.PutClient(w, r, dbConn)
    //})//обновление клиента

    http.HandleFunc("/get_osts", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOsts(w, r, dbConn)
    })//получение всех остатков

    http.HandleFunc("/get_osts_name", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOstsName(w, r, dbConn)
    })//получение всех остатков по названию ткани

    http.HandleFunc("/get_ost", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOst(w, r, dbConn)
    })//получение одного остатка

    http.HandleFunc("/add_ost", func(w http.ResponseWriter, r *http.Request){
        handlers.AddOst(w, r, dbConn)
    })//добовление остатка

    //http.HandleFunc("/put_ost", func(w http.ResponseWriter, r *http.Request){
        //handlers.PutOst(w, r, dbConn)
    //})//обновление остатка

    http.HandleFunc("/get_orders", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOrders(w, r, dbConn)
    })//получение всех заказов

    http.HandleFunc("/get_order", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOrder(w, r, dbConn)
    })//получение одного заказа

    http.HandleFunc("/time_order", func(w http.ResponseWriter, r *http.Request){
        handlers.TimeOrders(w, r, dbConn)
    })//получение всех заказов на ближайшее время

    http.HandleFunc("/add_order", func(w http.ResponseWriter, r *http.Request){
        handlers.AddOrder(w, r, dbConn)
    })//добавлние заказа

    //http.HandleFunc("/put_order", func(w http.ResponseWriter, r *http.Request){
        //handlers.PutOrder(w, r, dbConn)
    //})//обновление заказа

    http.HandleFunc("/client_orders", func(w http.ResponseWriter, r *http.Request){
        handlers.GetOrdersClient(w, r, dbConn)
    })//получение всех заказов одного клиета

    http.HandleFunc("/add_warehouse", func(w http.ResponseWriter, r *http.Request){
        handlers.AddWarehouse(w, r, dbConn)
    })//добавление комплектации

    http.HandleFunc("/get_waregouse", func(w http.ResponseWriter, r *http.Request){
        handlers.GetWarehouse(w, r, dbConn)
    })//получение комплектации

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Сервер запущен на порту %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}