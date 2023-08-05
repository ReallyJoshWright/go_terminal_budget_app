package main

import (
    "fmt"
    "flag"
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func parse_args(file string, initial_balance float64, transactions int, balance bool, credits bool, debits bool, spending bool, flags int) {
    if flags == 0 {
        init_balance := get_initial_balance()
        bal := get_balance()
        fmt.Println("Initial Balance:", init_balance)
        fmt.Println("Balance:", bal)
    }
}

func create_connection() {
    db, err := sql.Open("sqlite3", "config/account.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    query := `
    SELECT *
    FROM account;
    `
    result, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
    for result.Next() {
        var id int
        result.Scan(&id)
        fmt.Println("id:", id)
    }
}

func get_initial_balance() float64 {
    return 0
}

func get_balance() float64 {
    return 0
}

func main() {
    var file string
    var initial_balance float64
    var transactions int
    var balance bool
    var credits bool
    var debits bool
    var spending bool
    var flags int

    flag.StringVar(&file, "f", "none", "Enter a file")
    flag.Float64Var(&initial_balance, "B", 0, "Enter an initial balance")
    flag.IntVar(&transactions, "n", 10, "Enter number of transactions")
    flag.BoolVar(&balance, "b", false, "Returns the balance")
    flag.BoolVar(&credits, "c", false, "Returns only credits")
    flag.BoolVar(&debits, "d", false, "Returns only debits")
    flag.BoolVar(&spending, "g", false, "Returns spending info")

    flag.Parse()
    flags = flag.NFlag()

    create_connection()
    parse_args(file, initial_balance, transactions, balance, credits, debits, spending, flags)
}
