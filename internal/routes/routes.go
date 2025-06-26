package routes

const (
	Prefix       = "/api"
	Login        = "/login"
	usersPath    = "/users"
	userIDParam  = "/:user_id"
	GetUsers     = usersPath
	GetUsersById = usersPath + userIDParam
	CreateUsers  = usersPath
	UpdateUsers  = usersPath + userIDParam
	DeleteUsers  = usersPath + userIDParam
)
