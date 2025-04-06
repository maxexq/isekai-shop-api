package repository

type itemShopRepositoryMock struct{}

func NewItemShopRepositoryMock() ItemShopRepository {
	return &itemShopRepositoryMock{}
}
