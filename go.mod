module github.com/lucie-cupcakes/pepino

replace github.com/lucie-cupcakes/pepino/engine => ./engine

replace github.com/lucie-cupcakes/pepino/service => ./service

replace github.com/lucie-cupcakes/pepino/httpservice => ./httpservice

replace github.com/lucie-cupcakes/pepino/untar => ./untar

require github.com/lucie-cupcakes/pepino/engine v0.0.0 // indirect

require github.com/lucie-cupcakes/pepino/service v0.0.0 // indirect

require github.com/lucie-cupcakes/pepino/httpservice v0.0.0

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/lucie-cupcakes/pepino/untar v0.0.0 // indirect
)

go 1.17
