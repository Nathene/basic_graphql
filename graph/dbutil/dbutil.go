package dbutil

import (
	"database/sql"
	"fmt"
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
	GetProjectsByStaffID(staffID string) ([]*model.Project, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

// DB is a wrapper for the SQLite database connection.
type DB struct {
	*sql.DB
}

// NewDB initializes a new in-memory SQLite database.
func NewDB() *DB {
	conn, err := sql.Open("sqlite", "graphql.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &DB{conn}
}

// CreateStaff inserts a new staff record into the database.
func (db *DB) CreateStaff(staff *model.Staff) error {
	_, err := db.Exec(
		"INSERT INTO staff (id, first_name, last_name, email, role, department, salary) VALUES (?, ?, ?, ?, ?, ?, ?)",
		staff.ID, staff.FirstName, staff.LastName, staff.Email, staff.Role, staff.Department, staff.Salary,
	)
	return err
}

// GetStaffByID retrieves a staff record by its ID.
func (db *DB) GetStaffByID(id string) (*model.Staff, error) {
	row := db.QueryRow("SELECT id, first_name, last_name, email, role, department, salary FROM staff WHERE id = ?", id)
	staff := model.Staff{}
	err := row.Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Email, &staff.Role, &staff.Department, &staff.Salary)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &staff, err
}

// GetAllStaff retrieves all staff records.
func (db *DB) GetAllStaff() ([]*model.Staff, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, role, department, salary FROM staff")
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

func (db *DB) GetProjectsByStaffID(staffID string) ([]*model.Project, error) {
	query := `SELECT p.id, p.name, p.description 
	          FROM projects p 
	          JOIN staff_projects sp ON sp.project_id = p.id 
	          WHERE sp.staff_id = ?`

	rows, err := db.Query(query, staffID)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %v", err)
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}

// FetchProjectsForStaff retrieves all projects for a given staff member
func (db *DB) FetchProjectsForStaff(staffID int) ([]*model.Project, error) {
	rows, err := db.Query(`
        SELECT p.id, p.name, p.description, p.department, p.budget
        FROM projects p
        INNER JOIN staff_projects sp ON p.id = sp.project_id
        WHERE sp.staff_id = ?
    `, staffID)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []*model.Project
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Department, &project.Budget); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}
