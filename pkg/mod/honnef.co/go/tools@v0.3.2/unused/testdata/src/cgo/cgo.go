package pkg

//go:cgo_export_dynamic
func foo() {} //@ used(true)

func bar() {} //@ used(false)
