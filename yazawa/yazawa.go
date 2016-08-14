package yazawa

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gojp/kana"
	"github.com/shogo82148/go-mecab"
)

// Convert convert text to yazawanized sentence
func Convert(text string, atRandom bool) string {
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

	id := findSuitableIndexForReplace(node, atRandom)
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

func findSuitableIndexForReplace(node mecab.Node, atRandom bool) int {
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

		if atRandom {
			rand.Seed(time.Now().UnixNano())
			score += rand.Intn(20)
		} else {
			score += utf8.RuneCountInString(word)
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
	hasKatakana := false
	hasKanji := false
	for _, r := range word {
		if !hasKatakana && kana.IsKatakana(fmt.Sprintf("%c", r)) {
			hasKatakana = true
		}
		if !hasKanji && kana.IsKanji(fmt.Sprintf("%c", r)) {
			hasKanji = true
		}
	}
	if hasKatakana {
		score += 10
	}
	if hasKanji {
		score += 10
	}
	if hasKatakana || hasKanji {
		score += 100
	}
	return score
}
