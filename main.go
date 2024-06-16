package main

import (
    "database/sql"
    "flag"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

type Barang struct {
    Name  string `json:"name"`
}

type Login struct {
    Emp_no  string `json:"emp_no"`
}

func main() {
    port := flag.String("port", "8080", "port to run the server on")
    flag.Parse()
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ale_project")
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()
    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging database:", err)
    }
    r := gin.Default()

    r.GET("/barang", func(c *gin.Context) {
        var barang []Barang

        rows, err := db.Query("SELECT name FROM tbl_barang")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        for rows.Next() {
            var item Barang
            if err := rows.Scan(&item.Name); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            barang = append(barang, item)
        }

        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, barang)
    })

	r.GET("/login", func(c *gin.Context) {
        var login []Login

        rows, err := db.Query("SELECT emp_no FROM tbl_login")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        for rows.Next() {
            var emp_no Login
            if err := rows.Scan(&emp_no.Emp_no); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            login = append(login, emp_no)
        }

        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, login)
    })

    addr := fmt.Sprintf(":%s", *port)
    log.Printf("Server running on %s", addr)
    r.Run(addr)
}
