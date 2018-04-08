package client

type Client struct {
	id    int32 `db:"ID"`
	name  string
	email string
	phone string
}
