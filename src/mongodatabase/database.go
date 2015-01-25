package mongodatabase

//For Database I've choose MongoDB !!
import (
	//	"APIs/youtube"
	"config"
	//	"fmt"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Session *mgo.Session
	DB      *mgo.Database
}

//Connect using config constants
func (m *Mongo) Connect() (err error) {
	m.Session, err = mgo.Dial(config.MONGODB_CONNECTION)
	if err != nil {
		return
	}
	m.Session.SetMode(mgo.Monotonic, true)
	m.DB = m.Session.DB(config.MONGO_DATABASE)
	return
}

func (m *Mongo) Insert(collection string, docs ...interface{}) (err error) {
	c := m.DB.C(collection)
	err = c.Insert(docs)
	return
}

func (m *Mongo) Find(collection string, conditions map[string]interface{}) *mgo.Query {
	c := m.DB.C(collection)
	return c.Find(conditions)
}

func (m *Mongo) CloseConnection() {
	m.Session.Close()
}
