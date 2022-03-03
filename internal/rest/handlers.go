package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/MarySmirnova/pereval/internal/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//putDataHandler - добавить в базу данные, вернуть id записи (pereval).
func (wr *Worker) postDataHandler(w http.ResponseWriter, r *http.Request) {
	var pereval data.Pereval
	var images data.Images

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &images)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = data.Validate(&pereval, &images); err != nil {
		log.WithError(err).Warn(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := wr.storage.PutDataToDB(&pereval, &images)
	if err != nil {
		log.WithError(err).Warn("unable to added data to DB") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	jsonID, err := json.Marshal(id)
	if err != nil {
		log.WithError(err).Warn("unable to parse response") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonID)
}

//getStatusHandler - получить статус модерации отправленных данных.
func (wr *Worker) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	status, err := wr.storage.GetStatusFromDB(id)
	if err != nil {
		log.WithError(err).Warn("incorrect parameter id") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonStatus, err := json.Marshal(status)
	if err != nil {
		log.WithError(err).Warn("unable to parse response") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonStatus)
}

//changeDataHandler - отредактировать существующую запись (замена), если она в статусе new.
//Редактировать можно все поля, кроме ФИО, почта, телефон.
func (wr *Worker) changeDataHandler(w http.ResponseWriter, r *http.Request) {
	//	id, _ := strconv.Atoi(mux.Vars(r)["id"])

}

//getAllDataHandler - список всех данных для отображения, которые этот пользователь отправил на сервер
//через приложение с возможностью фильтрации по данным пользователя (ФИО, телефон, почта), если передан объект.
func (wr *Worker) getAllDataHandler(w http.ResponseWriter, r *http.Request) {

}

//getDataHandler - получить одну запись (перевал) по её id.
func (wr *Worker) getDataHandler(w http.ResponseWriter, r *http.Request) {
	//	id, _ := strconv.Atoi(mux.Vars(r)["id"])

}
