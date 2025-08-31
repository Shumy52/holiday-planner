package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "yourpassword"
    dbname   = "vacation_planner"
)

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }
    fmt.Println("Connected to database!")

    _, err = db.Exec(`
        INSERT INTO users (full_name, email, role, department)
        VALUES
        ('Ion Popescu', 'ion.popescu@example.com', 'employee', 'IT'),
        ('Maria Ionescu', 'maria.ionescu@example.com', 'manager', 'HR'),
        ('Andrei Georgescu', 'andrei.georgescu@example.com', 'admin', 'Operations')
    `)
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec(`
        INSERT INTO vacation_requests (user_id, start_date, end_date, status)
        VALUES
        (1, '2025-08-01', '2025-08-10', 'approved'),
        (1, '2025-09-15', '2025-09-20', 'pending')
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Seed Approvals
    _, err = db.Exec(`
        INSERT INTO approvals (request_id, manager_id, decision)
        VALUES
        (1, 2, 'approved'),
        (2, 2, 'rejected')
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Seed Audit Logs
    _, err = db.Exec(`
        INSERT INTO audit_log (user_id, action, details)
        VALUES
        (1, 'Created request', 'Request for 01.08.2025 - 10.08.2025'),
        (2, 'Approved request', 'Approved request ID 1 for Ion Popescu')
    `)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database seeded successfully!")
}
