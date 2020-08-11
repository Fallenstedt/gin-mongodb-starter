package models


type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID(id string) (*User, error) {
	// Pretend to access a database
	user := User {
		ID: id,
		Name: "Alex",
		Birthday: "02/05/1991",
		Gender: "Male",
		PhotoURL: "some/photo/url",
		Time: 84928942,
		Active: true,
		UpdatedAt: 8942894289,
	}

	return &user, nil
}
