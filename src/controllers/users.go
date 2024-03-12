package controllers

import "net/http"

func Create(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Create User"))
}

func FindUsers(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("buscar Users"))
}

func FindUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("buscar User"))
}

func Update(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("upa User"))
}

func Delete(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("deletar User"))
}
