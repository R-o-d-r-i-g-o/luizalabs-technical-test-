package zipcode

import (
	"luizalabs-technical-test/internal/common"
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler struct holds a reference to the service layer.
type handler struct {
	svc ServiceImp
}

// NewHandler creates and returns a new handler instance with the injected service.
func NewHandler(svc ServiceImp) server.HandlerImp {
	return &handler{svc}
}

// Register sets up the route for retrieving CEP information.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/zip-code")
	g.GET("/:zip-code", h.getAddressByZipCode)
}

// getAddressByZipCode handles the request to retrieve CEP information.
func (h *handler) getAddressByZipCode(c *gin.Context) {
	zipCode := c.Param("zip-code")
	zipCode = common.StripNonNumericCharacters(zipCode)

	if isAccepted := common.ValidateZipCode(zipCode); !isAccepted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong format provided"})
		return
	}

	for {
		response, _ := h.svc.GetAddressByZipCode(zipCode)
		if response != nil {
			c.JSON(http.StatusOK, response)
			break
		}

		zipCode = common.AdjustLastNonZeroDigit(zipCode)
		if zipCode == common.EmptyZipCodeValue {
			c.JSON(http.StatusNotFound, gin.H{"error": "zip code not found"})
			break
		}
	}
}
