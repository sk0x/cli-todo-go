package helpers

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// global db name
const DATABASE_NAME string = "todo.db"

func openDatabase()(*sql.DB, error){
    home_path, path_err := os.UserHomeDir()
    if path_err != nil {
        return nil, path_err
    }
    db_path := fmt.Sprintf("%s/%s", home_path, DATABASE_NAME)
    db, db_err := sql.Open("sqlite3", db_path)
    if db_err != nil {
        return nil, db_err
    }
    return db, nil
}

func CreateTodosTable(){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
    }
    defer db.Close()

    sqlQuery := `CREATE TABLE IF NOT EXISTS todos
    (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    todo TEXT,
    status BIT NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`

    _, err := db.Exec(sqlQuery)
    if err != nil {
        fmt.Println("Error: While create todos table")
        os.Exit(1)
        return
    }
}

func AddTask(task string){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer db.Close()
    _, ex_err := db.Exec("INSERT INTO todos(todo) VALUES(?)", task)
    if ex_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    fmt.Println("Task added successfully")
}

func UpdateTask(id int){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer db.Close()

    _, resErr := db.Exec("UPDATE todos SET status = 1 WHERE id = ?", id)
    if resErr != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    fmt.Println("Updated todo successfully")
    return
}

func DeleteTask(id int){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer db.Close()

    _, resErr := db.Exec("DELETE FROM todos WHERE id=?", id)
    if resErr != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    fmt.Println("Deleted todo successfully")
    return
}


func ListTasks(){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer db.Close()
    rows, rerr := db.Query("SELECT id, todo, status, created_at FROM todos")
    if rerr != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer rows.Close()

    // creating new tab writer
    tabw := tabwriter.NewWriter(os.Stdout, 4, 10, 4, ' ', 10)
    // writing column headers
    fmt.Fprintf(tabw, "%s\t%s\t%s\t%s\n", "ID", "TODOS", "STATUS", "CREATED")

    for rows.Next(){
        var id, status int
        var todo, created_at string
        err := rows.Scan(&id, &todo, &status, &created_at)
        if err != nil {
            fmt.Println(db_err)
            os.Exit(1)
            return
        }
        var isCompleted string
        if status == 0 {
            isCompleted = "pending"
        } else {
            isCompleted = "complete"
        }
        fmt.Fprintf(tabw, "%d\t%s\t%s\t%s\n", id, todo, isCompleted, created_at)
    }
    tabw.Flush()
}

func RemoveLastDayCompletedTasks(){
    db, db_err := openDatabase()
    if db_err != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    defer db.Close()
    loc, _ := time.LoadLocation("Asia/Kolkata")
    now := time.Now().In(loc)
    midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
    result, rerr := db.Exec("DELETE FROM todos WHERE created_at < ? AND status=1", midnight)
    if rerr != nil {
        fmt.Println(db_err)
        os.Exit(1)
        return
    }
    afrows, _ := result.RowsAffected()
    if afrows > 0{
        fmt.Println("Deleted yesterdays completed todo(s)")
    }
    return
}
