package domain

type CustomerRepositoryStub struct{
	customers []Customer
}

func (s CustomerRepositoryStub)FindAll()([]Customer,error){
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers:=[]Customer{
		{"1001","Rob","Roma","110011","2000-01-01","1"},
		{"1002","Bob","Roma","110011","2000-01-01","1"},
		{"1003","Pop","Roma","110011","2000-01-01","1"},
		{"1004","Bho","Roma","110011","2000-01-01","1"},
	}
	return CustomerRepositoryStub{customers: customers}
}