package application

import "errors"

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (service *ProductService) Get(id string) (ProductInterface, error) {
	product, err := service.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	productIsValid, invalidProductError := product.IsValid()

	if !productIsValid {
		return &Product{}, invalidProductError
	}

	productCreated, persistenceError := service.Persistence.Save(product)

	if persistenceError != nil {
		return &Product{}, persistenceError
	}

	return productCreated, nil
}

func (service *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if product.GetStatus() == ENABLED {
		return &Product{}, errors.New("product already is enabled")
	}

	errorEnablingProduct := product.Enable()

	if errorEnablingProduct != nil {
		return &Product{}, errorEnablingProduct
	}

	productCreated, persistenceError := service.Persistence.Save(product)

	if persistenceError != nil {
		return &Product{}, persistenceError
	}

	return productCreated, nil
}

func (service *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if product.GetStatus() == DISABLED {
		return &Product{}, errors.New("product already is disabled")
	}

	errorEnablingProduct := product.Disable()

	if errorEnablingProduct != nil {
		return &Product{}, errorEnablingProduct
	}

	productCreated, persistenceError := service.Persistence.Save(product)

	if persistenceError != nil {
		return &Product{}, persistenceError
	}

	return productCreated, nil
}
