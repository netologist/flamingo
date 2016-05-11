package vodka

import "github.com/spf13/viper"

type DbClientFactory struct {
	loggerFactory *LoggerFactory
	MongoClient   MongoClient
}

// type Person struct {
// 	Name  string
// 	Phone string
// }

func NewDbClientFactory(loggerFactory *LoggerFactory) *DbClientFactory {
	uri := viper.GetString("mongo.uri")
	mongoClient := NewMongoClient(uri, loggerFactory)
	return &DbClientFactory{loggerFactory, mongoClient}
	/*
		logger := loggerFactory.NewLogger("DbClientFactory")
		session, err := mgo.Dial("localhost")

		if err != nil {
			panic(err)
		}
		//logger := loggerFactory.NewLogger("asd").(log.Logger)
		//mgo.SetLogger(logger)
		mgo.SetDebug(true)
		defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)

		c := session.DB("test").C("people")
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		c.FindId(id interface{})
		query := c.Find(bson.M{"name": "Ale"})
		logger.Info("Query %#v", query)
		err = query.One(&result)
		logger.Info("Query %#v, Result %#v", query, result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)
	*/

}
