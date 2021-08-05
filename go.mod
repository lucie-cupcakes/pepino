module github.com/lucie-cupcakes/pepino

replace github.com/lucie-cupcakes/pepino/engine => ./engine

replace github.com/lucie-cupcakes/pepino/service => ./service

replace github.com/lucie-cupcakes/pepino/httpservice => ./httpservice

require github.com/lucie-cupcakes/pepino/engine v0.0.0

require github.com/lucie-cupcakes/pepino/service v0.0.0

require github.com/lucie-cupcakes/pepino/httpservice v0.0.0

go 1.16
