package ambulance_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
  updateAmbulanceFunc(ctx, func(
    ctx *gin.Context,
    ambulance *Ambulance,
) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
    result := ambulance.PredefinedConditions
    if result == nil {
        result = []Condition{}
    }
    return nil, result, http.StatusOK
})
}