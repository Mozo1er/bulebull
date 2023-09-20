package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required,lte=20,gte=1,alphanum=Alphanumeric"`
	Password   string `json:"password" binding:"required,lte=20,gte=4,alphanum=Alphanumeric"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
