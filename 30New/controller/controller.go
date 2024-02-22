package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type UseCase interface {
	ContCreate(string, int) string
	ContDelete(int) (string, bool)
	ContGetFriends(int) (string, bool)
	ContAddFriend(int, int) (string, bool)
	ContUpdateAge(int, int) (string, bool)
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller { return &Controller{uc: uc} }

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var id int
		if err := json.Unmarshal(content, &id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, status := c.uc.ContDelete(id)

		jsonResult, _ := json.Marshal(result)

		if !status {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}

	}

	w.WriteHeader(http.StatusBadRequest)
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u user

		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result := c.uc.ContCreate(u.Name, u.Age)

		jsonResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)

		return
	}

	w.WriteHeader(http.StatusBadRequest)

}

func (c *Controller) UserFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var idUser int
		if err := json.Unmarshal(content, &idUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, status := c.uc.ContGetFriends(idUser)

		jsonResult, _ := json.Marshal(result)

		if !status {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)
}

type FriendsID struct {
	Source_id int `json:"source_id"`
	Target_id int `json:"target_id"`
}

func (c *Controller) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var friends FriendsID

		if err := json.Unmarshal(content, &friends); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, status := c.uc.ContAddFriend(friends.Source_id, friends.Target_id)

		jsonResult, _ := json.Marshal(result)

		if !status {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

type ElementUser struct {
	IDUser int `json:"iduser"`
	NewAge int `json:"newage"`
}

func (c *Controller) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var elements ElementUser

		if err := json.Unmarshal(content, &elements); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, status := c.uc.ContUpdateAge(elements.IDUser, elements.NewAge)

		jsonResult, _ := json.Marshal(result)

		if !status {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResult)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResult)
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)
}

func Server(c *Controller) {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", c.Create)
	mux.HandleFunc("/get", c.UserFriends)
	mux.HandleFunc("/friends", c.MakeFriends)
	mux.HandleFunc("/delete", c.Delete)
	mux.HandleFunc("/update", c.UpdateAge)
	http.ListenAndServe("localhost:8080", mux)
}
