package members

type Member struct{
	ID string `json:"id"`
	FullName string `json:"fullName"`
	Initials string `json:"initials"`
	Username string `json:"username"`
	Email string `json:"email"`
}