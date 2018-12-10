package membership

import (
	"github.com/ggalihpp/rbmc-backend/database"
	"gopkg.in/mgo.v2/bson"
)

func registerUser(data *Member) (err error) {
	dbSession, collection, err := database.Connection("membership")
	if err != nil {
		return
	}
	defer database.Disconnect(dbSession)

	err = collection.Insert(data)

	return
}

func updateMember(data *Member) (err error) {
	dbSession, collection, err := database.Connection("membership")
	if err != nil {
		return
	}
	defer database.Disconnect(dbSession)

	err = collection.Update(bson.M{"username": data.Username}, data)
	return
}

func getAllMembers(h database.QueryHelper) (result []Member, statusCode int, err error) {
	dbSession, col, err := database.Connection("membership")
	if err != nil {
		statusCode = 500
		return
	}
	defer database.Disconnect(dbSession)

	if !h.Ascending {
		h.SortBy = "-" + h.SortBy
	}

	err = col.Find(bson.M{}).Limit(h.Limit).Skip(h.Skip).Sort(h.SortBy).All(&result)
	if err != nil {
		statusCode = 500
		return
	}

	statusCode = 200
	return
}

func getMembersByTerritory(territory string, h database.QueryHelper) (result []Member, statusCode int, err error) {
	dbSession, col, err := database.Connection("membership")
	if err != nil {
		statusCode = 500
		return
	}
	defer database.Disconnect(dbSession)

	if !h.Ascending {
		h.SortBy = "-" + h.SortBy
	}

	err = col.Find(bson.M{"territory": territory}).Limit(h.Limit).Skip(h.Skip).Sort(h.SortBy).All(&result)
	if err != nil {
		statusCode = 500
		return
	}

	statusCode = 200
	return
}

// GetMemberByUsername - Will return you a data of a member
func GetMemberByUsername(username string) (result Member, statusCode int, err error) {
	dbSession, col, err := database.Connection("membership")
	if err != nil {
		statusCode = 500
		return
	}
	defer database.Disconnect(dbSession)

	err = col.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		statusCode = 404
		return
	}

	statusCode = 200
	return
}
