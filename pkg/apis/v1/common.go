package v1

type CommonPaginationReq struct {
	Offset int `d:"0"              dc:"Offset number for pagination"`
	Limit  int `d:"10" v:"max:100" dc:"Limit size for pagination"`
}

type CommonPaginationRes struct {
	Total int `dc:"Total count number for pagination"`
}
