package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"search-service/config"
	"search-service/internal/dto"
	"search-service/internal/repository"
)

type (
	SearchService interface {
	}

	SearchServiceImpl struct {
		repo   repository.SearchRepository
		config config.Config
	}
)

func NewSearchService(repo repository.SearchRepository, config config.Config) SearchService {
	return &SearchServiceImpl{repo: repo, config: config}
}


func getInfoFromOtherServiceById(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	r := dto.ResponseError{}
	json.Unmarshal(body, &r)

	fmt.Println(r)

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf(r.Errors[0])
	}
	return resp.StatusCode, nil
}

func getInfoFromOtherServiceByIds(url string, ids []int64) (int, error) {
	body := dto.GetByIDsRequest{
		IDs: ids,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	r := struct {
		Errors  []string `json:"errors"`
		Total   int64    `json:"total"`
		Message string   `json:"message"`
	}{}
	json.Unmarshal(bodyResp, &r)

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf(r.Errors[0])
	}

	if r.Total != int64(len(ids)) {
		return http.StatusBadRequest, fmt.Errorf("invalid variant ids")
	}

	return resp.StatusCode, nil
}