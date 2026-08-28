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

func (cfg *apiConfig) getRecordsHandler(w http.ResponseWriter, r *http.Request) {
	type Record struct {
		ID           int32  `json:"id"`
		Record_Date  string `json:"record_date"`
		Record_Title string `json:"record_title"`
		Record_Type  string `json:"record_type"`
	}

	var allRecords []Record

	//query := r.URL.Query()
	//date := query.Get("record_date")
	//sortType := query.Get("sort")

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

	/* if sortType == "desc" {
		sort.Slice(allRecords, func(i, j int) bool {
			return allRecords[i].Record_Date.After(allRecords[j].CreatedAt)
		})
	} else {
		sort.Slice(allChirps, func(i, j int) bool {
			return allChirps[i].CreatedAt.Before(allChirps[j].CreatedAt)
		})
	} */

	respondWithJSON(w, 200, allRecords)

	/* if authorID != "" {
		parsedUUID, err := uuid.Parse(authorID)
		if err != nil {
			msg := fmt.Sprintf("Error parsing UUID: %s", err)
			respondWithError(w, 400, msg)
			return
		}
		chirps, err := cfg.db.GetChirpsByAuthorID(r.Context(), parsedUUID)
		if err != nil {
			msg := fmt.Sprintf("Error getting all chirps for author: %s", err)
			respondWithError(w, 500, msg)
			return
		}

		for _, chirp := range chirps {
			convertChirp := Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				User_ID:   chirp.UserID,
			}
			allChirps = append(allChirps, convertChirp)
		}
	} else {
		chirps, err := cfg.db.GetAllChirps(r.Context())
		if err != nil {
			msg := fmt.Sprintf("Error getting all chirps: %s", err)
			respondWithError(w, 500, msg)
			return
		}

		for _, chirp := range chirps {
			convertChirp := Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				User_ID:   chirp.UserID,
			}
			allChirps = append(allChirps, convertChirp)
		}
	}

	if sortType == "desc" {
		sort.Slice(allChirps, func(i, j int) bool {
			return allChirps[i].CreatedAt.After(allChirps[j].CreatedAt)
		})
	} else {
		sort.Slice(allChirps, func(i, j int) bool {
			return allChirps[i].CreatedAt.Before(allChirps[j].CreatedAt)
		})
	}

	respondWithJSON(w, 200, allChirps) */
}
