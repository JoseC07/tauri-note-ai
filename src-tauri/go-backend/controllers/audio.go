package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gordonklaus/portaudio"
)

func ListAudioDevices(c *gin.Context) {
	devices, err := portaudio.Devices()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list audio devices"})
		return
	}

	deviceList := make([]map[string]interface{}, len(devices))
	for i, device := range devices {
		deviceList[i] = map[string]interface{}{
			"name":              device.Name,
			"maxInputChannels":  device.MaxInputChannels,
			"maxOutputChannels": device.MaxOutputChannels,
			"defaultSampleRate": device.DefaultSampleRate,
		}
	}

	c.JSON(200, deviceList)
}

func ProcessAudioWithPython(c *gin.Context) {
	// Your audio processing logic here
	c.JSON(200, gin.H{"message": "Audio processing endpoint"})
} 