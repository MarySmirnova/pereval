package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/MarySmirnova/pereval/internal/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// @Summary Post New Pereval
// @Tags Pereval API
// @Description post new entry. Valid e-mail, coordinates and url values must be entered
// @ID post_data
// @Accept  json
// @Produce  json
// @Param input body data.Pereval true "pereval info"
// @Success 200 {integer} integer 1
// @Router /submitData [post]
func (wr *Worker) postDataHandler(w http.ResponseWriter, r *http.Request) {
	var pereval data.Pereval

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = data.Validate(&pereval); err != nil {
		log.WithError(err).Warn(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := wr.storage.PutDataToDB(body)
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

// @Summary Get Pereval Status
// @Tags Pereval API
// @Description returns the status of the entry
// @ID get_status
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} string 1
// @Router /submitData/{id}/status [get]
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

// @Summary Change Pereval
// @Tags Pereval API
// @Description all fields can be edited except email, phone, full name
// @ID update_data
// @Accept  json
// @Param id path int true "id"
// @Param input body data.Pereval true "pereval info"
// @Success 200
// @Router /submitData/{id} [put]
func (wr *Worker) changeDataHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var pereval data.Pereval

	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = data.Validate(&pereval); err != nil {
		log.WithError(err).Warn(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = wr.storage.UpdateDataToDB(id, body)
	if err != nil {
		log.WithError(err).Warn("unable to update data to DB") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Get All Data From User
// @Tags Pereval API
// @Description returns all user records (at least one parameter is required)
// @ID get_all_data
// @Produce  json
// @Param email query string false "email"
// @Param phone query string false "phone"
// @Param fam query string false "fam"
// @Param name query string false "name"
// @Param otc query string false "otc"
// @Success 200 {object} data.AllPereval
// @Router /submitData/ [get]
func (wr *Worker) getAllDataHandler(w http.ResponseWriter, r *http.Request) {
	paramsKeys := []string{"email", "phone", "fam", "name", "otc"}
	userParams := make(map[string]string)

	for _, key := range paramsKeys {
		if r.FormValue(key) != "" {
			userParams[key] = r.FormValue(key)
		}
	}

	if len(userParams) == 0 {
		err := fmt.Errorf("parameter not passed")
		log.WithError(err).Warn("incorrect parameter id") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	selectPereval, err := wr.storage.GetAllDataFromDB(userParams)
	if err != nil {
		log.WithError(err).Warn("failed to get data") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	perevalJson, err := json.Marshal(selectPereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse response") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(perevalJson)
}

// @Summary Get Pereval
// @Tags Pereval API
// @Description get a record by its id
// @ID get_data
// @Produce  json
// @Param id path integer true "id"
// @Success 200 {object} data.Pereval
// @Router /submitData/{id} [get]
func (wr *Worker) getDataHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	pereval, err := wr.storage.GetDataFromDB(id)
	if err != nil {
		log.WithError(err).Warn("incorrect parameter id") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	perevalJson, err := json.Marshal(pereval)
	if err != nil {
		log.WithError(err).Warn("unable to parse response") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(perevalJson)
}
