package geofence

import "math"

func Check(lat, lon float64) bool {

	centerLat := -6.2088
	centerLon := 106.8456

	const R = 6371000

	dLat := (centerLat - lat) * math.Pi / 180
	dLon := (centerLon - lon) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat*math.Pi/180)*math.Cos(centerLat*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R*c <= 50
}
