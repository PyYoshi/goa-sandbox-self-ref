package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("selfRef", func() {
	Title("Self Reference")
})

var Person = Type("Person", func() {
	Field(1, "person_id", String)
	Field(2, "name", String)
	Field(3, "friends", ArrayOf("Person"))
})

var _ = Service("selfRef", func() {
	Method("retrieve_person", func() {
		Payload(func() {
			Field(1, "person_id", String)
			Required("person_id")
		})
		Result(Person)

		GRPC(func() {
			Response(StatusOK)
		})

		HTTP(func() {
			Response(StatusOK)
		})
	})
})
