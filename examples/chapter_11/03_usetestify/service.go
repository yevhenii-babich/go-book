package usetestify

type Data struct {
	Value string
}

type DataService interface {
	GetData(id string) (*Data, error)
}

type Service struct{}

func (Service) GetData(id string) (*Data, error) {
	return &Data{Value: "Hello, " + id}, nil
}
