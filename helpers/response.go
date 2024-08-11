package helpers

import (
  "encoding/json"
  "net/http"
)

type ResponseWithData struct {
  Status  string `json:"status"`
  Message string `json:"message"`
  Data    any    `json:"data"`
}

type ResponseWithoutData struct {
  Status  string `json:"status"`
  Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)

  var response any
  status := "Success"

  if code >= 400 {
    status = "Error"
  }

  if payload != nil {
    response = &ResponseWithData{
      Status:  status,
      Message: message,
      Data:    payload,
    }
  } else {
    response = &ResponseWithoutData{
      Status:  status,
      Message: message,
    }
  }

  res, _ := json.Marshal(response)
  _, err := w.Write(res)
  if err != nil {
    return
  }
}
