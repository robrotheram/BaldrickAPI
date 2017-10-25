package db

import (
	"gopkg.in/mgo.v2"
	"github.com/robrotheram/baldrick_engine/app/configuration"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
)
// A ValidationError is an error that is used when the required input fails validation.
// swagger:response Bot
type Bot struct{
	BotName   	string	    `bson:"name"`
	Channel		string    	`bson:"channels"`
	Rules   	[]Rule    	`bson:"rules"`
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
	Prifix			string    	`bson:"prifix"`
	Function		string    	`bson:"function"`
	Paramaters		string    	`bson:"paramaters"`
}

func (r*Rule) Process() string {
	return "hello this is a test";
}


type Channel struct {
	Name		string	 		`bson:"name"`
	Description string 			`bson:"description"`
	StartTime	string			`bson:"start_time"`
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

func NewBot() Bot {

	r1 := Rule{
		Name:   "testName",
		Prifix: "!broadcast",
		Function: "hello",
		Paramaters: "",
	}
	r2 := Rule{
		Name:   "NewRule",
		Prifix: "!hello",
		Function: "bob",
		Paramaters: "",
	}



	rules := []Rule{r1,r2}

	return Bot{
		BotName: "BotName",
		Channel: "discord",
		Rules:	rules,
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
	bot := NewBot()
	if err := c.Insert(bot); err != nil {
		panic(err)
	}
	fmt.Println("Inserted a new BOT");
}

func GetRulesFromChannel (channel string) []Rule {
	c := Session.DB("YOUR_DB").C("BOT")
	var result []Bot;
	err := c.Find(bson.M{"channels":channel}).All(&result)
	if err != nil {
		panic(err)
	}
	if (len(result) > 0){
		return result[0].Rules;
	} else {
		return []Rule{}
	}
}

func Close()  {
	Session.Close();
}