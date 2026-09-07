package db

import(
    "database/sql"
    "log"
    "os"
    "time"

    _"github.com/lib/pq"
    "ecosystem_mir_j/models"
)
//работа с клиентами
func AddClient(db *sql.DB, name string, email string, phone string) error{
    _, err := db.Exec(`INSERT INTO clients_list (name, email, phone) VALUES ($1, $2, $3)`, name, email, phone)

    if err != nil{
        log.Fatal("error create client: ", err)
    }

    return err
}//добавляем клиета

func GetClient(db *sql.DB, id int64) ([]models.ClientResponse, error){
    rows, err :=db.Query(`SELECT id, name, email, phone, created_at FROM clients_list WHERE id=$1`, id)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.ClientResponse
    for rows.Next(){
        var c models.ClientResponse
        err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получаем клиета

func GetClients(db *sql.DB) ([]models.ClientResponse, error){
    rows, err :=db.Query(`SELECT id, name, email, phone, created_at FROM clients_list`)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.ClientResponse
    for rows.Next(){
        var c models.ClientResponse
        err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone, &c.CreatedAt)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получаем клиетов

//работа с остатками

func AddOst(db *sql.DB, status string, name string, height int64, wight int64) error{
    _, err := db.Exec(`INSERT INTO ost_list (status, name, height, wight) VALUES ($1, $2, $3, $4)`, status, name, height, wight)

    if err != nil{
        log.Fatal("error create slot: ", err)
    }

    return err
}//добавляем остаток

func GetOstsName(db *sql.DB, name string) ([]models.OstResponse, error){
    rows, err :=db.Query(`SELECT id, status, name, height, wight, created_at FROM ost_list WHERE name = $1`, name)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OstResponse
    for rows.Next(){
        var c models.OstResponse
        err = rows.Scan(&c.ID, &c.Status, &c.Name, &c.Height, &c.Wight, &c.CreatedAt)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получение остатков по имени

func GetOSts(db *sql.DB) ([]models.OstResponse, error){
    rows, err :=db.Query(`SELECT id, status, name, height, wight, created_at FROM ost_list`)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OstResponse
    for rows.Next(){
        var c models.OstResponse
        err = rows.Scan(&c.ID, &c.Status, &c.Name, &c.Height, &c.Wight, &c.CreatedAt)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получаем остатки

func GetOSt(db *sql.DB, id int64) ([]models.OstResponse, error){
    rows, err :=db.Query(`SELECT id, status, name, height, wight, created_at FROM ost_list WHERE id=$1`, id)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OstResponse
    for rows.Next(){
        var c models.OstResponse
        err = rows.Scan(&c.ID, &c.Status, &c.Name, &c.Height, &c.Wight, &c.CreatedAt)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получаем остаток

//работа с заказми

func AddOrder(db *sql.DB, title string, description string, measurer string, target_date time.Time, status string, client_id int64) error{
    _, err := db.Exec(`INSERT INTO orders_list (title, description, measurer, target_date, status, client_id)`, title, description, measurer, target_date, status, client_id)

    if err != nil{
        log.Fatal("error create order: ", err)
    }

    return err
}

func TimeOrders(db *sql.DB) ([]models.OrderResponse, error){
    target_date := time.Now().AddDate(0, 0, 3)

    rows, err :=db.Query(`SELECT id, title, description, measurer, created_at, target_date, status, client_id  created_at FROM orders_list WHERE target_date<$1`, target_date)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OrderResponse
    for rows.Next(){
        var c models.OrderResponse
        err = rows.Scan(&c.ID, &c.Title, &c.Descriptoin, &c.CreatedAt, &c.Target_date, &c.Status, &c.Measurer, &c.Client_id)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получение бдижайших заказов

func GetOrdersClient(db *sql.DB, id int64) ([]models.OrderResponse, error){
    rows, err :=db.Query(`SELECT id, title, description, measurer, created_at, target_date, status, client_id  created_at FROM orders_list WHERE client_id=$1`, id)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OrderResponse
    for rows.Next(){
        var c models.OrderResponse
        err = rows.Scan(&c.ID, &c.Title, &c.Descriptoin, &c.CreatedAt, &c.Target_date, &c.Status, &c.Measurer, &c.Client_id)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получение заказов клиента

func GetOrder(db *sql.DB, id int64) ([]models.OrderResponse, error){
    rows, err :=db.Query(`SELECT id, title, description, measurer, created_at, target_date, status, client_id  created_at FROM orders_list WHERE id=$1`, id)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OrderResponse
    for rows.Next(){
        var c models.OrderResponse
        err = rows.Scan(&c.ID, &c.Title, &c.Descriptoin, &c.CreatedAt, &c.Target_date, &c.Status, &c.Measurer, &c.Client_id)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получение заказа

func GetOrders(db *sql.DB) ([]models.OrderResponse, error){
    rows, err :=db.Query(`SELECT id, title, description, measurer, created_at, target_date, status, client_id  created_at FROM orders_list`)

    if err != nil{
        log.Fatal("error select clients: ", err)
    }

    defer rows.Close()

    var clients []models.OrderResponse
    for rows.Next(){
        var c models.OrderResponse
        err = rows.Scan(&c.ID, &c.Title, &c.Descriptoin, &c.CreatedAt, &c.Target_date, &c.Status, &c.Measurer, &c.Client_id)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        clients = append(clients, c)
    }

    return clients, nil
}//получение заказа

//работа с складом комплектации
func AddWarehouse(db *sql.DB, name string, quantity int64, category string, number string) error{
    _, err := db.Exec(`INSERT INTO warehouse_list (name, quantity, category, number) VALUES ($1, $2, $3, $4)`, name, quantity, category, number)

    if err != nil{
        log.Fatal("error create warehouse object: ", err)
    }
    return nil
}

func GetWarehouse(db *sql.DB) ([]models.Warehouse, error){
    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list`)
    if err != nil{
        log.Fatal("error get warehouse: ", err)
    }

    defer rows.Close()

    var warehouse []models.Warehouse
    for rows.Next(){
        var w models.Warehouse
        err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.Category, &w.Number)
        if err != nil{
            log.Fatal("error scen client: ", err)
            continue
        }

        warehouse = append(warehouse, w)
    }

    return warehouse, nil
}//получение комплектации с склада

func PutWarehoyse(db *sql.DB, id int64, category string) error{
    _, err := db.Exec(`UPDATE warehouse_list SET category = $1 WHERE id = $2`, category, id) 

    if err!= nil{
        log.Fatal("error update warehouse: ", err) 
    }
}

func InitDB() *sql.DB{
    //подключаемся к БД
    connstr := os.Getenv("DATABASE_URL")
    if connstr == ""{
        connstr = "postgresql://postgres:36863686@localhost:5432/ecosystem_mir_j?sslmode=disable"
    }

    db, err := sql.Open("postgres", connstr)
    if err != nil{
        log.Fatal("error db open: ", err)
    }

    createTableSQLClients := `
    CREATE TABLE IF NOT EXISTS clients_list (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        phone TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW()
    );`

    _, err = db.Exec(createTableSQLClients)
    if err != nil{
        log.Fatal("error create clients_list table: ", err)
    }

    //добавляем таблицы
    createTableSQLOrders := `
    CREATE TABLE IF NOT EXISTS orders_list (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
        measurer TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW(),
        target_date TIMESTAMP DEFAULT (NOW() + INTERVAL '3 days'),
        status TEXT DEFAULT 'new',
        client_id INT REFERENCES clients_list(id) ON DELETE SET NULL
    );`

    _, err = db.Exec(createTableSQLOrders)
    if err != nil{
        log.Fatal("error create orders_list table: ", err)
    }

    createTableSQLOSt := `
    CREATE TABLE IF NOT EXISTS ost_list (
        id SERIAL PRIMARY KEY,
        status TEXT DEFAULT 'in stock',
        name TEXT NOT NULL,
        height INT NOT NULL,
        wight INT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW()
    );`

    _, err = db.Exec(createTableSQLOSt)
    if err != nil{
        log.Fatal("error create ost_list table: ", err)
    }

    createTableSQLWarehouse := `
    CREATE TABLE IF NOT EXISTS warehouse_list(
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        quantity INT NOT NULL,
        category TEXT NOT NULL,
        number TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW()
    );`

    _, err = db.Exec(createTableSQLWarehouse)
    if err != nil{
        log.Fatal("error create warehouse_list table: ", err)
    }

    return db
}
