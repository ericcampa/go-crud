package db

import (
	"github.com/go-pg/pg/v10"
)

type Home struct {
	ID          int64  `json:"id"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	AgentID     int64  `json:"agent_id"`
	Agent       *Agent `pg:"rel:has-one" json:"agent"`
}

func CreateHome(db *pg.DB, req *Home) (*Home, error) {
	// _, err := db.Model(req).Insert()
	_, err := db.QueryOne(req, "INSERT INTO HOMES (price, description, address, agent_id) VALUES (?price, ?description, ?address, ?agent_id) RETURNING id", req)

	if err != nil {
		return nil, err
	}

	home := &Home{}
	_, saveErr := db.QueryOne(home, "select h.*, a.id as agent__id, a.name as agent__name  from homes h join agents a on h.agent_id=a.id where h.id = ?;", req.ID)

	if saveErr != nil {
		panic(saveErr)
	}

	return home, err
}

func GetHome(db *pg.DB, homeID string) (*Home, error) {
	home := &Home{}

	// err := db.Model(home).Relation("Agent").Where("home.id = ?", homeID).Select()
	_, err := db.QueryOne(home, "select h.*, a.id as agent__id, a.name as agent__name  from homes h join agents a on h.agent_id=a.id where h.id = ?;", homeID)

	if err != nil {
		panic(err)
	}

	return home, err
}
