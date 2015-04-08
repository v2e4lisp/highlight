package highlight

type tokenType int

const (
        tUnkown tokenType = iota
        tVar
        tNum
        tConst
        tString
        // ...
)

type token struct {
        t tokenType
        v string
}
