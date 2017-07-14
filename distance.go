package distance

import (
    "math"
)

const earthRadius = 6378137

func rad(value float64) (float64) {
    return value * math.Pi / 180.0
}

func GetDistance(longitude, latitude, otherLongitude, otherLatitude float64) (float64) {
    radLatitude := rad(latitude)
    radOtherLatitude := rad(otherLatitude)

    a := radLatitude - radOtherLatitude
    b := rad(longitude) - rad(otherLongitude)

    return earthRadius * 2 * math.Asin(
        math.Sqrt(
            math.Pow(
                math.Sin(a / 2), 2) + math.Cos(radLatitude) * math.Cos(radOtherLatitude) * math.Pow(math.Sin(b / 2), 2)))
}
