package service

type Compare interface{}

func (c Compare) Daxiao(hash int64) int8 {
	if hash > 4 {
		return 1
	}
}
