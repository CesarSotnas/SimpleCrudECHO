package routes

const (
	Prefix       = "/api"
	Login        = "/login"
	usersPath    = "/users"
	GetUsers     = usersPath
	GetUsersById = usersPath + "/:user_id"
	CreateUsers  = usersPath
)
