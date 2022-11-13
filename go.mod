module simulatorv2

go 1.19

require (
	client v0.0.0
	datanode v0.0.0
	namenode v0.0.0
)

require shell v0.0.0 // indirect

replace (
	client => ./client
	datanode => ./datanode
	namenode => ./namenode
	shell => ./shell
)
