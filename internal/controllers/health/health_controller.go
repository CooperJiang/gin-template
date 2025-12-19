package health

import (
	"net/http"
	"runtime"
	"template/pkg/cache"
	"template/pkg/database"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

type HealthResponse struct {
	Status    string                 `json:"status"`
	Timestamp string                 `json:"timestamp"`
	Uptime    string                 `json:"uptime"`
	Version   string                 `json:"version"`
	Checks    map[string]CheckResult `json:"checks"`
}

type CheckResult struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Latency string `json:"latency,omitempty"`
}

func (h *HealthController) Health(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
		Version:   "1.0.0",
		Checks:    make(map[string]CheckResult),
	}

	dbCheck := checkDatabase()
	response.Checks["database"] = dbCheck
	if dbCheck.Status == "error" {
		response.Status = "unhealthy"
	} else if dbCheck.Status == "degraded" {
		if response.Status != "unhealthy" {
			response.Status = "degraded"
		}
	}

	cacheCheck := checkCache()
	response.Checks["cache"] = cacheCheck
	if cacheCheck.Status == "error" {
		if response.Status == "healthy" {
			response.Status = "degraded"
		}
	}

	memoryCheck := checkMemory()
	response.Checks["memory"] = memoryCheck

	statusCode := http.StatusOK
	if response.Status == "unhealthy" {
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, response)
}

func (h *HealthController) Readiness(c *gin.Context) {
	checks := make(map[string]bool)

	db := database.GetDB()
	if db != nil {
		sqlDB, err := db.DB()
		if err == nil {
			if err := sqlDB.Ping(); err == nil {
				checks["database"] = true
			} else {
				checks["database"] = false
			}
		} else {
			checks["database"] = false
		}
	} else {
		checks["database"] = false
	}

	ready := true
	for _, status := range checks {
		if !status {
			ready = false
			break
		}
	}

	if ready {
		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
			"checks": checks,
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "not ready",
			"checks": checks,
		})
	}
}

func (h *HealthController) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func checkDatabase() CheckResult {
	start := time.Now()
	db := database.GetDB()

	if db == nil {
		return CheckResult{
			Status:  "error",
			Message: "database not initialized",
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		return CheckResult{
			Status:  "error",
			Message: err.Error(),
		}
	}

	if err := sqlDB.Ping(); err != nil {
		return CheckResult{
			Status:  "error",
			Message: err.Error(),
		}
	}

	latency := time.Since(start)
	result := CheckResult{
		Status:  "ok",
		Latency: latency.String(),
	}

	if latency > 100*time.Millisecond {
		result.Status = "degraded"
		result.Message = "high latency"
	}

	return result
}

func checkCache() CheckResult {
	start := time.Now()
	cacheClient := cache.GetCache()

	if cacheClient == nil {
		return CheckResult{
			Status:  "degraded",
			Message: "using memory cache",
		}
	}

	testKey := "health_check_test"
	testValue := "ok"

	if err := cacheClient.Set(testKey, testValue, time.Second); err != nil {
		return CheckResult{
			Status:  "error",
			Message: err.Error(),
		}
	}

	if _, err := cacheClient.Get(testKey); err != nil {
		return CheckResult{
			Status:  "error",
			Message: err.Error(),
		}
	}

	latency := time.Since(start)
	return CheckResult{
		Status:  "ok",
		Latency: latency.String(),
	}
}

func checkMemory() CheckResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	allocMB := m.Alloc / 1024 / 1024

	result := CheckResult{
		Status: "ok",
	}

	if allocMB > 500 {
		result.Status = "degraded"
		result.Message = "high memory usage"
	}

	return result
}
