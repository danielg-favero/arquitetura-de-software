// Definição de um produto

package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// Por padrão no GO o método "init" executa primeiro
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Contrato de definição de um Produto
// Metodos que começam com letra MAIUSCULA, podem ser acessados de forma externa
// Metodos que começam com letra minuscula, NÃO podem ser acessados de forma externa
type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	ChangePrice(price float64) error
}

// Definição do serviço de Produto
type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

// Interfaces para criar o contrato entre a aplicação e a comunicação com o Banco de Dados
type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWrite interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWrite
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

// Definição do tipo Produto
type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Price float64 `valid:"float,optional"`
	Status string `valid:"required"`
}

// Criar novo produto (semelhante ao construtor de classes em POO)
func NewProduct() *Product {
	product := Product {
		ID:		uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

// Implementação dos métodos da interface
func (product *Product) IsValid() (bool, error) {
	if product.Status == ""{
		product.Status = DISABLED
	}

	if product.Status != ENABLED && product.Status != DISABLED {
		return false, errors.New("O status do produto precisa ser ENABLED ou DISABLED")
	}

	if product.Price < 0 {
		return false, errors.New("O preço do produto precisa ser maior que 0")
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (product *Product) Enable() error {
	if product.Price > 0 {
		product.Status = ENABLED
		return nil
	}

	return errors.New("O preço precisa ser maior que 0 para habilitar o produto")
}

func (product *Product) Disable() error {
	if product.Price == 0 {
		product.Status = DISABLED
		return nil
	}

	return errors.New("O preço precisa ser 0 para desabilitar o produto")
}

func (product *Product) GetID() string {
	return product.ID
}

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) GetStatus() string {
	return product.Status
}

func (product *Product) GetPrice() float64 {
	return product.Price	
}

func (product *Product) ChangePrice(price float64) error {
	if product.Price < 0 {
		return errors.New("Preço precisa ser maior que 0")
	}
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return err
	}
	return nil
}