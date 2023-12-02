package graph

import "github.com/lucasferreirajs/13-GraphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// onde injetamos dependência
type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
