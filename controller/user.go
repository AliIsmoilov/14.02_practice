package controller

import (
	"app/models"
	"errors"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id string, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {

	users, err := c.store.User().GetList(req)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	return users, nil
}

func (c *Controller) GetUserByIdController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetUserById(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (c *Controller) UpdateUserController(req *models.UpdateUser) (models.User, error) {
	user, err := c.store.User().UpdateUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
func (c *Controller) DeleteUserController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().DeleteUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (c *Controller) WithdrawUserBalance(id string, balance float64) error {

	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{Id: id})
	if err != nil {
		return err
	}

	if user.Balance < balance{
		return errors.New("Not enough money")
	}

	user.Balance = user.Balance - balance

	_, err = c.store.User().UpdateUser(&models.UpdateUser{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Balance: user.Balance,
		
	})

	if err != nil {
		return err
	}

	err = c.store.ShopCart().UpdateProductStatus(id)

	if err != nil{
		return err
	}
	return nil
}

func (c *Controller) SendMoney(req *models.ExchangeMoney) (err error){

	sender, err := c.store.User().GetUserById(&models.UserPrimaryKey{req.SenderID})
	if err != nil{
		return err
	}

	reciver, err := c.store.User().GetUserById(&models.UserPrimaryKey{req.RecieverID})
	if err != nil{
		return err
	}

	if sender.Balance < req.Summa + (req.Summa*10/100){
		return errors.New("Not enough money")
	}

	_, err = c.store.User().UpdateUser(&models.UpdateUser{
		Id: sender.Id,
		Name: sender.Name,
		Surname: sender.Surname,
		Balance: sender.Balance - (req.Summa + req.Summa*10/100),
	}) 
	
	if err != nil{
		return err
	}

	_, err = c.store.User().UpdateUser(&models.UpdateUser{
		Id:  reciver.Id,
		Name: reciver.Name,
		Surname: reciver.Name,
		Balance: reciver.Balance + req.Summa,
	})

	if err != nil{
		return err
	}

	err = c.store.Komissiya().UpdateCommision(models.Komissiya{Balance: req.Summa*10/100})
	
	if err != nil{
		return err
	}


	return nil
	
}