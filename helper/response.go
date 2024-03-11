package helper

import "inventory-manajemen-system/dto"

type ResponsWithData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponsWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Respons(p dto.ResponsParams) any {
	var response any
	var status string

	if p.StatusCode >= 200 && p.StatusCode < 300 {
		status = "success"
	} else {
		status = "failed"
	}

	if p.Data != nil {
		response = &ResponsWithData{
			Code:    p.StatusCode,
			Status:  status,
			Message: p.Message,
			Data:    p.Data,
		}
	} else {
		response = &ResponsWithoutData{
			Code:    p.StatusCode,
			Status:  status,
			Message: p.Message,
		}
	}

	return response
}
