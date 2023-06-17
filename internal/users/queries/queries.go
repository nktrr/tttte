package queries

type UserQueries struct {
	GetUserByEmail GetUserByEmailHandler
}

func NewUserQueries(getUserByEmail GetUserByEmailHandler) *UserQueries {
	return &UserQueries{GetUserByEmail: getUserByEmail}
}

type GetUserByEmailQuery struct {
	UserEmail    string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
}

func NewGetUserByEmailQuery(userEmail string, userPassword string) *GetUserByEmailQuery {
	return &GetUserByEmailQuery{
		UserEmail:    userEmail,
		UserPassword: userPassword,
	}
}
