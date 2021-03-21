package Fibonacci

import (
	G "Fibonacci/GRPC/apiGRPC"
	"Fibonacci/REST/apiREST"
	"Fibonacci/src"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct{
	name string
	x int32
	y int32
	exp map[int32]int32
	wantErr bool
}

var testCases = []testCase{
	{
		name: "x=0, y=1",
		x: 0,
		y: 1,
		exp: map[int32]int32{0: 0},
	},
	{
		name: "x=0, y=2",
		x: 0,
		y: 2,
		exp: map[int32]int32{0: 0, 1: 1},
	},
	{
		name: "x=0, y=3",
		x: 0,
		y: 3,
		exp: map[int32]int32{0: 0, 1: 1,2: 1},
	},
	{
		name: "x=0, y=4",
		x: 0,
		y: 4,
		exp: map[int32]int32{0: 0, 1: 1,2: 1,3: 2},
	},
	{
		name: "x=0, y=0",
		x: 0,
		y: 0,
		wantErr: true,
	},
	{
		name: "x=0, y=-1",
		x: 0,
		y: -1,
		wantErr: true,
	},
	{
		name: "x=-1, y=0",
		x: -1,
		y: 0,
		wantErr: true,
	},

}

func TestFibonacci(t *testing.T) {
	for _, tc := range testCases {
		x := fmt.Sprint(tc.x)
		y := fmt.Sprint(tc.y)

		result, err := src.GetFibonacci(x,y)
		if err != nil {
			if !tc.wantErr {
				t.Fatalf("Unexpected error: %s", err)
			}
		} else {
			assert.Equal(t, tc.exp, result) }
		}
	}

func TestREST(t *testing.T)  {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			r, _ := http.NewRequest("GET", "/fib", nil)
			w := httptest.NewRecorder()

			vars := map[string]string{
				"x": fmt.Sprint(tc.x),
				"y": fmt.Sprint(tc.y),
			}

			r = mux.SetURLVars(r, vars)
			apiREST.HandleGet(w, r)

			isJson := func(in []byte) bool {
				str := string(in)
				return strings.Contains(str, "{") && strings.Contains(str, "}")
			}

			result := w.Body.Bytes()

			if !isJson(result) {
				if !tc.wantErr {
					t.Fatalf("Unexpected error: %s", result)}
				} else {
					// преобразуем map массив байт
					js, _ := json.Marshal(tc.exp)
					js = append(js, 10)

					assert.Equal(t, js, result) }
			})
		}
}

func TestGRPS(t *testing.T) {
	s := G.GRPCServer{}

	for _, tc := range testCases {
		req := &G.FibonacciRequest{X: tc.x, Y: tc.y}
		resp, err := s.Get(context.Background(), req)

		if err == nil {
			assert.Equal(t, tc.exp, resp.Result )
			}

		if err != nil && !tc.wantErr {
			t.Fatalf("Unexpected error: %s", err.Error())
		}
	}

}