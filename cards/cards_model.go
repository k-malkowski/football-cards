package card

type Card struct {
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Club        string     `json:"club"`
	Nationality string     `json:"nationality"`
	Statistics  Statistics `json:"statistics"`
}

type Statistics struct {
	Pac int8 `json:"pac"`
	Sho int8 `json:"sho"`
	Pas int8 `json:"pas"`
	Dri int8 `json:"dri"`
	Def int8 `json:"def"`
	Phy int8 `json:"phy"`
}

type Cards []Card

var CardsList Cards = make(Cards, 0)

func (c *Cards) Add(card Card) {
	*c = append(*c, card)
}

func (s *Statistics) GetOverall() int8 {
	return (s.Def + s.Dri + s.Pac + s.Pas + s.Phy + s.Sho) / 6
}
