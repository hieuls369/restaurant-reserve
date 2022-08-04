package reserve_model

type ReserveModel struct {
	ReserveID    string `json:"reserveID"`
	AmountPeople int    `json:"amountPeople"`
	PhoneNumber  int    `json:"phoneNumber"`
	Date         string `json:"date"`
	TableID      string `json:"tableID"`
}

func NewReserve(
	ReserveID string,
	AmountPeople int,
	PhoneNumber int,
	Date string,
	TableID string,
) ReserveModel {
	return ReserveModel{
		ReserveID,
		AmountPeople,
		PhoneNumber,
		Date,
		TableID,
	}
}

func (model ReserveModel) GetReserveID() string { return model.ReserveID }
func (model ReserveModel) GetAmountPeople() int { return model.AmountPeople }
func (model ReserveModel) GetPhoneNumber() int  { return model.PhoneNumber }
func (model ReserveModel) GetDate() string      { return model.Date }
func (model ReserveModel) GetTableID() string   { return model.TableID }
