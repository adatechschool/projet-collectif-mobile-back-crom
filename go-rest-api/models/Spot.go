package models

type Spot struct {
	id          string `field: "id"`
	destination string
	photoURL    string
	address     string
	favorites   bool
}
