package platform

import (
	"github.com/charmingruby/clowork/internal/platform/http/endpoint"
	"github.com/charmingruby/clowork/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *postgres.Client) {
	endpoint.New(r, db)
}
