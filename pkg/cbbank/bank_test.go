package cbbank_test

import (
	"fmt"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/cbbank"
)

func TestUSD(t *testing.T) {
	usd := cbbank.USD()
	if usd == 0.0 {
		t.Error("USD: Доллар стоит 0.0. Ошибка")
	}
	t.Log(fmt.Sprintf("USD is %f RUB", usd))
}
