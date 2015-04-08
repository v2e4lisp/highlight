package highlight

import "path"

type lang int

const (
        LANG_UNKNOWN lang = iota
        LANG_GO
)

func detect(filename string) Lang {
        ext := path.Ext(filename)
        lang := LANG_UNKNOWN

        switch ext {
        case "go":
                lang = LANG_GO
        }

        return lang
}
