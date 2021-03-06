package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/cleardataeng/gotoearth"
	"github.com/cleardataeng/gotoearth/example/foo"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle lambda event.
// This is what your shim would require.
func Handle(evt gotoearth.Event, ctx *runtime.Context) (interface{}, error) {
	r := gotoearth.Router{Handlers: map[string]gotoearth.Handler{
		"GET:/foo/{fooID}": foo.Handler{},
		"GET:/bar/{barID}": gotoearth.Lambda{lambda.InvokeInput{
			FunctionName:   aws.String("bar"),
			InvocationType: aws.String("Event"),
		}},
		"GET:/baz/{bazID}": gotoearth.SimpleLambda{"baz"},
	}}
	return r.Route(evt)
}

// main is simply used to mimic calls from the shim.
func main() {
	var res interface{}
	var err error
	evt1 := gotoearth.Event{
		Path:  map[string]string{"fooID": "foo1"},
		Route: "GET:/foo/{fooID}",
	}
	evt2 := gotoearth.Event{
		Path:  map[string]string{"barID": "bar2"},
		Route: "GET:/bar/{barID}",
	}
	evt3 := gotoearth.Event{
		Path:  map[string]string{"bazID": "baz3"},
		Route: "GET:/baz/{bazID}",
	}
	res, err = Handle(evt1, &runtime.Context{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("evt1: %s\n", res)
	res, err = Handle(evt2, &runtime.Context{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("evt2: %v\n", res)
	res, err = Handle(evt3, &runtime.Context{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("evt3: %v\n", res)
}
