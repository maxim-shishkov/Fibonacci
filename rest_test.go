package Fibonacci

import (
	"Fibonacci/REST/apiREST"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)




type in struct {
	x string
	y string
}


func TestREST(t *testing.T)  {

	testCases := []struct{
		name string
		in in
		exp []byte
	}{
		{	name: "zero",
			in: in{x: "0",y: "1"},
			exp: []byte(`{"0":0}`+"\n"),
		},
		{	name: "one",
			in: in{x: "0",y: "2"},
			exp: []byte(`{"0":0,"1":1}`+"\n"),
		},
		{	name: "two",
			in: in{x: "0",y: "3"},
			exp: []byte(`{"0":0,"1":1,"2":1}`+"\n"),
		},
		{	name: "three",
			in: in{x: "0",y: "4"},
			exp: []byte(`{"0":0,"1":1,"2":1,"3":2}`+"\n"),
		},
		{	name: "four",
			in: in{x: "0",y: "0"},
			exp: []byte(`Invalid input  x >= y`),
		},

	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r, _ := http.NewRequest("GET", "/fib", nil)
			w := httptest.NewRecorder()

			vars := map[string]string{
				"x": tc.in.x,
				"y": tc.in.y,
			}

			r = mux.SetURLVars(r, vars)

			apiREST.HandleGet(w,r)

			assert.Equal(t, tc.exp, w.Body.Bytes())

		})
	}


}

func TestGRPC(t *testing.T) {

}
