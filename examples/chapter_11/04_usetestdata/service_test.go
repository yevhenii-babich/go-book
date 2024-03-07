// Пакет usetestdata
package usetestdata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/stretchr/testify/assert"  // імпортуємо пакет для перевірки умов
	"github.com/stretchr/testify/require" // імпортуємо пакет для вимог
	"testing"                             // імпортуємо пакет для тестування

	"usetestdata/products" // імпортуємо пакет з продуктами
)

// TestService_GetProductList - тестова функція для перевірки отримання списку продуктів
func TestService_GetProductList(t *testing.T) {
	// визначаємо структуру для аргументів
	type args struct {
		from float64
		to   float64
	}
	// завантажуємо тестові дані
	toCompare := loadFixture(t)
	// визначаємо тести
	tests := []struct {
		name    string
		fields  products.ClientAPI
		args    args
		want    ProductStat
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "без помилок",
			fields: &mockProductAPI{err: nil, t: t},
			args:   args{from: toCompare.Price, to: toCompare.Price},
			want: ProductStat{
				Min:      toCompare.Price,
				Max:      toCompare.Price,
				Avg:      toCompare.Price,
				Products: []products.Product{toCompare, toCompare, toCompare, toCompare},
			},
			wantErr: assert.NoError,
		},
		{
			name:    "помилка API",
			fields:  &mockProductAPI{err: assert.AnError, t: t},
			args:    args{from: toCompare.Price, to: toCompare.Price},
			want:    ProductStat{},
			wantErr: assert.Error,
		},
	}
	// запускаємо тести
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				productAPI: tt.fields,
			}
			got, err := s.GetProductList(tt.args.from, tt.args.to)
			tt.wantErr(t, err)
			if err != nil {
				assert.Empty(t, got)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestService_GetProduct(t *testing.T) {
	expectedData := loadFixture(t)
	tests := []struct {
		name    string
		fields  products.ClientAPI
		args    int
		want    *products.Product
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "без помилок",
			fields:  &mockProductAPI{err: nil, t: t},
			args:    2,
			want:    &expectedData,
			wantErr: assert.NoError,
		},
		{
			name:    "API повертає помилку",
			fields:  &mockProductAPI{err: assert.AnError, t: t},
			args:    2,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				productAPI: tt.fields,
			}
			got, err := s.GetProduct(tt.args)
			tt.wantErr(t, err, fmt.Sprintf("GetProduct(%v)", tt.args))
			if err != nil {
				assert.Nil(t, got)
				return
			}
			t.Logf("Got: %+v", got)
			assert.Equalf(t, tt.want, got, "GetProduct(%v)", tt.args)
		})
	}
}

// ---- MOCKS ----
// mockProductAPI - макет API продуктів
type mockProductAPI struct {
	err error
	t   *testing.T
}

// GetProduct - отримує продукт за ID
func (m *mockProductAPI) GetProduct(id int) (*products.Product, error) {
	p := loadFixture(m.t)
	p.ID = id
	if m.err != nil {
		return nil, m.err
	}
	return &p, m.err
}

// GetProductList - отримує список продуктів
func (m *mockProductAPI) GetProductList(from, to float64) (products.ProductList, error) {
	p := loadFixture(m.t)
	return products.ProductList{
		Products: []products.Product{p, p, p, p},
		Total:    4,
		Range:    products.RangeInList{From: from, To: to},
	}, m.err
}

// loadFixture - завантажує тестові дані
func loadFixture(t *testing.T) products.Product {
	t.Helper()
	data, err := os.ReadFile("testdata/example_response.json")
	require.NoError(t, err)
	var p products.Product
	err = json.Unmarshal(data, &p)
	require.NoError(t, err)
	return p
}
