package graphql

import (
	"net/http"
	"net/http/httptest"
	"siop/noticias/pkg/apis/service"
	"testing"

	"github.com/rodrigobotelho/graphql-kit"
)

func TestHandler(t *testing.T) {
	type args struct {
		request, response string
	}
	tests := []struct {
		name string
		args []args
	}{
		//TODO: acrescentar casos
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewBasicExampleService()
			h := NewHandler(svc, "schema.graphql", string(graphqlkit.Secret), nil)
			mux := http.NewServeMux()
			mux.Handle("/", h.Handler())
			for _, args := range tt.args {
				req, _ := graphqlkit.CreateGraphqlRequest(args.request)
				rec := httptest.NewRecorder()
				mux.ServeHTTP(rec, req)
                var v interface{}
				err := json.Unmarshal([]byte(args.response), &v)
				if err != nil {
					t.Error(err)
				}
				b, _ := json.Marshal(v)
				resp := string(b)
				if !strings.Contains(resp, "data") &&
					!strings.Contains(resp, "errors") {
					resp = fmt.Sprintf(`{"data":%v}`, resp)
				}
				if rec.Body.String() != resp {
					t.Errorf(
						"Body = %v, want %v",
						rec.Body.String(),
						resp,
					)
				}
			}
		})
	}
}
