package opensecrets

import (
	"testing"
	"log"
)


func TestOpenSecretsQuery(t *testing.T) {

	args := []map[string]string{}
	arg := map[string]string{}
	arg["key"] = "cid"
	arg["value"] = lastName
	args = append(args, arg)

	method := "candContrib"

	// @todo set up test
	// data := env.Votesmart.Query(method, args)
}
