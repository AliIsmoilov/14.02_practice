package controller

import(
	// "errors"

	"app/models"
)

func (c *Controller) GetAll(search string) ([]models.User, error){

	users, err := c.store.User().GetAll(search)
	
	if err != nil{
		return nil, err
	}

	return users, nil
}

func (c *Controller) Update(id string, UpdateUser models.User)error{

	err := c.store.User().Update(id, UpdateUser)

	if err != nil{
		return err
	}

	return nil
}

func (c *Controller) GetById(id string) (found_user models.User, err error){

	found_user, err =  c.store.User().GetById(id)

	if err != nil{
		return found_user, err
	}

	return found_user, nil
}


func (c *Controller) Delete(id string) error{

	err := c.store.User().Delete(id)

	return err
}