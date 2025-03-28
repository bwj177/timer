package vo

type SignUpReq struct {
	UserName   string `form:"user_name" json:"user_name" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	RePassword string `form:"re_password" json:"re_password" binding:"required"`
}

type SignUpResp struct {
	CodeMsg
}

func NewSignUpResp(codeMsg CodeMsg) *SignUpResp {
	return &SignUpResp{
		CodeMsg: codeMsg,
	}
}

type LoginReq struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResp struct {
	CodeMsg
}

func NewLoginResp(codeMsg CodeMsg) *LoginResp {
	return &LoginResp{
		CodeMsg: codeMsg,
	}
}
