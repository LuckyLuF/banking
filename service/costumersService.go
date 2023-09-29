package service

import "github.com/LuckyLuF/banking/domain"

type CustomerService interface {
	GetAllCostumer()([]domain.Customer, error)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (s DefaultCustomerService)GetAllCostumer()([]domain.Customer,error){
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}