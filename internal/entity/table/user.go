package table

// User represents schema table users.
var User = user{
	ID:        "id",
	Username:  "username",
	FullName:  "full_name",
	Balance:   "balance",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Dot: &user{
		ID:        "users.id",
		Username:  "users.username",
		FullName:  "users.full_name",
		Balance:   "users.balance",
		CreatedAt: "users.created_at",
		UpdatedAt: "users.updated_at",
	},
}

type user struct {
	ID string

	Username string
	FullName string
	Balance  string

	CreatedAt string
	UpdatedAt string

	Dot *user
}

func (user) TableName() string {
	return "users"
}

// UserCredential represents schema table user_credentials.
var UserCredential = userCredential{
	ID:        "id",
	UserID:    "user_id",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Dot: &userCredential{
		ID:        "user_credentials.id",
		UserID:    "user_credentials.user_id",
		Password:  "user_credentials.password",
		CreatedAt: "user_credentials.created_at",
		UpdatedAt: "user_credentials.updated_at",
	},
}

type userCredential struct {
	ID string

	UserID   string
	Password string

	CreatedAt string
	UpdatedAt string

	Dot *userCredential
}

func (userCredential) TableName() string {
	return "user_credentials"
}
