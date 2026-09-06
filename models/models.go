package models

import(
 "time"
)

type ClientResponse struct{
    ID int `json: "ID"`
    Name string `json: "name"`
    Email string `json: "email"`
    Phone string `json: "phone"`
    CreatedAt time.Time `json: "created_at"`
}

type OrderResponse struct{
    ID int `json: "ID"`
    Title string `json: "title"`
    Descriptoin string `json: "description"`
    CreatedAt time.Time `json: "created_at"`
    Target_date time.Time `json: "target_date"`
    Status string `json: "status"`
    Measurer string `json: "measurer"`
    Client_id int `json: "client_id"`
}

type OstResponse struct{
    ID int `json: "ID"`
    Status string `json: "status"`
    Name string `json: "name"`
    Height int `json: "height"`
    Wight int `json: "wight"`
    CreatedAt time.Time `json: "created_at"`
}

type Warehouse struct{
    ID int `json: "id"`
    Name string `json: "name"`
    Quantity int `json: "quantity"`
    Category string `json: "category"`
    Number string `json: "number"`
}
