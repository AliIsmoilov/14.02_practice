package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type komissiyaRepo struct{
	fileName string
	file     *os.File
}

// Constructor
func NewKomissiyaRepo(fileName string, file *os.File) *komissiyaRepo {
	return &komissiyaRepo{
		fileName: fileName,
		file:     file,
	}
}

func (k *komissiyaRepo) CreateKomissiya(req *models.Komissiya) (err error){

	var commisions []*models.Komissiya
	err = json.NewDecoder(k.file).Decode(&commisions)
	if err != nil {
		return err
	}

	commisions = append(commisions, &models.Komissiya{Balance: req.Balance})


	body, err := json.MarshalIndent(commisions, "", "   ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(k.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (k *komissiyaRepo) UpdateCommision(req models.Komissiya) (err error){

	var commisions []*models.Komissiya
	err = json.NewDecoder(k.file).Decode(&commisions)
	if err != nil {
		return err
	}
	
	commisions[0].Balance += req.Balance 

	body, err := json.MarshalIndent(commisions, "", "   ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(k.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}