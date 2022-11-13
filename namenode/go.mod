module namenode

go 1.18
require (
	datanode v0.0.0
	shell v0.0.0
)

replace (
	datanode => ../datanode
	shell => ../shell
)