package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"graphql/graph/model"
	"log"
)

// CreateStaff creates a new staff member using the NewStaffInput.
func (r *mutationResolver) CreateStaff(ctx context.Context, input model.NewStaffInput) (*model.Staff, error) {
	staff := &model.Staff{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      input.Email,
		Role:       input.Role,
		Department: input.Department,
		Salary:     input.Salary,
	}

	err := r.DB.CreateStaff(staff)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

// AssignProjectToStaff assigns a project to a specific staff member.
func (r *mutationResolver) AssignProjectToStaff(ctx context.Context, staffID string, projectID string) (*model.Project, error) {
	project, err := r.DB.AssignProjectToStaff(staffID, projectID)
	if err != nil {
		return nil, err // Return error if project assignment fails
	}
	return project, nil
}

// GetStaffByID fetches a staff member by their ID.
func (r *queryResolver) GetStaffByID(ctx context.Context, id string) (*model.Staff, error) {
	return r.DB.GetStaffByID(id)
}

// GetStaffWithProjects retrieves the staff and their associated projects from the database
func (r *queryResolver) ListAllStaff(ctx context.Context) ([]*model.Staff, error) {
	var staffList []*model.Staff

	// Assuming r.DB implements the Query method
	rows, err := r.DB.Query(`
		SELECT s.id, s.first_name, s.last_name, s.department, s.role, s.salary 
		FROM staff s`)
	if err != nil {
		log.Println("Error fetching staff data:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the staff rows
	for rows.Next() {
		var staff model.Staff
		err := rows.Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Department, &staff.Role, &staff.Salary)
		if err != nil {
			log.Println("Error scanning staff data:", err)
			return nil, err
		}

		// Fetch associated projects for the current staff member
		var projects []*model.Project
		projectRows, err := r.DB.Query(`
			SELECT p.name, p.department 
			FROM projects p
			JOIN staff_projects sp ON p.id = sp.project_id
			WHERE sp.staff_id = ?`, staff.ID)
		if err != nil {
			log.Println("Error fetching projects for staff", staff.ID, ":", err)
			return nil, err
		}

		// Iterate through the project rows and add them to the staff member
		for projectRows.Next() {
			var project model.Project
			err := projectRows.Scan(&project.Name, &project.Department)
			if err != nil {
				log.Println("Error scanning project data:", err)
				return nil, err
			}
			projects = append(projects, &project)
		}
		staff.Projects = projects

		// Add the staff member to the result list
		staffList = append(staffList, &staff)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over staff rows:", err)
		return nil, err
	}

	return staffList, nil
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
