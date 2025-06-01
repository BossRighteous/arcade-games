package arcadegame

type SensorType = uint8

type Sensor struct {
	XMin int32
	XMax int32
	YMin int32
	YMax int32
	Type SensorType
}

func OffsetSensor(s Sensor, xf float32, yf float32) Sensor {
	x := int32(xf)
	y := int32(yf)
	return Sensor{
		XMin: x + s.XMin,
		XMax: x + s.XMax,
		YMin: y + s.YMin,
		YMax: y + s.YMax,
	}
}

func TestIntersect(a Sensor, b Sensor) bool {
	return a.XMin <= b.XMax &&
		a.XMax >= b.XMin &&
		a.YMin <= b.YMax &&
		a.YMax >= b.YMin
}
