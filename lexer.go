package highlight

// l = goLexer{}
// for tok := range lexer {
//         colorChan <- tok
// }
type lexer struct {
        tokens chan token
}

