package main

import (
	"log"
	"testing"
)

func TestAny(t *testing.T) {
	/*
		SELECT * FROM `travel-sample`.inventory.route
		WHERE airline="KL" AND sourceairport="ABQ" AND destinationairport="ATL"
		AND ANY departure IN schedule SATISFIES departure.utc > "23:41" END;
	*/

	statement := Select("*").From("travel-sample.inventory.route").
		Where(
			X("airline").Eq(S("KL")).
				And("sourceairport").Eq(S("ABQ")).
				And("destinationairport").Eq(S("ATL")).
				And(
					Any("departure").In("schedule").Satisfies("departure.utc").Gt(S("23:41")),
				).
				End(),
		)

	log.Println(statement.String())
}
