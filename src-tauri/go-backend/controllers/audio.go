package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gordonklaus/portaudio"
	"os/exec"
)

// POTENTIAL ISSUE: Need to handle different audio formats
// POTENTIAL ISSUE: Need proper error handling for device access
func ListAudioDevices(c *gin.Context) {
	devices, err := portaudio.Devices()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list audio devices"})
		return
	}

	// Convert devices to a simpler format for JSON
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

// POTENTIAL ISSUE: Need to handle Python process lifecycle
func ProcessAudioWithPython(c *gin.Context) {
	// Execute Python script for AI processing
	cmd := exec.Command("python3", "notes_ai_helper.py", "process_audio")
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to process audio"})
		return
	}

	c.JSON(200, gin.H{"result": string(output)})
} 