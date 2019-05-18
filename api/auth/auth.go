package auth

import (
	"encoding/json"
	"fmt"
	"github.com/frayeralex/go-api/db"
	"github.com/frayeralex/go-api/middleware"
	. "github.com/frayeralex/go-api/models"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func getCollection() *bongo.Collection {
	return db.Connection.Collection(db.Users)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cred Credentials
	var user User
	_ = json.NewDecoder(r.Body).Decode(&cred)

	err := getCollection().FindOne(bson.M{
		"username": cred.Username,
	}, &user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"msg": "User not found",
		})
		return
	}

	if err := user.CheckPassword(cred.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"msg": "Invalid credentials",
		})
		return
	}

	token, err := middleware.GenerateJwt(string(user.GetId()), cred.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"id":       fmt.Sprintf(`%x`, string(user.GetId())),
		"username": user.Username,
		"token":    token,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cred Credentials
	var user User
	_ = json.NewDecoder(r.Body).Decode(&cred)

	err := getCollection().FindOne(bson.M{
		"username": cred.Username,
	}, &user)

	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		user = User{
			Username: cred.Username,
			Password: cred.Password,
		}
		err = getCollection().Save(&user)

		if _, ok := err.(*bongo.ValidationError); ok {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"msg": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(user)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]string{
		"msg": "User Already exists",
	})
}
