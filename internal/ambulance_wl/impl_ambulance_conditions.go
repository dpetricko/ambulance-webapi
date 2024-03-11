package ambulance_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
  ctx.AbortWithStatus(http.StatusNotImplemented)
}