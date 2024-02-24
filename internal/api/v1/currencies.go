package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hukuma-finance/internal/model"
)

// @Summary Get all currencies
// @Produce  json
// @Success 200 {object} model.Result{data=[]model.Curency} "desc"
// @Router /api/v1/currencies/ [get]
func GetArticle(c *gin.Context) {
	currenies := []model.Curency{
		{
			Tiker: "USD",
			Logo: "unknown",
		}, 
		{
			Tiker: "EUR",
			Logo: "unknown",
		}, 
		{
			Tiker: "RUB",
			Logo: "unknown",
		}, 
	}
		c.JSON(http.StatusOK, model.Result{
			Data: currenies,
		})
}
