package woocommerce_test

import (
	"reflect"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func TestMapTags(t *testing.T) {
	testInput := []woocommerce.Tag{{Name: "baby", Slug: "baby", Id: 53}, {Name: "Bags", Slug: "bags", Id: 56}, {Name: "barber", Slug: "barber", Id: 59}, {Name: "Beachwear", Slug: "beachwear", Id: 63}, {Name: "Belt", Slug: "belt", Id: 67}, {Name: "besties gifts", Slug: "besties-gifts", Id: 68}, {Name: "Book", Slug: "book", Id: 74}, {Name: "Boy", Slug: "boy", Id: 437}, {Name: "Bra", Slug: "bra", Id: 76}, {Name: "Coffee", Slug: "coffee", Id: 91}, {Name: "Denim", Slug: "denim", Id: 101}, {Name: "Dress", Slug: "dress", Id: 103}, {Name: "electronic", Slug: "electronic", Id: 108}, {Name: "Fashion", Slug: "fashion", Id: 110}, {Name: "Food", Slug: "food", Id: 113}, {Name: "furniture", Slug: "furniture", Id: 118}, {Name: "Girl", Slug: "girl", Id: 436}, {Name: "Hats", Slug: "hats", Id: 129}, {Name: "House paint", Slug: "house-paint", Id: 132}, {Name: "jacket", Slug: "jacket", Id: 140}, {Name: "juice", Slug: "juice", Id: 145}, {Name: "Man", Slug: "man", Id: 433}, {Name: "Minimog", Slug: "minimog", Id: 159}, {Name: "Nail", Slug: "nail", Id: 164}, {Name: "Notebook", Slug: "notebook", Id: 167}, {Name: "Pet lovers", Slug: "pet-lovers", Id: 173}, {Name: "plants", Slug: "plants", Id: 179}, {Name: "Print", Slug: "print", Id: 184}, {Name: "Sandal", Slug: "sandal", Id: 191}, {Name: "Shirt", Slug: "shirt", Id: 193}, {Name: "Shop our look", Slug: "shop-our-look", Id: 198}, {Name: "Sneaker", Slug: "sneaker", Id: 208}, {Name: "socks", Slug: "socks", Id: 212}, {Name: "Sunglasses", Slug: "sunglasses", Id: 215}, {Name: "Toy", Slug: "toy", Id: 227}, {Name: "Unisex", Slug: "unisex", Id: 435}, {Name: "Unisex Adults", Slug: "unisex_adults", Id: 438}, {Name: "Vagabond", Slug: "vagabond", Id: 232}, {Name: "watch", Slug: "watch", Id: 238}, {Name: "Women", Slug: "women", Id: 434}}
	testAnswer := map[string]int{"man": 433, "women": 434}

	MatTags_test := woocommerce.MapTags(testInput)
	if !reflect.DeepEqual(MatTags_test, testAnswer) {
		t.Log("Неверно обработаный вывод мапы")
	}
}
