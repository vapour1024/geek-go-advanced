package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// Yes
func getErrNorows() error {
	// original error
	return errors.New("sql: no rows in result set")
}
func main() {
	fmt.Println("week02:code")
	var ErrNoRows = getErrNorows()
	// print original error
	fmt.Printf("original error:\n%s\n", ErrNoRows)
	// stack & message
	errors.Wrap(ErrNoRows, "main error")
	// just message
	// errors.WithMessage(ErrNoRows, "main error")
	// just stack
	// errors.WithStack(ErrNoRows)
	fmt.Printf("wraped error:\n%+v", ErrNoRows)
}
