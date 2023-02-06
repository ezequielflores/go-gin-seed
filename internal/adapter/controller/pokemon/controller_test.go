package pokemon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/redbeestudios/go-seed/mocks"
	"github.com/redbeestudios/go-seed/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type controllerDependencies struct {
	getByName   *mocks.MockGetByName
	savePokemon *mocks.MockSavePokemon
}

func makePokemonControllerDependencies(t *testing.T) *controllerDependencies {
	return &controllerDependencies{
		getByName: mocks.NewMockGetByName(gomock.NewController(t)),
	}
}

func TestNewPokemonController(t *testing.T) {
	dependencies := makePokemonControllerDependencies(t)

	controller := NewPokemonController(
		dependencies.getByName,
		dependencies.savePokemon,
	)
	assert.NotNil(t, controller)
}

func TestGetPokemon(t *testing.T) {
	pokemon := testdata.Pokemon()

	type test struct {
		name         string
		mock         func(*controllerDependencies)
		expectedBody string
		expectedCode int
		pathParam    string
	}

	tests := []test{
		{
			name: "Get pokemon by name",
			mock: func(dependencies *controllerDependencies) {
				dependencies.getByName.EXPECT().
					Get(gomock.Any(), pokemon.Name()).
					Return(pokemon, nil)
			},
			expectedBody: `
				{
					"id": 3,
					"name": "venusaur",
					"type": "grass"
				}
			`,
			expectedCode: 200,
			pathParam:    "venusaur",
		},
		{
			name: "500 if service fails to return pokemon",
			mock: func(dependencies *controllerDependencies) {
				dependencies.getByName.EXPECT().
					Get(gomock.Any(), pokemon.Name()).
					Return(nil, fmt.Errorf("Internal server error"))
			},
			expectedCode: 500,
			pathParam:    "venusaur",
		},
		{
			name: "400 If It's invalid pokemon name",
			mock: func(dependencies *controllerDependencies) {
				dependencies.getByName.EXPECT().
					Get(gomock.Any(), pokemon.Name()).
					Times(0)

			},
			expectedCode: 400,
			pathParam:    "_",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dependencies := makePokemonControllerDependencies(t)
			if test.mock != nil {
				test.mock(dependencies)
			}

			controller := NewPokemonController(
				dependencies.getByName,
				dependencies.savePokemon,
			)
			router := gin.Default()
			router.GET("/pokemon/:name", controller.GetPokemon)
			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/pokemon/%s", test.pathParam),
				bytes.NewReader([]byte(test.expectedBody)),
			)
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, test.expectedCode, rec.Code)

			if test.expectedBody != "" {
				buff := new(bytes.Buffer)
				err := json.Compact(buff, []byte(test.expectedBody))
				assert.NoError(t, err)
				assert.Equal(t, buff.String(), rec.Body.String())
			}

		})
	}

}
