package main

type ConnectionCache map[string]*DbConnection

var cache ConnectionCache

func init() {
	cache = make(ConnectionCache)
}

func AddConnection(id string, db *DbConnection) {
	cache[id] = db
}

func GetConnection(id string) *DbConnection {
	return cache[id]
}
