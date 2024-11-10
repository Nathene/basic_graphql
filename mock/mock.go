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

// Sample mock data for staff and projects
var mockStaffData = []Staff{
	{"John", "Doe", "johndoe@example.com", "Software Engineer", "Engineering", 65000.00},
	{"Jane", "Smith", "janesmith@example.com", "Product Designer", "Design", 55000.00},
	{"Alice", "Brown", "alicebrown@example.com", "Product Manager", "Project Management", 75000.00},
	{"Bob", "Johnson", "bobjohnson@example.com", "Data Analyst", "Data", 60000.00},
	{"Charlie", "Davis", "charliedavis@example.com", "Site Reliability Engineer", "SRE", 68000.00},
	{"David", "Miller", "davidmiller@example.com", "HR Manager", "HR", 80000.00},
	{"Eva", "Martinez", "evamartinez@example.com", "UX Designer", "Design", 72000.00},
	{"Frank", "Wilson", "frankwilson@example.com", "Software Engineer", "Engineering", 95000.00},
	{"Grace", "Taylor", "gracetaylor@example.com", "Project Coordinator", "Project Management", 53000.00},
	{"Henry", "Moore", "henrymoore@example.com", "Full Stack Developer", "Engineering", 85000.00},
	{"Isabelle", "Clark", "isabelleclark@example.com", "Data Scientist", "Data", 93000.00},
	{"Jack", "Lewis", "jacklewis@example.com", "Software Engineer", "SRE", 78000.00},
	{"Katherine", "Walker", "katherinewalker@example.com", "Product Owner", "Product", 67000.00},
	{"Leo", "Harris", "leoharris@example.com", "Business Analyst", "Project Management", 58000.00},
	{"Mona", "Young", "monayoung@example.com", "Data Architect", "Data", 105000.00},
	{"Nathan", "Scott", "nathanscott@example.com", "Marketing Manager", "Marketing", 54000.00},
	{"Olivia", "King", "oliviaking@example.com", "Cloud Architect", "Engineering", 110000.00},
	{"Peter", "Wright", "peterwright@example.com", "Web Developer", "Engineering", 62000.00},
	{"Quincy", "Adams", "quincyadams@example.com", "IT Support", "IT", 45000.00},
}

var mockProjectData = []Project{
	{
		Name:        "Project Apollo",
		Description: "A lunar landing project",
		Department:  "Engineering",
		Budget:      1500000.00,
	},
	{
		Name:        "Project Hermes",
		Description: "A high-speed transportation project",
		Department:  "Product",
		Budget:      3000000.00,
	},
	{
		Name:        "Project Athena",
		Description: "Data analysis automation project",
		Department:  "Data",
		Budget:      500000.00,
	},
	{
		Name:        "Project Zeus",
		Description: "Design overhaul for main website",
		Department:  "Design",
		Budget:      750000.00,
	},
	{
		Name:        "Project Poseidon",
		Description: "Ocean exploration initiative",
		Department:  "Research",
		Budget:      2500000.00,
	},
	{
		Name:        "Project Orion",
		Description: "AI-driven customer service platform",
		Department:  "Engineering",
		Budget:      1200000.00,
	},
	{
		Name:        "Project Atlas",
		Description: "Global expansion logistics",
		Department:  "Product",
		Budget:      2200000.00,
	},
	{
		Name:        "Project Titan",
		Description: "Quantum computing research",
		Department:  "Research",
		Budget:      3000000.00,
	},
	{
		Name:        "Project Gemini",
		Description: "Social media platform redesign",
		Department:  "Design",
		Budget:      850000.00,
	},
	{
		Name:        "Project Apollo II",
		Description: "Lunar base development",
		Department:  "Engineering",
		Budget:      2500000.00,
	},
	{
		Name:        "Project Nimbus",
		Description: "Cloud-native infrastructure",
		Department:  "Engineering",
		Budget:      1700000.00,
	},
	{
		Name:        "Project Elysium",
		Description: "Advanced cybersecurity solutions",
		Department:  "SRE",
		Budget:      500000.00,
	},
	{
		Name:        "Project Nexus",
		Description: "Augmented reality marketing campaign",
		Department:  "Marketing",
		Budget:      1300000.00,
	},
	{
		Name:        "Project Lynx",
		Description: "AI-based fraud detection system",
		Department:  "Data",
		Budget:      1000000.00,
	},
	{
		Name:        "Project Spartan",
		Description: "Mobile application development",
		Department:  "Product",
		Budget:      800000.00,
	},
	{
		Name:        "Project Aurora",
		Description: "Smart city project",
		Department:  "Research",
		Budget:      3500000.00,
	},
	{
		Name:        "Project Vortex",
		Description: "Autonomous vehicle development",
		Department:  "Engineering",
		Budget:      4200000.00,
	},
	{
		Name:        "Project Cyclone",
		Description: "Global logistics platform",
		Department:  "Product",
		Budget:      2300000.00,
	},
	{
		Name:        "Project Eclipse",
		Description: "Privacy-focused search engine",
		Department:  "Engineering",
		Budget:      1400000.00,
	},
}

// Define relationships between staff and projects
var mockStaffProjects = map[string][]string{
	"johndoe@example.com":         {"Project Apollo", "Project Zeus", "Project Orion"},
	"janesmith@example.com":       {"Project Hermes", "Project Atlas"},
	"alicebrown@example.com":      {"Project Apollo", "Project Athena", "Project Gemini"},
	"bobjohnson@example.com":      {"Project Athena", "Project Titan"},
	"charliedavis@example.com":    {"Project Poseidon", "Project Elysium"},
	"davidmiller@example.com":     {"Project Orion", "Project Titan"},
	"evamartinez@example.com":     {"Project Apollo", "Project Atlas", "Project Gemini"},
	"frankwilson@example.com":     {"Project Orion", "Project Vortex"},
	"gracetaylor@example.com":     {"Project Apollo", "Project Gemini"},
	"henrymoore@example.com":      {"Project Elysium", "Project Vortex"},
	"isabelleclark@example.com":   {"Project Titan", "Project Nexus"},
	"jacklewis@example.com":       {"Project Gemini", "Project Cyclone"},
	"katherinewalker@example.com": {"Project Atlas", "Project Vortex"},
	"leoharris@example.com":       {"Project Lynx", "Project Titan"},
	"monayoung@example.com":       {"Project Apollo", "Project Vortex"},
	"nathanscott@example.com":     {"Project Atlas", "Project Eclipse"},
	"oliviaking@example.com":      {"Project Nimbus", "Project Eclipse"},
	"peterwright@example.com":     {"Project Lynx", "Project Aurora"},
	"quincyadams@example.com":     {"Project Cyclone", "Project Apollo"},
}

// createTables creates the staff, projects, and staff_projects tables if they do not already exist
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

	staffProjectsTableStmt := `
	CREATE TABLE IF NOT EXISTS staff_projects (
    staff_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    FOREIGN KEY (staff_id) REFERENCES staff(id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    PRIMARY KEY (staff_id, project_id)
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

	_, err = db.Exec(staffProjectsTableStmt)
	if err != nil {
		log.Fatalf("Failed to create staff_projects table: %v", err)
	}
}

// InsertMockData inserts batches of staff, project, and staff-project relationship data into the database
func InsertMockData(db *sql.DB) error {
	createTables(db)

	staffInsertStmt := `INSERT INTO staff (first_name, last_name, email, role, department, salary) VALUES (?, ?, ?, ?, ?, ?)`
	projectInsertStmt := `INSERT INTO projects (name, description, department, budget) VALUES (?, ?, ?, ?)`
	staffProjectInsertStmt := `INSERT INTO staff_projects (staff_id, project_id) VALUES (?, ?)`

	// Insert staff data
	for _, staff := range mockStaffData {
		_, err := db.Exec(staffInsertStmt, staff.FirstName, staff.LastName, staff.Email, staff.Role, staff.Department, staff.Salary)
		if err != nil {
			if sqliteErr, ok := err.(interface{ Code() int }); ok && sqliteErr.Code() == 2067 { // SQLITE_CONSTRAINT_UNIQUE
				log.Printf("Skipping duplicate entry for email: %s", staff.Email)
				continue
			}
			return err
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

	// Insert staff-project relationships
	for email, projects := range mockStaffProjects {
		var staffID int
		err := db.QueryRow("SELECT id FROM staff WHERE email = ?", email).Scan(&staffID)
		if err != nil {
			log.Printf("Failed to get staff ID for %s: %v", email, err)
			continue
		}

		for _, projectName := range projects {
			var projectID int
			err = db.QueryRow("SELECT id FROM projects WHERE name = ?", projectName).Scan(&projectID)
			if err != nil {
				log.Printf("Failed to get project ID for %s: %v", projectName, err)
				continue
			}

			// Check if relationship already exists
			var exists bool
			err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM staff_projects WHERE staff_id = ? AND project_id = ?)", staffID, projectID).Scan(&exists)
			if err != nil {
				log.Printf("Failed to check existing relationship for staff %d and project %d: %v", staffID, projectID, err)
				continue
			}

			// Insert only if the relationship does not exist
			if !exists {
				_, err := db.Exec(staffProjectInsertStmt, staffID, projectID)
				if err != nil {
					log.Printf("Failed to insert staff-project relationship for staff %d and project %d: %v", staffID, projectID, err)
				}
			} else {
				log.Printf("Skipping duplicate staff-project relationship for staff %d and project %d", staffID, projectID)
			}
		}
	}

	return nil
}
