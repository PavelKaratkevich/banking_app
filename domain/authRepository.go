package domain

import (
	"banking/logger"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Port
type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)
	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		log.Printf("RESPONSE looks like %v", response) // {200 OK 200 HTTP/1.1 1 1 map[Content-Length:[22] Content-Type:[application/json] Date:[Sun, 08 Aug 2021 19:08:30 GMT]] 0xc0000de280 22 [] false false map[] 0xc0000ee300 <nil>}
		m := map[string]bool{}

		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())

			return false
		}
		log.Printf("M looks like %v", m) // map[isAuthorized:true]
		return m["isAuthorized"]
	}
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: ":8181", Path: "/auth/verify", Scheme: "http"}
	log.Printf("U looks like: %v", u) // {http   :8181 /auth/verify  false   }
	q := u.Query()
	log.Printf("Q looks like: %v", q) // map[]
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	log.Printf("Q looks like: %v", q) // map[customer_id:[2001] routeName:[GetCustomer] token:[eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6WyI5NTQ3MiIsIjk1NDczIiwiOTU0NzQiLCI5NTQ3NSJdLCJjdXN0b21lcl9pZCI6IjIwMDEiLCJleHAiOjE2Mjg0NTE4MjYsInJvbGUiOiJ1c2VyIiwidXNlcm5hbWUiOiIyMDAxIn0.G--zBekvtAiLij_9pXoWL9MEW2wqXRu3RBL98dCVXTI]]
	u.RawQuery = q.Encode()
	log.Printf("u.RawQuery looks like: %v", u.RawQuery) // customer_id=2001&routeName=GetCustomer&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6WyI5NTQ3MiIsIjk1NDczIiwiOTU0NzQiLCI5NTQ3NSJdLCJjdXN0b21lcl9pZCI6IjIwMDEiLCJleHAiOjE2Mjg0NTE4MjYsInJvbGUiOiJ1c2VyIiwidXNlcm5hbWUiOiIyMDAxIn0.G--zBekvtAiLij_9pXoWL9MEW2wqXRu3RBL98dCVXTI
	log.Printf("u.String() looks like: %v", u.String()) // http://:8181/auth/verify?customer_id=2001&routeName=GetCustomer&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6WyI5NTQ3MiIsIjk1NDczIiwiOTU0NzQiLCI5NTQ3NSJdLCJjdXN0b21lcl9pZCI6IjIwMDEiLCJleHAiOjE2Mjg0NTE4MjYsInJvbGUiOiJ1c2VyIiwidXNlcm5hbWUiOiIyMDAxIn0.G--zBekvtAiLij_9pXoWL9MEW2wqXRu3RBL98dCVXTI
	return u.String()

}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
