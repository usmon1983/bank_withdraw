package card

import 	"github.com/usmon1983/bank/pkg/bank/types"

//Issue создаёт экземпляр карты с предопределенными полями
func Issue(currency types.Currency, color string, name string)  {
	return types.Card {
		ID:			1000,
		PAN: 		"5058 xxxx xxxx 0001",
		Balance:	0,
		Currency:	currency,
		Color:		color,
		Name: 		name,
		Active:		true,
	}
}