package product

type ProductUseCase struct {
	Repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		Repository: repository,
	}
}

func (usecase *ProductUseCase) GetProduct(id int) (Product,error) {
	return usecase.Repository.GetProduct(id),nil
}