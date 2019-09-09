package decoder

import (
	"encoding/binary"
	"math"
)

type TKAQ struct {
	Timestamp        float32 // Ask Hans Jørgen
	Lat              float32 // Latitude from the GPS
	Long             float32 // Longitude from the GPS
	Altitude         float32 // Altitide from the GPS
	RelativeHumidity float32 // Relative humidity from the BNO080
	Temperature      float32 // Temperature from the BNO080
	Status           byte    // IAQ Core Status byte
	CO2PPM           uint16  // CO2 in PPM
	TVOCPPB          uint16  // TVOC in PPB
	PM25             uint16  // PM2.5 in ug/m3
	PM10             uint16  // PM10 in ug/m3
}

const (
	IAQStatusDataValid    = 0x00 // Valid data from IAQ Core
	IAQStatusSensorBusy   = 0x01 // IAQ Core is busy
	IAQStatusSensorWarmup = 0x10 // IAQ Core is in warm-up phase
	IAQStatusSensorError  = 0x80 // IAQ Core error
)

func decodeFloat(value []byte) float32 {
	bits := binary.BigEndian.Uint32(value)
	return math.Float32frombits(bits)
}

func decodeUint16(value []byte) uint16 {
	var b1 uint16
	var b2 uint16
	b1 = uint16(value[0])
	b2 = uint16(value[1])
	return b1<<8 | b2
}

func decodeUint32(value []byte) uint32 {
	var b1 uint32
	var b2 uint32
	var b3 uint32
	var b4 uint32
	b1 = uint32(value[0])
	b2 = uint32(value[1])
	b3 = uint32(value[0])
	b4 = uint32(value[1])
	return b1<<24 | b2<<16 | b3<<8 | b4
}

// DecodePayload decodes the payload from the Trondheim Kommune Air
// Quality sensor.
func DecodePayload(payload []byte) (*TKAQ, error) {
	return &TKAQ{
		Timestamp:        decodeFloat(payload[0:4]),
		Long:             decodeFloat(payload[4:8]),
		Lat:              decodeFloat(payload[8:12]),
		Altitude:         decodeFloat(payload[12:16]),
		RelativeHumidity: decodeFloat(payload[16:20]),
		Temperature:      decodeFloat(payload[20:24]),
		Status:           payload[24],
		CO2PPM:           decodeUint16(payload[25:27]),
		TVOCPPB:          decodeUint16(payload[27:29]),
		PM10:             decodeUint16(payload[33:35]),
		PM25:             decodeUint16(payload[35:37]),
	}, nil
}
