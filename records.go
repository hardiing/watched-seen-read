package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/google/uuid"
	//"github.com/hardiing/watched-seen-read/internal/auth"
	"github.com/hardiing/watched-seen-read/internal/database"
)

func (cfg *apiConfig) addRecordHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Record_Date  string `json:"record_date"`
		Record_Title string `json:"record_title"`
		Record_Type  string `json:"record_type"`
	}

	type returnVals struct {
		Record_Date  time.Time `json:"record_date"`
		Record_Title string    `json:"record_title"`
		Record_Type  string    `json:"record_type"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		msg := fmt.Sprintf("Error decoding parameters: %s", err)
		respondWithError(w, 500, msg)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", params.Record_Date)
	if err != nil {
		respondWithError(w, 400, "Invalid date format. Expected YYYY-MM-DD")
		return
	}

	/* 	token, err := auth.GetBearerToken(r.Header)
	   	if err != nil {
	   		msg := fmt.Sprintf("Error getting authentication header: %s", err)
	   		respondWithError(w, 500, msg)
	   		return
	   	}
	   	validatedUser, err := auth.ValidateJWT(token, cfg.secret)
	   	if err != nil {
	   		msg := fmt.Sprintf("Error validating JWT: %s", err)
	   		respondWithError(w, 401, msg)
	   		return
	   	} */

	query_params := database.CreateRecordParams{
		RecordDate:  parsedDate,
		RecordTitle: params.Record_Title,
		RecordType:  params.Record_Type,
	}

	record, err := cfg.db.CreateRecord(r.Context(), query_params)
	if err != nil {
		msg := fmt.Sprintf("Error creating chirp: %s", err)
		respondWithError(w, 500, msg)
		return
	}

	respondWithJSON(w, 201, returnVals{
		Record_Date:  record.RecordDate,
		Record_Title: record.RecordTitle,
		Record_Type:  record.RecordType,
	})
}
