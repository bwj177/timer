package webserver

import (
	"context"
	"errors"
	po2 "github.com/xiaoxuxiansheng/xtimer/common/model/po"
	"github.com/xiaoxuxiansheng/xtimer/common/model/vo"
	dao "github.com/xiaoxuxiansheng/xtimer/dao/user"
)

type UserService struct {
	dao *dao.UserDAO
}

func NewUserService(dao *dao.UserDAO) *UserService {
	return &UserService{dao: dao}
}

func (t *UserService) SignUp(ctx context.Context, req *vo.SignUpReq) error {
	po := &po2.User{UserName: req.UserName, Password: req.Password}
	err := t.dao.CreateUser(ctx, po)
	if err != nil {
		return err
	}
	return nil
}

func (t *UserService) Login(ctx context.Context, req *vo.LoginReq) error {
	user, err := t.dao.GetUser(ctx, dao.WithUserName(req.UserName))
	if err != nil {
		return err
	}
	if len(user) == 0 {
		return errors.New("用户不存在/密码错误")
	}
	return nil
}
