package simulator

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/domain/services/simulator/tests"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
	"testing"
)
var injection tests.MockRepository

func TestCreate (t *testing.T) {
	data := tests.MockCreate()

	result, err := Create(&data, injection)
	if result != nil {
		return
	}
	t.Error("Expect: ", data, "Received: ", err )
}

func BenchmarkCreate (b *testing.B) {
	data := tests.MockCreate()

	for i :=0; i< b.N; i++ {
		Create(&data, injection)
	}
}

func TestList (t *testing.T) {
	data := tests.MockCreate()

	result, err := List([]entities.Simulator{}, injection)
	if result != nil {
		return
	}
	t.Error("Expect: ", data, "Received:", err )
}

func BenchmarkList (b *testing.B) {
	for i :=0; i< b.N; i++ {
		List([]entities.Simulator{}, injection)
	}
}

func TestUpdate (t *testing.T) {
	data := tests.MockCreate()
	data.Role = "Co-Piloto"

	result, err := Update(&data, injection)
	if result != nil {
		return
	}
	t.Error("Expect: ", data, "Received:", err )
}

func BenchmarkUpdate (b *testing.B) {
	data := tests.MockCreate()
	data.Role = "Co-Piloto"

	for i :=0; i< b.N; i++ {
		Update(&data, injection)
	}
}

func TestDelete (t *testing.T) {
	id, _ := uuid.Parse("3bdbc864-51bd-4a6f-a7e9-667970793947")
	result := Delete(id, injection)
	if result != nil {
		t.Error("Expect:", nil, "Received:", result)
	}
	return
}

func BenchmarkDelete (b *testing.B) {
	id, _ := uuid.Parse("3bdbc864-51bd-4a6f-a7e9-667970793947")
	for i :=0; i< b.N; i++ {
		Delete(id, injection)
	}
}
