// Code generated by goctl. DO NOT EDIT.
package types

type SignupRequest struct {
	Username   string `path:"name,options=dengliwei|no"`
	Password   string `form:"password"`
	Gender     int    `json:"gender,options=0|1|2,default=0"`
	RePassword string `form:"re_password"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username   string `path:"name,options=dengliwei|no"`
	Password   string `form:"password"`
	Gender     int    `json:"gender,options=0|1|2,default=0"`
	RePassword string `form:"re_password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}
