
package styles

import (
    "github.com/alecthomas/chroma"
)

// Emacs style.
var Emacs = Register(chroma.NewStyle("emacs", chroma.StyleEntries{
    chroma.TextWhitespace: "#bbbbbb",
    chroma.Comment: "italic #008800",
    chroma.CommentPreproc: "noitalic",
    chroma.CommentSpecial: "noitalic bold",
    chroma.Keyword: "bold #AA22FF",
    chroma.KeywordPseudo: "nobold",
    chroma.KeywordType: "bold #00BB00",
    chroma.Operator: "#666666",
    chroma.OperatorWord: "bold #AA22FF",
    chroma.NameBuiltin: "#AA22FF",
    chroma.NameFunction: "#00A000",
    chroma.NameClass: "#0000FF",
    chroma.NameNamespace: "bold #0000FF",
    chroma.NameException: "bold #D2413A",
    chroma.NameVariable: "#B8860B",
    chroma.NameConstant: "#880000",
    chroma.NameLabel: "#A0A000",
    chroma.NameEntity: "bold #999999",
    chroma.NameAttribute: "#BB4444",
    chroma.NameTag: "bold #008000",
    chroma.NameDecorator: "#AA22FF",
    chroma.LiteralString: "#BB4444",
    chroma.LiteralStringDoc: "italic",
    chroma.LiteralStringInterpol: "bold #BB6688",
    chroma.LiteralStringEscape: "bold #BB6622",
    chroma.LiteralStringRegex: "#BB6688",
    chroma.LiteralStringSymbol: "#B8860B",
    chroma.LiteralStringOther: "#008000",
    chroma.LiteralNumber: "#666666",
    chroma.GenericHeading: "bold #000080",
    chroma.GenericSubheading: "bold #800080",
    chroma.GenericDeleted: "#A00000",
    chroma.GenericInserted: "#00A000",
    chroma.GenericError: "#FF0000",
    chroma.GenericEmph: "italic",
    chroma.GenericStrong: "bold",
    chroma.GenericPrompt: "bold #000080",
    chroma.GenericOutput: "#888",
    chroma.GenericTraceback: "#04D",
    chroma.Error: "border:#FF0000",
    chroma.Background: " bg:#f8f8f8",
}))

