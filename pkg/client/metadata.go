package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

const (
	token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlBVOGRWQ2NzLWxDcXdlS0ttTGF4RSJ9.eyJpc3MiOiJodHRwczovL3ZtMDAxLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJnb29nbGUtb2F1dGgyfDExMjQyMzg4MTEwMjUwMzk5NTY1MCIsImF1ZCI6WyJhcGlzZXJ2ZXIiLCJodHRwczovL3ZtMDAxLnVzLmF1dGgwLmNvbS91c2VyaW5mbyJdLCJpYXQiOjE2MTM1OTQ2NDgsImV4cCI6MTYxMzY4MTA0OCwiYXpwIjoiT2NEZ0lIeVZvaUczOWhhRGI4Vm04OGRUYTk2bWdHVGciLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIn0.NNR3b6uikgwKJtVDB6cuptR0HEvHbnXBz2Uo4HMXZMV0hlnrJPn2Tafm-AvK382FHonQ2shVSrn9uk_l5iaKpvl9Zov0MRnp0UVxDuHGkSj7OQ2jJxBYyRHWV5pUhWL5F7mvu8vkhYOlmidSv1fj5yqJpUNTLcrNAT-gxLs75WNWa-urkeLsAE81hdydIuYL6ZYrOSy6lIFi0QLTv5OqAMBmyfyDMNDLZmedwSyNUhon7DfiiuclW8ZPZm7bJKw_kDy2ip7eyjLyl0u5LDS1HLalrYcoWuoY5PVhWC6R5O0DRMf4DyuE97bbR4fxrgfBiKBq-QSVRWfw6CaLc-g3iA"
)

func Credential() grpc.DialOption {
	cre, err := oauth.NewJWTAccessFromKey([]byte(token))
	if err != nil {
		panic(err)
	}

	return grpc.WithPerRPCCredentials(cre)
}
