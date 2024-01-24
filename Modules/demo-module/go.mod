module github/module-repo

go 1.21.5

require (
	github.com/deviant101/crypto v0.0.0-00010101000000-000000000000
	github.com/pborman/uuid v1.2.1
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/deviant101/crypto => ../crypto   //using replace directive if the module that we want to use is also availabe locally
