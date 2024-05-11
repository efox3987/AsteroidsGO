
package state

type State int

const (
    Start State = iota
    Play State = iota
    GameOver State = iota
)
