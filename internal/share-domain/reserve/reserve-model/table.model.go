package reserve_model

type TableModel struct {
	ID        string `json:"ID"`
	MaxPeople int    `json:"maxPeople"`
}

func NewTable(
	ID string,
	MaxPeople int,
) TableModel {
	return TableModel{
		ID,
		MaxPeople,
	}
}

func (model TableModel) GetID() string     { return model.ID }
func (model TableModel) GetMaxPeople() int { return model.MaxPeople }
