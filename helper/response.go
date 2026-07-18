package helper

import (
    "encoding/json"
    "net/http"
)

type ResponseWithData struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    status := "success"
    if code >= 400 {
        status = "failed"
    }

    var res []byte
    var err error

    if payload != nil {
        res, err = json.Marshal(ResponseWithData{
            Status:  status,
            Message: message,
            Data:    payload,
        })
    } else {
        res, err = json.Marshal(ResponseWithoutData{
            Status:  status,
            Message: message,
        })
    }

    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Write(res)
}
