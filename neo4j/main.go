package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	dbUri := "bolt://localhost:7687"

	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "test", ""), func(c *neo4j.Config) { c.Encrypted = false })
	checkError(err)
	defer driver.Close()

	session, err := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	checkError(err)
	defer session.Close()

	session, err = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	checkError(err)
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("CREATE (a:Person {name: $name})", map[string]interface{}{"name": "Andrey Shkunov"})
		if err != nil {
			return nil, err
		}

		return result.Consume()
	})

	checkError(err)

	getPeople(driver)

	session, err = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	checkError(err)
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(`MATCH (p:Person {name: 'Andrey Shkunov'})
		SET p.name = 'Andrey Brukhanov'
		RETURN p`, make(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		return result.Consume()
	})
	checkError(err)

	getPeople(driver)

	session, err = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	checkError(err)

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(`MATCH (p:Person {name: 'Andrey Brukhanov'})
		DELETE p`, make(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		return result.Consume()
	})

	checkError(err)

	getPeople(driver)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getPeople(driver neo4j.Driver) {
	session, err := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	checkError(err)

	defer session.Close()

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []interface{}

		result, err := tx.Run("MATCH (a:Person) RETURN a.name ORDER BY a.name", nil)
		checkError(err)

		for result.Next() {
			list = append(list, result.Record().Values())
		}

		err = result.Err()

		checkError(err)

		return list, nil
	})
	checkError(err)

	fmt.Println(people)
}
