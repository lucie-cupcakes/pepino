module github.com/lucie-cupcakes/pepino/httpservice

replace github.com/lucie-cupcakes/pepino/service => ../service

replace github.com/lucie-cupcakes/pepino/engine => ../engine

require github.com/lucie-cupcakes/pepino/service v0.0.0

go 1.17
