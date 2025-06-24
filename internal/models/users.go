package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UserRequests struct {
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Email *string `json:"email"`
}
