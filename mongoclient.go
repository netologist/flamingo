package vodka

type Selectors map[string]interface{}

/******************************************
	Mongo Query
 ******************************************/

type MongoQuery struct {
}

func (ms *MongoQuery) One(result interface{}) error {
	return nil
}

func (ms *MongoQuery) All(result interface{}) error {
	return nil
}

/******************************************
	Mongo Collection
 ******************************************/
type MongoCollection struct {
}

func (ms *MongoCollection) Query(selectors Selectors) *MongoQuery {
	/*	session, err := mgo.Dial("server1.example.com,server2.example.com")
		if err != nil {
			panic(err)
		}
		defer session.Close()*/
	return &MongoQuery{}
}

func (ms *MongoCollection) FindById(id interface{}) *MongoQuery {
	return &MongoQuery{}
}

func (ms *MongoCollection) Insert(entity interface{}) error {
	return nil
}

func (ms *MongoCollection) Update(entity interface{}) error {
	return nil
}

func (ms *MongoCollection) Remove(selectors Selectors) error {
	return nil
}

func (ms *MongoCollection) RemoveById(id interface{}) error {
	return nil
}

/******************************************
	Mongo Session
 ******************************************/

type MongoSession struct {
}

func (ms *MongoSession) Close() error {
	return nil
}

func (ms *MongoSession) Db(name string) *MongoSession {
	return ms
}

func (ms *MongoSession) Collection(name string) *MongoCollection {
	return &MongoCollection{}
}

/******************************************
	Mongo client
 ******************************************/

type MongoClient struct {
	loggerFactory *LoggerFactory
}

func (m *MongoClient) OpenSession() (MongoSession, error) {
	return MongoSession{}, nil
}

func NewMongoClient(uri string, loggerFactory *LoggerFactory) MongoClient {
	return MongoClient{loggerFactory}
}
