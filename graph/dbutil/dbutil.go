package dbutil

import (
	"database/sql"
	"graphql/graph/model"
	"log"

	_ "modernc.org/sqlite"
)

// Database defines the methods that the application uses to access staff data.
type Database interface {
	CreateStaff(staff *model.Staff) error
	GetStaffByID(id string) (*model.Staff, error)
	GetAllStaff() ([]*model.Staff, error)
	AssignProjectToStaff(staffID, projectID string) (*model.Project, error)
}

// DB is a wrapper for the SQLite database connection.
type DB struct {
	Conn *sql.DB
}

// NewDB initializes a new in-memory SQLite database.
func NewDB() *DB {
	conn, err := sql.Open("sqlite", "graphql.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &DB{Conn: conn}
}

// CreateStaff inserts a new staff record into the database.
func (db *DB) CreateStaff(staff *model.Staff) error {
	_, err := db.Conn.Exec(
		"INSERT INTO staff (id, first_name, last_name, email, role, department, salary) VALUES (?, ?, ?, ?, ?, ?, ?)",
		staff.ID, staff.FirstName, staff.LastName, staff.Email, staff.Role, staff.Department, staff.Salary,
	)
	return err
}

// GetStaffByID retrieves a staff record by its ID.
func (db *DB) GetStaffByID(id string) (*model.Staff, error) {
	row := db.Conn.QueryRow("SELECT id, first_name, last_name, email, role, department, salary FROM staff WHERE id = ?", id)
	staff := model.Staff{}
	err := row.Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Email, &staff.Role, &staff.Department, &staff.Salary)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &staff, err
}

// GetAllStaff retrieves all staff records.
func (db *DB) GetAllStaff() ([]*model.Staff, error) {
	rows, err := db.Conn.Query("SELECT id, first_name, last_name, email, role, department, salary FROM staff")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffList []*model.Staff
	for rows.Next() {
		staff := model.Staff{}
		err := rows.Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Email, &staff.Role, &staff.Department, &staff.Salary)
		if err != nil {
			return nil, err
		}
		staffList = append(staffList, &staff)
	}
	return staffList, nil
}

func (db *DB) AssignProjectToStaff(staffID, projectID string) (*model.Project, error) {
	return nil, nil
}
