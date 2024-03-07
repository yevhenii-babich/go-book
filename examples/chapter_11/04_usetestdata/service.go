package usetestdata

import "usetestdata/products"

type Service struct {
	productAPI products.ClientAPI
}
type ProductStat struct {
	Min      float64
	Max      float64
	Avg      float64
	Products []products.Product
}

func New(productAPI products.ClientAPI) *Service {
	return &Service{productAPI: productAPI}
}

func (s *Service) GetProduct(id int) (*products.Product, error) {
	return s.productAPI.GetProduct(id)
}

func (s *Service) GetProductList(from, to float64) (ProductStat, error) {
	list, err := s.productAPI.GetProductList(from, to)
	if err != nil {
		return ProductStat{}, err
	}
	list.Range.From = 9000000000 - 1
	list.Range.To = -1
	list.Total = len(list.Products)
	totalPrice := 0.0 // reset the total price
	for _, v := range list.Products {
		list.Range.From = min(list.Range.From, v.Price)
		list.Range.To = max(list.Range.To, v.Price)
		totalPrice += v.Price
	}
	return ProductStat{
		Min:      list.Range.From,
		Max:      list.Range.To,
		Avg:      totalPrice / float64(list.Total), // calculate the average price
		Products: list.Products,
	}, nil
}
