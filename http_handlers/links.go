package http_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"time"
	"urlShortener/helper"
	"urlShortener/models"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		var body UrlPost
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		verifyRegex := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
		matched, regexpErr := regexp.Match(verifyRegex, []byte(body.Url))
		if regexpErr != nil || !matched {
			fmt.Println(regexpErr, matched)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: "Wrong URL format",
				Data:    nil,
			})
			return
		}

		existLink, _ := models.SearchLink(body.Url)

		if existLink != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Message: fmt.Sprintf("Url already exists: %s", body.Url),
				Data:    nil,
			})
			return
		}

		shortKey := helper.GenerateShortKey(8)

		var saveLink models.LinkSave
		saveLink.Short = shortKey
		saveLink.ExpiredAt = time.Now().AddDate(0, 0, 30).Format(time.RFC3339)
		saveLink.Url = body.Url
		saveLink.UserId = 1
		saveErr := models.SaveLink(saveLink)
		if saveErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Message: "not ok",
				Data:    nil,
			})
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(Response{
				Message: "ok",
				Data:    saveLink,
			})
			return
		}
	} else {
		links, err := models.UserLinks(1)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{
				Message: "not ok",
				Data:    nil,
			})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{
				Message: "ok",
				Data:    links,
			})
			return
		}
	}
}

func LinksShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	link, err := models.GetLink(vars["link"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "not ok",
			Data:    nil,
		})
		return
	} else {
		err := models.LinkAddCounter(*link)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Message: "not ok",
				Data:    nil,
			})
			return
		}
		http.Redirect(w, r, link.Url, http.StatusSeeOther)
		return
	}
}

func LinksDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	link, err := models.GetLink(vars["link"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "not ok",
			Data:    nil,
		})
		return
	} else {
		delErr := models.DeleteLink(*link)
		if delErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Message: "not ok",
				Data:    nil,
			})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{
				Message: "Link deleted",
				Data:    nil,
			})
			return
		}
	}
}

func LinksStatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	link, err := models.GetLink(vars["link"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Response{
			Message: "not ok",
			Data:    nil,
		})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{
			Message: "ok",
			Data:    link,
		})
		return
	}
}
