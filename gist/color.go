package gist

import (
	"bytes"
	"regexp"
	"strings"
	"time"
	"math/rand"
)

const(
	sep = "* * *"
)

func Color(contents []byte) (s []byte){
	
	pages := strings.Split(string(contents), sep)

	t := templates()
	var keys []string
	for k := range t {
		keys = append(keys,k)
	}

	design := designd(t,keys)
	var buffer bytes.Buffer	
	rand.Seed(time.Now().UnixNano())
	for i, p := range pages {
		// 見出し（Markdownの "====" ）を見つけたらデザインを変更する
		if m, _ := regexp.MatchString("={4,}",p); m {
			design = designd(t,keys)
		}

		if i != 0 {
			buffer.WriteString(sep)
		}
		buffer.WriteString(design)
		buffer.WriteString(p)
	}
	
	return []byte(buffer.String())
}

func designd(t map[string]string, keys []string) (s string){
	index := rand.Intn(len(keys))
	return t[keys[index]]	
}

func templates()(m map[string]string) {
	return map[string]string{
			"Stedelijk" : `
<!-- background: #ffffeb -->
<!-- color: #ff0000 -->
<!-- font: helvetica -->
`,
		"MOCAK" : `
<!-- background: #92117e -->
<!-- color: #ffd595 -->
<!-- font: verdana -->
`,
		"ReinaSofia" : `
<!-- background: #9bd1e7 -->
<!-- color: #72003c -->
<!-- font: tradegothic -->
`,
		"Pompidou" : `
<!-- background: #e4dadf -->
<!-- color: #774c43 -->
<!-- font: univers -->
`,
		"CCBB" : `
<!-- background: #f1f16d -->
<!-- color: #0d1c8b -->
<!-- font: frutiger -->
`,
		"SMAK" : `
<!-- background: #00acec -->
<!-- color: #fff -->
<!-- font: din -->
`,
		"LONDON" : `
<!-- background: #6e391b -->
<!-- color: #fff28c -->
<!-- font: centurygothic -->
`,
		"Oslo" : `
<!-- background: #50b187 -->
<!-- color: #fff -->
<!-- font: monaco -->
`,
		"Amsterdam" : `
<!-- background: red -->
<!-- color: #fff -->
<!-- font: frutiger -->
`,
		"HongKong" : `
<!-- background: #e9ca77 -->
<!-- color: #9f031e -->
<!-- font: futura -->
`,
		"Split" : `
<!-- background: #c8e4f6 -->
<!-- color: #15025e -->
<!-- font: brandon -->
`,
		"Marrakech" : `
<!-- background: #f8ebe5 -->
<!-- color: #a10318 -->
<!-- font: metronova -->
`,
		"SigmarPolke" : `
<!-- background: #14174a -->
<!-- color: #ffc8d9 -->
<!-- font: din -->
`,
		"DavidHockney" : `
<!-- background: #fffa28 -->
<!-- color: #25a9ce -->
<!-- font: centurygothic -->
`,
		"PabloPicasso" : `
<!-- background: #e75e05 -->
<!-- color: #ffd5fd -->
<!-- font: sabon -->
`,
		"SalvadorDali" : `
<!-- background: #ffb205 -->
<!-- color: #a10100 -->
<!-- font: metronova -->
`,
		"JacksonPollock" : `
<!-- background: #000100 -->
<!-- color: #feffd4 -->
<!-- font: rockwell -->
`,
		"BarbaraHepworth" : `
<!-- background: #6f6f6f -->
<!-- color: #fff -->
<!-- font: univers -->
`,
	}
}
