package model

const (
	Normal int = iota + 1
	Personal
)

const AESKey = "tuhgutkshhfkagbfjdjfarfh"
const (
	NoDeleted int = iota
	Deleted
)
const (
	NoArchive int = iota
	Archive
)
const (
	Open int = iota
	Private
	custom
)
const (
	Default = "default"
	Simple  = "simple"
)

const (
	NoCollected int = iota
	Collected
)
const (
	NoOwner int = iota
	Owner
)

const (
	NoExecutor = iota
	Executor
)

const (
	NoCanRead = iota
	CanRead
)

const (
	NoComment = iota
	Comment
)
