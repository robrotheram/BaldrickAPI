package db

import (
	"gopkg.in/mgo.v2"
	"github.com/robrotheram/baldrick_engine/app/configuration"
//	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)
// A ValidationError is an error that is used when the required input fails validation.
// swagger:response Bot
type Bot struct{
	BotName   	string	    `bson:"name"`
	Type   		string    	`bson:"type"`
}

type User struct {
	Username	   	string   	`bson:"username"`
	Password   		string		`bson:"password"`
	Email   		string		`bson:"email"`
	TimeCreated   	string		`bson:"time_created"`
	Lastlogged   	string		`bson:"loggedin_time"`
}

type Rule struct {
	Name			string    	`bson:"name"`
	Type			string    	`bson:"type"`
	Formatter 		string    	`bson:"formatter"`

}

func NewRule(name, typeID, formatter string) Rule {
	return Rule{
		Name:   name,
		Type: typeID,
		Formatter: formatter,
	}
}

func NewUser(name, typeID, formatter string) User {
	now := time.Now()
	secs := now.Unix()

	return User{
		Username:   name,
		Password: typeID,
		Email: formatter,
		TimeCreated: time.Unix(secs, 0).String(),
	}
}

func NewBot(name, typeID string) Bot {
	return Bot{
		BotName: name,
		Type: typeID,
	}
}

var Session *mgo.Session;


func CreateSession(){
	Host := []string{
		configuration.Config.Host,
		// replica set addrs...
	}
	const (
		Username   = "YOUR_USERNAME"
		Password   = "YOUR_PASS"
		Database   = "YOUR_DB"
		Collection = "YOUR_COLLECTION"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,
		// Username: Username,
		// Password: Password,
		// Database: Database,
		// DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
		// },
	})
	if err != nil {
		panic(err)
	}
	Session = session
}
func GenerateBot(){
	c := Session.DB("YOUR_DB").C("BOT")
	bot := NewBot("Test","testID")
	if err := c.Insert(bot); err != nil {
		panic(err)
	}
	fmt.Println("Inserted a new BOT");
}


func Close()  {
	Session.Close();
}