
package state

type State int

const (
    Start State = iota
    Menu State = iota
    Play State = iota
    Pause State = iota
    GameOver State = iota
)
