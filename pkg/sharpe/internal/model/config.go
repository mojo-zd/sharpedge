package model

type ConfigItems struct {
	Items []*Config
}

type Config struct {
	Name    string
	Content string
	Id      int64
}
