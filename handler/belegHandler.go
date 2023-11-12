// Package handler beinhaltet die Implementierung des BelegHandlers.
// Der BelegHandler dient als Schnittstelle zwischen der Außenwelt (in diesem Fall HTTP-Anfragen)
// und der internen Anwendungslogik, die im BelegService implementiert ist.
//
// Dieser Handler verarbeitet eingehende HTTP-Anfragen, extrahiert und validiert die benötigten Daten,
// ruft die entsprechenden Methoden des BelegService auf und sendet die passenden HTTP-Antworten zurück.
// Er ist verantwortlich für:
// - Das Entgegennehmen von Anfragen und deren Weiterleitung an den BelegService.
// - Die Validierung von Anfrageparametern.
// - Das Transformieren von Daten aus dem Format, das der Service zurückgibt, in ein für den Client geeignetes Format.
// - Das Senden von HTTP-Antworten basierend auf den Ergebnissen der Service-Operationen.
//
// Der BelegHandler sollte keine Geschäftslogik oder direkte Datenbankinteraktionen enthalten.
// Diese Verantwortlichkeiten liegen in der Service- bzw. Repository-Schicht.
// Dadurch wird eine klare Trennung der Verantwortlichkeiten gemäß dem Domain-Driven Design (DDD) Prinzip erreicht.

package handler

import (
	"beleg-app/api/domain"
	"beleg-app/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BelegHandler struct {
	service *service.BelegService
}

func NewBelegHandler(service *service.BelegService) *BelegHandler {
	return &BelegHandler{service: service}
}

func (h *BelegHandler) GetBelegById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	beleg, err := h.service.GetBelegById(c.Request.Context(), id)
	if err != nil {
		// Hier können Sie die Fehlerbehandlung verfeinern, z.B. unterschiedliche Antworten je nach Art des Fehlers
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if beleg == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Beleg not found"})
		return
	}

	c.JSON(http.StatusOK, beleg)
}

func (h *BelegHandler) CreateBeleg(c *gin.Context) {
	var beleg domain.Beleg
	if err := c.ShouldBindJSON(&beleg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateBeleg(c.Request.Context(), &beleg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, beleg)
}

func (h *BelegHandler) GetAllBelege(c *gin.Context) {
	belege, err := h.service.GetAllBelege(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, belege)
}

func (h *BelegHandler) DeleteBelegById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.service.DeleteBelegById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Beleg deleted")

}
