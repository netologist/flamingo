package vodka

type DbClientFactory struct {
	loggerFactory *LoggerFactory
}

type Person struct {
	Name  string
	Phone string
}

func NewDbClientFactory(loggerFactory *LoggerFactory) *DbClientFactory {
	/*
		session, err := mgo.Dial("server1.example.com,server2.example.com")
		if err != nil {
			panic(err)
		}
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

		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)
	*/
	return &DbClientFactory{loggerFactory}
}
