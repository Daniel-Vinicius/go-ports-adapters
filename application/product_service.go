package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (service *ProductService) Get(id string) (ProductInterface, error) {
	product, err := service.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
