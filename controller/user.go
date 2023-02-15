package controller

import (
	"app/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Controller) GetPkeyUser(req *models.UserPrimaryKey) (res *models.User, err error) {

	user, err := c.store.User().GetPkey(req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
