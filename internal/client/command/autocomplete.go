package command

import (
	"github.com/knz/bubbline/computil"
	"github.com/knz/bubbline/editline"
	"strings"
)

func Autocomplete(v [][]rune, line, col int) (msg string, completions editline.Completions) {
	firstLinePairs := strings.Split(string(v[0]), " ")
	if val, ok := commands[strings.TrimSpace(firstLinePairs[0])]; ok {
		return val.autocomplete()(v, line, col)
	}

	// Detect the word under the cursor.
	word, wstart, wend := computil.FindWord(v, line, col)

	var candidates []string
	for _, name := range commandNames {
		if strings.HasPrefix(name, word) {
			candidates = append(candidates, name)
		}
	}

	if len(candidates) == 0 {
		return msg, nil
	}

	return msg, editline.SimpleWordsCompletion(candidates, "commands", col, wstart, wend)
}
