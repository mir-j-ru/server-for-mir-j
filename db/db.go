package db

import(
    "database/sql"
    "log"
    "os"
    "time"
    "strconv"
    "fmt"
    "math"

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

func UpdateWarehouse(db *sql.DB, id int64, amount int64) error{
    _, err := db.Exec(`UPDATE warehouse_list SET quantity = quantity + $1 WHERE id = $2`, amount, id)

    if err != nil{
        log.Fatal("error updated warehouse: ", err)
    }

    return err
}


func GetWarehouse(db *sql.DB) ([]models.Warehouse, error){
    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list ORDER BY id`)
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
    return nil
}

func DeleteWarehouse(db *sql.DB, id int64) error{
    _, err := db.Exec(`DELETE FROM warehouse_list WHERE id=$1`, id)

    if err!=nil{
        log.Fatal("error delete warehouse_list: ", err)
    }

    return err
}

func GorizontDefault(db *sql.DB, height int64, width int64) error {
    // height и width не используются в этом коде — если они нужны позже, оставь.
    orders := map[string]int{
        "67": -1, "69": -1, "72": -1, "73": -1,
        "76": -4, "77": -2, "78": -1, "79": -1,
        "80": -1, "84": -2, "85": -2, "87": -2,
        "92": -2,
    }

    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list`)
    if err != nil {
        return fmt.Errorf("error get warehouse: %w", err)
    }
    defer rows.Close()

    var warehouse []models.Warehouse
    for rows.Next() {
        var w models.Warehouse
        err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.Category, &w.Number)
        if err != nil {
            return fmt.Errorf("error scan warehouse row: %w", err)
        }
        warehouse = append(warehouse, w)
    }
    if rows.Err() != nil {
        return fmt.Errorf("error after rows iteration: %w", rows.Err())
    }

    // 1. Проверка достаточности товара
    // Предполагаем, что ключ orders — это ID склада (как строка).
    for idStr, amount := range orders {
        for _, w := range warehouse {
            _, err := strconv.ParseInt(idStr, 10, 64)
            if err != nil {
                return fmt.Errorf("error parsing id to int: %w", err)
            }

            if w.Quantity+amount < 0 {
                return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
            }

            switch w.ID {
            case 62, 63:
                if int64(w.Quantity)-height < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 83:
                if int64(w.Quantity)-width-150 < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 397:
                if w.Quantity-800 < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            }
        }
    }

    // 2. Обновление остатков
    for idStr, amount := range orders {
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            return fmt.Errorf("error parsing id to int: %w", err)
        }

        _, err = db.Exec(`UPDATE warehouse_list SET quantity = quantity + $1 WHERE id = $2`, amount, id)
        if err != nil {
            return fmt.Errorf("error updating warehouse: %w", err)
        }
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=62`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 62: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=63`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 63: %w", err)
    }

    if width > 1500{
        _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-1000 WHERE id=397`)
        if err != nil{
            return fmt.Errorf("error update warehouse 397: %w", err)
        }
    }else{
        _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-800 WHERE id=397`)
        if err != nil{
            return fmt.Errorf("error update warehouse 397: %w", err)
        }
    }

    wifht_res := width+150

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=83`, wifht_res)
    if err != nil{
        return fmt.Errorf("error update warehouse 83: %w", err)
    }

    return nil
}

func VerDef(db *sql.DB, height int64, width int64) error{
    orders := map[string]int{"7" : 1, "8" : 1,"10":1,"11":1,"15":1,"17":1,"25":2,"35":1,"38":1,"43":1,"47":1,"48":1,}

    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list`)
    if err != nil {
        return fmt.Errorf("error get warehouse: %w", err)
    }
    defer rows.Close()

    var warehouse []models.Warehouse
    for rows.Next() {
        var w models.Warehouse
        err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.Category, &w.Number)
        if err != nil {
            return fmt.Errorf("error scan warehouse row: %w", err)
        }
        warehouse = append(warehouse, w)
    }
    if rows.Err() != nil {
        return fmt.Errorf("error after rows iteration: %w", rows.Err())
    }

    for idStr, amount := range orders {
        for _, w := range warehouse {
            _, err := strconv.ParseInt(idStr, 10, 64)
            if err != nil {
                return fmt.Errorf("error parsing id to int: %w", err)
            }

            if w.Quantity+amount < 0 {
                return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
            }

            switch w.ID {   // ← теперь внутри for
            case 6, 398:
                if int64(w.Quantity)-height < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 13, 17, 20:
                if int64(w.Quantity)-height/9 < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 24:
                if int64(w.Quantity)-int64(math.Round(float64(height)/9*2)) < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 23:
                if int64(w.Quantity)-int64(math.Round(float64(height)*2+(float64(width)-20)*2)) < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            }
        }
    }

    // 2. Обновление остатков
    for idStr, amount := range orders {
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            return fmt.Errorf("error parsing id to int: %w", err)
        }

        _, err = db.Exec(`UPDATE warehouse_list SET quantity = quantity + $1 WHERE id = $2`, amount, id)
        if err != nil {
            return fmt.Errorf("error updating warehouse: %w", err)
        }
    }    
    
    minus_d :=  math.Round(float64(height/9))

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=13`, minus_d)
    if err != nil{
        return fmt.Errorf("error update warehouse 13: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=6`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 6: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=398`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 398: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=17`, minus_d)
    if err != nil{
        return fmt.Errorf("error update warehouse 17: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=20`, minus_d)
    if err != nil{
        return fmt.Errorf("error update warehouse 20: %w", err)
    }

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=24`, minus_d*2)
    if err != nil{
        return fmt.Errorf("error update warehouse 24: %w", err)
    }

    for_dt := height*2+(width-20)*2

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=23`, for_dt)
    if err != nil{
        return fmt.Errorf("error update warehouse 23: %w", err)
    }

    for_tr := (width-20)*2

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=30`, for_tr)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    }

    return nil
}

func DefMINI(db *sql.DB, height int64, width int64) error{
    orders := map[string]int{"250":1,"251":2,"256":1,"259":2,"287":2,"244":1}

    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list`)
    if err != nil {
        return fmt.Errorf("error get warehouse: %w", err)
    }
    defer rows.Close()

    var warehouse []models.Warehouse
    for rows.Next() {
        var w models.Warehouse
        err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.Category, &w.Number)
        if err != nil {
            return fmt.Errorf("error scan warehouse row: %w", err)
        }
        warehouse = append(warehouse, w)
    }
    if rows.Err() != nil {
        return fmt.Errorf("error after rows iteration: %w", rows.Err())
    }

    for idStr, amount := range orders {
        for _, w := range warehouse {
            _, err := strconv.ParseInt(idStr, 10, 64)
            if err != nil {
                return fmt.Errorf("error parsing id to int: %w", err)
            }

            if w.Quantity+amount < 0 {
                return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
            }

            switch w.ID {   // ← теперь внутри for
            case 235, 236,316,314:
                if int64(w.Quantity)-height < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            case 242:
                if int64(w.Quantity)-int64(math.Round(float64(wight*2-30))) < 0 {
                    return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
                }
            }
        }
    }

    // 2. Обновление остатков
    for idStr, amount := range orders {
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            return fmt.Errorf("error parsing id to int: %w", err)
        }

        _, err = db.Exec(`UPDATE warehouse_list SET quantity = quantity + $1 WHERE id = $2`, amount, id)
        if err != nil {
            return fmt.Errorf("error updating warehouse: %w", err)
        }
    }  

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=235`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    }         

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=236`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    }  

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=316`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    } 

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=314`, height)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    } 

    _, err = db.Exec(`UPDATE warehouse_list SET quantity=quantity-$1 WHERE id=242`, wight*2-30)
    if err != nil{
        return fmt.Errorf("error update warehouse 30: %w", err)
    } 

    return nil
}

func DefUNI(db *sql.DB, height int64, width int64) error{
    orders := map[string]int{}

    rows, err := db.Query(`SELECT id, name, quantity, category, number FROM warehouse_list`)
    if err != nil {
        return fmt.Errorf("error get warehouse: %w", err)
    }
    defer rows.Close()

    var warehouse []models.Warehouse
    for rows.Next() {
        var w models.Warehouse
        err = rows.Scan(&w.ID, &w.Name, &w.Quantity, &w.Category, &w.Number)
        if err != nil {
            return fmt.Errorf("error scan warehouse row: %w", err)
        }
        warehouse = append(warehouse, w)
    }
    if rows.Err() != nil {
        return fmt.Errorf("error after rows iteration: %w", rows.Err())
    }

    for idStr, amount := range orders {
        for _, w := range warehouse {
            _, err := strconv.ParseInt(idStr, 10, 64)
            if err != nil {
                return fmt.Errorf("error parsing id to int: %w", err)
            }

            if w.Quantity+amount < 0 {
                return fmt.Errorf("недостаточно товара на складе %s (id=%d)", w.Name, w.ID)
            }
        }
    }

    // 2. Обновление остатков
    for idStr, amount := range orders {
        id, err := strconv.ParseInt(idStr, 10, 64)
        if err != nil {
            return fmt.Errorf("error parsing id to int: %w", err)
        }

        _, err = db.Exec(`UPDATE warehouse_list SET quantity = quantity + $1 WHERE id = $2`, amount, id)
        if err != nil {
            return fmt.Errorf("error updating warehouse: %w", err)
        }
    }  

    return nil
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
