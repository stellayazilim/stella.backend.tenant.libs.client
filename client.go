package client

import (
	"fmt"
	"net/http"
)

type IUserClient interface {
}

type userClient struct {
	TenantId string
}

func UserClient(i string) IUserClient {
	return &userClient{
		TenantId: i,
	}
}

type client struct {
	user IUserClient
}

func (u *userClient) GetUsers(l int, s int) (http.Response, error) {

	if resp, err := http.Get(fmt.Sprintf("http://localhost:8080/tenant/%v/users", u.TenantId)); err == nil {
		return *resp, nil
	}
	return http.Response{}, fmt.Errorf("can not get user(s) of tenant")
}
