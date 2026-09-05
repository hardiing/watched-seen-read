package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		msg := fmt.Sprintf("Error decoding parameters: %s", err)
		respondWithError(w, 400, msg)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", params.Record_Date)
	if err != nil {
		respondWithError(w, 400, "Invalid date format. Expected YYYY-MM-DD")
		return
	}

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

func (cfg *apiConfig) getRecordsHandler(w http.ResponseWriter, r *http.Request) {
	type Record struct {
		ID           int32  `json:"id"`
		Record_Date  string `json:"record_date"`
		Record_Title string `json:"record_title"`
		Record_Type  string `json:"record_type"`
	}

	allRecords := make([]Record, 0)

	records, err := cfg.db.GetAllRecords(r.Context())
	if err != nil {
		msg := fmt.Sprintf("Error getting records: %s", err)
		respondWithError(w, 500, msg)
		return
	}

	for _, record := range records {
		convertRecord := Record{
			ID:           record.ID,
			Record_Date:  record.RecordDate.Format("2006-01-02"),
			Record_Title: record.RecordTitle,
			Record_Type:  record.RecordType,
		}
		allRecords = append(allRecords, convertRecord)
	}

	respondWithJSON(w, 200, allRecords)
}

func (cfg *apiConfig) deleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("recordID")
	val64, err := strconv.ParseInt(path, 10, 32)
	if err != nil {
		msg := fmt.Sprintf("Error parsing path: %s", err)
		respondWithError(w, 500, msg)
		return
	}
	val32 := int32(val64)
	err = cfg.db.DeleteRecord(r.Context(), val32)
	if err != nil {
		msg := fmt.Sprintf("Error deleting record: %s", err)
		respondWithError(w, 500, msg)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
