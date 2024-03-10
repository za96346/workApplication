package enum

type OperationItemCode string

const (
	Edit OperationItemCode = "edit"
	Inquire = "inquire"
	Add = "add"
	Delete = "delete"
	ChangeBanch = "changeBanch"
	Copy = "copy"
	Print = "print"
)