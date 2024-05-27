module github.com/submodule-org/submodule.go/batteries/sub_cmd

go 1.22.2

replace (
	github.com/submodule-org/submodule.go => ../..
	github.com/submodule-org/submodule.go/batteries/sub_env => ../sub_env
)

require (
	github.com/submodule-org/submodule.go v1.7.0
	github.com/urfave/cli/v2 v2.27.2
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
)
