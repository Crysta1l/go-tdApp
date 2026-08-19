package users_postgres_repository

type userModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}
