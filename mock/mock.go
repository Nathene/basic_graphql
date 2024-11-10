package mock

import (
	"database/sql"
	"log"
)

// Staff represents a staff member with basic fields
type Staff struct {
	FirstName  string
	LastName   string
	Email      string
	Role       string
	Department string
	Salary     float64
}

// Project represents a project with basic fields
type Project struct {
	Name        string
	Description string
	Department  string
	Budget      float64
}

// Mock data to insert
var mockStaffData = []Staff{
	{"John", "Doe", "johndoe@example.com", "Software Engineer", "Engineering", 65000.00},
	{"Jane", "Smith", "janesmith@example.com", "Designer", "Design", 55000.00},
	{"Alice", "Brown", "alicebrown@example.com", "Product Manager", "Product", 75000.00},
	{"Bob", "Johnson", "bobjohnson@example.com", "Data Analyst", "Data", 60000.00},
	{"Charlie", "Davis", "charliedavis@example.com", "Engineer", "Engineering", 68000.00},
}

var mockProjectData = []Project{
	{"Project Apollo", "A lunar landing project", "Engineering", 1500000.00},
	{"Project Hermes", "A high-speed transportation project", "Product", 3000000.00},
	{"Project Athena", "Data analysis automation project", "Data", 500000.00},
	{"Project Zeus", "Design overhaul for main website", "Design", 750000.00},
	{"Project Poseidon", "Ocean exploration initiative", "Research", 2500000.00},
}

// createTables creates the staff and projects tables if they do not already exist
func createTables(db *sql.DB) {
	staffTableStmt := `
    CREATE TABLE IF NOT EXISTS staff (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        first_name TEXT,
        last_name TEXT,
        email TEXT UNIQUE,
        role TEXT,
        department TEXT,
        salary REAL
    );
    `

	projectTableStmt := `
    CREATE TABLE IF NOT EXISTS projects (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        description TEXT,
        department TEXT,
        budget REAL
    );
    `

	_, err := db.Exec(staffTableStmt)
	if err != nil {
		log.Fatalf("Failed to create staff table: %v", err)
	}

	_, err = db.Exec(projectTableStmt)
	if err != nil {
		log.Fatalf("Failed to create projects table: %v", err)
	}
}

// insertMockData inserts batches of staff and project data into the database
func InsertMockData(db *sql.DB) error {
	createTables(db)
	staffInsertStmt := `INSERT INTO staff (first_name, last_name, email, role, department, salary) VALUES (?, ?, ?, ?, ?, ?)`
	projectInsertStmt := `INSERT INTO projects (name, description, department, budget) VALUES (?, ?, ?, ?)`

	// Insert staff data
	for _, staff := range mockStaffData {
		_, err := db.Exec(staffInsertStmt, staff.FirstName, staff.LastName, staff.Email, staff.Role, staff.Department, staff.Salary)
		if err != nil {
			// Skip duplicate entries but log the error
			if sqliteErr, ok := err.(interface{ Code() int }); ok && sqliteErr.Code() == 2067 { // SQLITE_CONSTRAINT_UNIQUE
				log.Printf("Skipping duplicate entry for email: %s", staff.Email)
				continue
			}
			return err // Return error if it's not a duplicate
		}
	}

	// Insert project data
	for _, project := range mockProjectData {
		_, err := db.Exec(projectInsertStmt, project.Name, project.Description, project.Department, project.Budget)
		if err != nil {
			log.Printf("Failed to insert project %s: %v", project.Name, err)
			return err
		}
	}

	return nil
}
