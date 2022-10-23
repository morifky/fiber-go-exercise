package handler

type userSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userCreateRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userUpdateRequest struct {
	Email    string `json:"email" validate:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type postCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type postUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
