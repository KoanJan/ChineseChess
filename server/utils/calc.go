package utils

/*
Squared Euclidean Distance
*/
func SquaredEucDist(x1, y1, x2, y2 int32) int32 {

	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}
