package decoder

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoder(t *testing.T) {
	payload, err := base64.StdEncoding.DecodeString("SAHkwD45wbQ/jbR0Qc5mZkGSiIA/MjQAAAHLAIAAA+K0AAIAAw==")
	assert.Nil(t, err)

	data, err := DecodePayload(payload)
	assert.Nil(t, err)

	assert.Equal(t, float32(133011.000000), data.Timestamp)
	assert.Equal(t, float32(0.18140298), data.Long)
	assert.Equal(t, float32(1.1070695), data.Lat)
	assert.Equal(t, float32(25.799999), data.Altitude)
	assert.Equal(t, float32(18.316650), data.RelativeHumidity)
	assert.Equal(t, float32(0.69610596), data.Temperature)
	assert.Equal(t, uint8(0), data.Status)
	assert.Equal(t, uint16(459), data.CO2PPM)
	assert.Equal(t, uint16(128), data.TVOCPPB)
	assert.Equal(t, uint16(2), data.PM10)
	assert.Equal(t, uint16(3), data.PM25)
}
