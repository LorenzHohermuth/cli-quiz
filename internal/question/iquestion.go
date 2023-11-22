package question

type IQuestion interface {
	Ask()
	Check() bool
}