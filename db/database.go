package database


import(
    //"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


type db struct {
    db *sql.DB
    connected bool
    name string
}

func (d *db) setConnectedState (state bool){
    d.connected = state
}

func (d db) GetConnState () bool{
    return d.connected
}

func ConnectDB (name string) (*db, error) {
    dbase, err  := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/osousa")
    if err != nil {
        panic("Unable to connect to db")
    }
    db := db{ db: dbase, connected: false, name: "osousa"}

    db.setConnectedState(true)

    return &db, nil
}