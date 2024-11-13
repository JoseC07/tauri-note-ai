package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gordonklaus/portaudio"
	"net/http"
)

func SetupAudioRoutes(r *gin.RouterGroup) {
	audio := r.Group("/audio")
	{
		audio.GET("/devices", listAudioDevices)
		audio.POST("/record", handleRecording)
	}
}

func listAudioDevices(c *gin.Context) {
	devices, err := portaudio.Devices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list audio devices"})
		return
	}

	deviceList := make([]map[string]string, 0)
	for _, device := range devices {
		if device.MaxInputChannels > 0 {
			deviceList = append(deviceList, map[string]string{
				"id":    device.Name,
				"label": device.Name,
			})
		}
	}

	c.JSON(http.StatusOK, deviceList)
}

func handleRecording(c *gin.Context) {
	// Handle recording logic
	c.JSON(http.StatusOK, gin.H{"status": "recording processed"})
} 