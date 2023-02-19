package controller

import "app/models"

func (c *Controller) CreateKomissiya(req models.Komissiya) (err error){

	err = c.store.Komissiya().CreateKomissiya(&req)
	if err != nil{
		return err
	}

	return nil

}