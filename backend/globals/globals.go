package globals

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/maxkruse/magnusopus/backend/structs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	MAX_FILES_PER_USER = 5
)

var (
	Logger       *logrus.Logger
	Config       structs.Config
	DBConn       *gorm.DB
	SessionStore *session.Store

	AllowedSuperadmin []int
)
