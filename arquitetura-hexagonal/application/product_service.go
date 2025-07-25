package application

// A aplicação não depende e nem conhece que método de persistência usamos
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

	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	result, err := service.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
} 

func (service *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}

	result, err := service.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
} 

func (service *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}

	result, err := service.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
} 