package tushare

type TuShare struct {
	token string
}

func New(token string) *TuShare {
	return &TuShare{token}
}
