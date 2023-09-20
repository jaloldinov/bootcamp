package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"playground/cpp-bootcamp/models"
)

func (h *handler) PersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("person handler")
	switch r.Method {
	case http.MethodPost:
		fmt.Println("person handler create")
		h.CreatePerson(w, r)
	case http.MethodGet:
		fmt.Println("person handler get")

		if r.URL.Path[len("/person")] == '/' && r.URL.Path[len("/person")+1:] != "" {
			h.GetPerson(w, r)
		}
	case http.MethodPut:
		h.UpdatePerson(w, r)

	case http.MethodDelete:
		h.DeletePerson(w, r)
	}
}
func (h *handler) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.CreatePerson
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error ioutil.ReadAll:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("error Unmarshal:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	fmt.Println(h.storage)
	resp, err := h.storage.Person().Create(person)
	if err != nil {
		fmt.Println("error Person Create:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s", resp)))
}

func (h *handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	id := r.URL.Path[len("/person/"):]
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error ioutil.ReadAll:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		fmt.Println("error Unmarshal:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	person.Id = id
	fmt.Println(id)
	resp, err := h.storage.Person().Update(person)
	if err != nil {
		fmt.Println("error Person Update:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (h *handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("MethodGet")

	id := r.URL.Path[len("/person/"):]

	resp, err := h.storage.Person().Get(models.RequestByID{ID: id})
	if err != nil {
		fmt.Println("error Person Get:", err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error Marshal:", err.Error())
		w.Write([]byte("Internal server error"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)

}

func (h *handler) GetAllPersons(page, limit int) ([]models.Person, error) {
	resp, err := h.storage.Person().GetAll(models.GetAllRequest{Page: page, Limit: limit})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/person/"):]

	resp, err := h.storage.Person().Delete(models.RequestByID{ID: id})
	if err != nil {
		http.Error(w, "Failed to delete person", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(resp))
}
