package handler

// User represent the person who is making the request
type User struct {
	Name     string `json:"name,omitempty" form:"name,omitempty" query:"name,omitempty"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}
