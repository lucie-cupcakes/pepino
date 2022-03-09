module github.com/lucie-cupcakes/pepino/service

replace github.com/lucie-cupcakes/pepino/engine => ../engine

replace github.com/lucie-cupcakes/pepino/untar => ../untar

require (
	github.com/google/uuid v1.3.0
	github.com/lucie-cupcakes/pepino/engine v0.0.0
	github.com/lucie-cupcakes/pepino/untar v0.0.0
)

go 1.17
