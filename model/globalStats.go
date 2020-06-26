package model

// TODO: ちゃんと独自のtype作る
type GlobalStats struct {
	Solo  detailedStats
	Duo   detailedStats
	Squad detailedStats
}
