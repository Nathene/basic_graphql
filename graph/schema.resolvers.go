package graph

import (
	"context"
	"graphql/graph/model"
)

// CreateStaff creates a new staff member using the NewStaffInput.
func (r *mutationResolver) CreateStaff(ctx context.Context, input model.NewStaffInput) (*model.Staff, error) {
	staff := &model.Staff{
		ID:         generateID(),
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

// AssignProject assigns a project to a staff member.
func (r *mutationResolver) AssignProject(ctx context.Context, staffID string, projectID string) (*model.Project, error) {
	project, err := r.DB.AssignProjectToStaff(staffID, projectID)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// GetStaffByID fetches a staff member by their ID.
func (r *queryResolver) GetStaffByID(ctx context.Context, id string) (*model.Staff, error) {
	return r.DB.GetStaffByID(id)
}

// GetAllStaff retrieves all staff members.
func (r *queryResolver) GetAllStaff(ctx context.Context) ([]*model.Staff, error) {
	return r.DB.GetAllStaff()
}

// Helper function to generate a unique ID.
func generateID() string {
	// Replace with actual ID generation logic
	return "unique-id"
}

// ListAllStaff retrieves all staff members from the database.
func (r *queryResolver) ListAllStaff(ctx context.Context) ([]*model.Staff, error) {
	staff, err := r.DB.GetAllStaff()
	if err != nil {
		return nil, err // Return error if fetching staff fails
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
