package yazawa

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/gojp/kana"
	"github.com/shogo82148/go-mecab"
)

// ToYazawa convert text to yazawanized sentence
func ToYazawa(text string) string {
	tagger, err := mecab.New(map[string]string{})
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	tagger.Parse("")

	node, err := tagger.ParseToNode(text)
	if err != nil {
		panic(err)
	}

	id := findSuitableIndexForReplace(node)
	convertedSentence := []string{}
	for ; node != (mecab.Node{}); node = node.Next() {
		if node.Id() == id {
			feature := strings.Split(node.Feature(), ",")
			convertedSentence = append(convertedSentence, "『"+strings.ToUpper(kana.KanaToRomaji(feature[len(feature)-2]))+"』")
		} else {
			convertedSentence = append(convertedSentence, node.Surface())
		}
	}

	return strings.Join(convertedSentence, "")
}

func findSuitableIndexForReplace(node mecab.Node) int {
	idForReplace := 0
	maxScore := 0
	for ; node != (mecab.Node{}); node = node.Next() {
		score := 0
		word := node.Surface()
		score += examineWord(word)
		partsOfSpeech := strings.Split(node.Feature(), ",")[0]
		if partsOfSpeech == "形容詞" {
			score += 20
		} else if partsOfSpeech == "名詞" {
			score += 10
		} else if partsOfSpeech == "動詞" {
			score += 8
		}
		if score > maxScore {
			maxScore = score
			idForReplace = node.Id()
		}
	}
	return idForReplace
}

func examineWord(word string) int {
	score := 0
	hasKana := false
	hasKanji := false
	for _, r := range word {
		if !hasKana && kana.IsKana(fmt.Sprintf("%c", r)) {
			hasKana = true
		}
		if !hasKanji && kana.IsKanji(fmt.Sprintf("%c", r)) {
			hasKanji = true
		}
	}
	if hasKana {
		score += 10
	}
	if hasKanji {
		score += 10
	}
	if hasKana || hasKanji {
		score += 100
	}
	score += utf8.RuneCountInString(word)
	return score
}
