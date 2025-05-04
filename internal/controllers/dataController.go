package controllers

import (
	"EffectiveMobileTest/internal/controllers/apiModels"
	"EffectiveMobileTest/internal/database"
	"EffectiveMobileTest/internal/database/models"
	"EffectiveMobileTest/internal/services"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// @Tags Data create
// @Summary create data endpoint
// @Accept json
// @Produce json
// Success 200
// @Router /data [post]
func DataCreate(w http.ResponseWriter, r *http.Request) {
	var p apiModels.DataCreateRequest

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		log.Error(err)
	}

	if p.Name == "" || p.Surname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbCon, dbErr := database.GetDatabaseConnection()

	if dbErr != nil {
		log.Error(dbErr)
		return
	}
	ageUrl := "https://api.agify.io/?name=" + p.Name
	sexUrl := "https://api.genderize.io/?name=" + p.Name
	nationUrl := "https://api.nationalize.io/?name=" + p.Name

	age := services.FetchAge(ageUrl)

	sex := services.FetchSex(sexUrl)
	nation := services.FetchNation(nationUrl)
	boolSex := false
	if sex.Gender == "male" {
		boolSex = true
	}

	data := models.DataModel{
		Name:       p.Name,
		Surname:    p.Surname,
		Patronymic: p.Patronymic,
		Age:        age.Age,
		Nation:     nation.Country[0].Country_id,
		Sex:        boolSex,
	}

	result := dbCon.Create(&data)

	if result.Error != nil {
		log.Error(result.Error)
	}
}

// @Tags Data create
// @Summary get data endpoint
// @Accept json
// @Produce json
// Success 200
// @Router /data [get]
func DataGet(w http.ResponseWriter, r *http.Request) {
	nameFilter := r.URL.Query().Get("nameFilter")

	dbCon, dbErr := database.GetDatabaseConnection()

	if dbErr != nil {
		log.Error(dbErr)
		return
	}
	var data []models.DataModel
	dbQuery := dbCon.Session(&gorm.Session{})
	if nameFilter != "" {
		dbQuery = dbQuery.Where("name LIKE ?", "%"+nameFilter+"%")
	}
	dbQuery.Scopes(paginate(r)).Find(&data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

// @Tags Data delete
// @Summary delete data endpoint
// @Accept json
// @Produce json
// Success 200
// @Router /data [delete]
func DataDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	dbCon, dbErr := database.GetDatabaseConnection()

	if dbErr != nil {
		log.Error(dbErr)
		return
	}

	dbCon.Delete(&models.DataModel{}, id)
}

// @Tags Data update
// @Summary update data endpoint
// @Accept json
// @Produce json
// Success 200
// @Router /data [put]
func DataUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var p apiModels.DataCreateRequest

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		log.Error(err)
	}

	if p.Name == "" || p.Surname == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbCon, dbErr := database.GetDatabaseConnection()

	if dbErr != nil {
		log.Error(dbErr)
		return
	}

	var data models.DataModel

	dbCon.Where("id = ?", id).First(&data)

	ageUrl := "https://api.agify.io/?name=" + p.Name
	sexUrl := "https://api.genderize.io/?name=" + p.Name
	nationUrl := "https://api.nationalize.io/?name=" + p.Name

	age := services.FetchAge(ageUrl)
	sex := services.FetchSex(sexUrl)
	nation := services.FetchNation(nationUrl)
	boolSex := false
	if sex.Gender == "male" {
		boolSex = true
	}

	data.Name = p.Name
	data.Surname = p.Surname
	data.Patronymic = p.Patronymic
	data.Age = age.Age
	data.Nation = nation.Country[0].Country_id
	data.Sex = boolSex

	dbCon.Save(&data)
}

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
