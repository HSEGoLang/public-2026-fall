module example.com/hello

go 1.27.0

require (
	example.com/moreints v0.0.0
	example.com/morestrings v0.0.0
)

require github.com/google/go-cmp v0.6.0 // indirect

replace example.com/moreints => ../moreints

replace example.com/morestrings => ../morestrings
