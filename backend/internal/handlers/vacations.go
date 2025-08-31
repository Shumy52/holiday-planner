package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/shumy52/holiday-planner/backend/internal/models"
	// This needs to be actually commited.
)

type VacHandler struct{ DB *sqlx.DB }

type CreateReq struct {
	Start string `json:"start"` // ISO date
	End   string `json:"end"`
}

func (h *VacHandler) ListMine(c *gin.Context) {
	kc := c.MustGet("claims").(map[string]any)
	sub := kc["sub"].(string) // keycloak user id
	// map keycloak id -> app user_id (simple demo: store keycloak id directly)
	var out []models.Vacation
	err := h.DB.Select(&out, `select * from vacations v
      join users u on u.id = v.user_id
      where u.keycloak_id=$1 order by start_date`, sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *VacHandler) Create(c *gin.Context) {
	kc := c.MustGet("claims").(map[string]any)
	sub := kc["sub"].(string)

	var uID string
	if err := h.DB.Get(&uID, `select id from users where keycloak_id=$1`, sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not onboarded"})
		return
	}

	var req CreateReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "bad json"})
		return
	}
	s, _ := time.Parse("2006-01-02", req.Start)
	e, _ := time.Parse("2006-01-02", req.End)
	if !e.After(s) {
		c.JSON(400, gin.H{"error": "interval invalid"})
		return
	}
	days := int(e.Sub(s).Hours()/24) + 1

	var id string
	err := h.DB.Get(&id, `
    insert into vacations (user_id,start_date,end_date,total_days,status)
    values ($1,$2,$3,$4,'pending') returning id`, uID, s, e, days)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"id": id})
}
