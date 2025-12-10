module power4web

go 1.25

replace (
	power4web/game => ./game
	power4web/handlers => ./handlers
)

require (
	power4web/game v0.0.0-00010101000000-000000000000
	power4web/handlers v0.0.0-00010101000000-000000000000
)
