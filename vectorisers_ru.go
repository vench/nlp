package nlp

import (
	"strings"
	"github.com/opennota/morph"
	"unicode"
)



var globalStopWords  = map[string]bool{
	"нее"	: true,	"этого": true,	"между": true,
	"несколько": true,	"момент": true,
	"какойто": true,	"было": true,
	"жестко": true,	"всей": true,
	"числе": true,	"двое": true,
	"про": true,	"стать": true,
	"быстро": true,	"маленькое": true,
	"итоге": true,	"нечего": true,
	"дать": true,	"впрочем": true,
	"отсюда": true,	"мои": true,
	"возле": true,	"немедленно": true,
	"небольшое":true,	"своем": true,
	"каково": true,	"одна": true,
	"крепкое":true,	"хочет": true,
	"чуть": true,	"признать":true,
	"буду": true,	"забыли": true,
	"сюда": true,	"против": true,
	"едва": true,	"начать": true,
	"виновато": true,	"нормальные": true,
	"скорее": true,	"срочное": true,
	"вышли": true,	"другое": true,
	"думаю": true,	"сквозь": true,
	"нем": true,	"величайшей": true,
	"побыстрее": true,	"испуганно": true,
	"однозначно": true,	"особо": true,
	"нервно": true,	"любым": true,
	"пускай": true,	"абсолютно": true,
	"сначала": true,	"многое": true,
	"мне": true,	"пусть": true,
	"случае": true,	"тот": true,
	"свое": true,	"стала": true,
	"прочь": true,	"вовсе": true,
	"позже": true,	"меня": true,
	"сами": true,	"третьего": true,
	"прошу": true,	"данное": true,
	"которого": true,	"им": true,
	"высоко": true,	"ваша": true,
	"когда": true,	"одной": true,
	"гдето": true,	"хоть": true,
	"мол": true,	"почти": true,
	"могла": true,	"низкие": true,
	"можно": true,	"любимое": true,
	"столько": true,	"хочешь": true,
	"тяжело": true,	"должен": true,
	"поводу": true,	"хотят": true,
	"какието": true,	"нас": true,
	"тоже": true,	"этот": true,
	"бывает": true,	"своего": true,
	"идет": true,	"серьезных": true,
	"некоторых": true,	"якобы": true,
	"прежде": true,	"крупные": true,
	"этих": true,	"минимум": true,
	"ведении": true,	"очевидно": true,
	"нам": true,	"сомнения": true,
	"существуют": true,	"иногда": true,
	"мало": true,	"необходимых": true,
	"целом": true,	"ваши": true,
	"иметь": true,	"вполне": true,
	"долго": true,	"лишь": true,
	"откуда": true,	"сколько": true,
	"безусловно": true,	"самое": true,
	"порой": true,	"такая": true,
	"разным": true,	"невозможно": true,
	"вероятно": true,	"легко": true,
	"лучше": true,	"примеру": true,
	"каждого": true,	"куда": true,
	"разной": true,	"весьма": true,
	"совсем": true,	"необходима": true,
	"тд": true,	"обычной": true,
	"нельзя": true,	"новый": true,
	"ею": true,	"ваше": true,
	"обойтись": true,
	"высокого": true,
	"лучшие": true,
	"непросто": true,
	"изза": true,
	"отметим": true,
	"третье": true,
	"втором": true,
	"ждать": true,
	"хотите": true,
	"трудно": true,
	"вообще": true,
	"никоим": true,
	"частью": true,
	"нужны": true,
	"типа": true,
	"негативно": true,
	"легкие": true,
	"каким": true,
	"действительно": true,
	"абсурд": true,
	"прошлое": true,
	"напомним": true,
	"быля": true,
	"своя": true,
	"та": true,
	"вы": true,
	"ли": true,
	"хватает": true,
	"тут": true,
	"таким": true,
	"некоторое": true,
	"тое": true,
	"вместе": true,
	"нормально": true,
	"часто": true,
	"недавно": true,
	"слишком": true,
	"некоторые": true,
	"принципе": true,
	"надо": true,
	"многие": true,	"возможность": true,
	"образом": true,	"одного": true,
	"значительно": true,
	"другие": true,	"относительно": true,
	"сделать": true,	"пытается": true,	"снова": true,	"важно": true,	"нуждаются": true,	"тех": true,	"самых": true,
	"необходимости": true,
	"далеко": true,	"два": true,	"деле": true,
	"через": true,	"нашего": true,	"которое": true,
	"каждое":true,
	"менее": true,
	"шесть": true,
	"много": true,
	"надеемся": true,
	"каждая": true,
	"конечно": true,
	"дали": true,
	"бывают": true,
	"чему": true,
	"сама": true,
	"другого": true,
	"самого": true,
	"далее": true,
	"также": true,
	"обязан": true,
	"кто": true,
	"уже": true,
	"ты": true,
	"должна": true,
	"наш": true,
	"при": true,
	"вот": true,
	"этом": true,
	"первую": true,
	"особые": true,
	"всего": true,
	"поэтому": true,
	"такие": true,
	"назад": true,
	"нашим": true,
	"ни": true,
	"немного": true,
	"ровно": true,
	"этому": true,
	"себе": true,
	"которые": true,
	"глядя": true,
	"под": true,
	"какоето": true,
	"начали": true,
	"одно": true,
	"введена": true,
	"где": true,
	"например": true,
	"свои": true,
	"каждый": true,
	"те": true,
	"то": true,
	"такой": true,
	"несет": true,
	"второй": true,
	"такое": true,
	"охотно": true,
	"три": true,
	"той": true,
	"находятся": true,
	"бывшую": true,
	"остальные": true,
	"ему": true,
	"отмечается": true,
	"нужно": true,
	"свою": true,
	"же": true,
	"четыре": true,
	"первому": true,
	"пояснил": true,
	"самом": true,
	"который": true,
	"всем": true,
	"быть": true,
	"со": true,
	"ничего": true,
	"могут": true,
	"чего": true,
	"затем": true,
	"около": true,
	"весь": true,
	"двух": true,
	"трех": true,
	"несут": true,
	//"объект:21"
	"восемь": true,
	"нет": true,
	"оттуда": true,
	"этим": true,
	"выше": true,
	"московского": true,
	"чем": true,
	"своей": true,
	"которому": true,
	"тем": true,
	"сейчас": true,
	"вас": true,
	"первой": true,
	"которая": true,
	"сможем": true,
	"важных": true,
	"начала": true,
	"принял": true,
	"эта": true,
	"которых": true,
	"я": true,
	"других": true,
	"после": true,
	"кроме": true,
	"некоторый": true,
	"ее": true,
	"непонятных": true,
	"остальных": true,
	"отметил": true,
	"ушел": true,
	"данная": true,
	"а": true,
	"да": true,
	"даже": true,
	"какихлибо": true,
	"каких": true,
	"либо": true,
	"новое": true,
	"такого": true,
	"том": true,
	"один": true,
	"почему": true,
	"сегодня": true,
	"своим": true,
	"насчет": true,
	"тебе": true,
	"тебя": true,
	"важна": true,
	"хотим": true,
	"частные": true,
	"мое": true,
	"очень": true,
	"какой": true,
	"какое": true,
	"никто": true,
	"кому": true,
	"чтобы": true,
	"большой": true,
	"плохо": true,
	"среди": true,
	"должно": true,
	"ждут": true,
	"бывшие": true,
	"чаще": true,
	"другой": true,
	"берут": true,
	"ведь": true,
	"пяти": true,
	"просто": true,
	"готов": true,
	"давно": true,
	"в": true,
	"всех": true,
	"стало": true,
	"может": true,
	"пор": true,
	"текущей": true,
	"принять": true,
	"или": true,
	"на": true,
	"есть": true,
	"еще": true,
	"первого": true,
	"четверки": true,
	"изменило": true,
	"причем": true,
	"широкой": true,
	"более": true,
	"тогда": true,
	"ранее": true,
	"стали": true,
	"нашей": true,
	"пониманием": true,
	"теперь": true,
	"другая": true,
	"будут": true,
	"считаю": true,
	"третьему": true,
	"раз": true,
	"для": true,
	"только": true,
	"итогам": true,
	"какие": true,
	"нескольких": true,
	"старый": true,
	"наших": true,
	"первый": true,
	"однако": true,
	"прочих": true,
	"того": true,
	"если": true,
	"о": true,
	"него": true,
	"известно": true,
	"прямо": true,
	"вновь" : true,
	"больше": true,
	"невелик": true,
	"именно": true,
	"из": true,
	"пока": true,
	"стал": true,
	"у":true,
	"данный": true,
	"зачастую": true,
	"сам": true,
	"за": true,
	"как":true,
	"во":true,
	"что":true,
	"и":true,
	"a":true,
	"об":true,
	"с":true,
	"по":true,
	"хочу":true,
	"к":true,
	"ей":true,
	"эти":true,
	"не":true,
	"потому":true,
	"так":true,
	"он":true,
	"она":true,
	"этой":true,
	"многих":true,
	"были":true,
	"вплоть":true,
	"самые":true,
	"мы":true,
	"должны":true,
	"но":true,
	"самый":true,
	"хотя":true,
	"здесь":true,
	"от":true,
	"до":true,
	"это":true,
	"них":true,
	"наши":true,
	"его":true,
	"свой":true,
	"своих":true,
	"сове":true,
	"таких":true,
	"все":true,
	"имеет":true,
	"был":true,
	"эх": true,
	"ээ": true,
	"ю": true,
	"тому":true,
	"любых":true,
	"любой":true,
	"там":true,
	"бы":true,
	"без":true,
	"их":true,
	"будет":true,
	"они": true,
	"онa": true,
	"оно": true,
}


//
type VsTokeniser struct {
	ModeWord bool
	RemoveStopwords bool
}

// ForEachIn iterates over each token within text and invokes function
// f with the token as parameter
func (t *VsTokeniser) ForEachIn(text string, f func(token string)) {
	tokens := t.tokenise(text)
	for _, token := range tokens {
		if globalStopWords != nil {
			if globalStopWords[token] {
				continue
			}
		}
		f(token)
	}
}


// tokenise returns a slice of all the tokens contained in string
// text.
func (t *VsTokeniser) tokenise(text string) []string {

	words := make([]string, 0)
	runes := []rune(text)
	buf := make([]rune, 0)

	for _,r := range runes {

		if unicode.IsSpace(r) && len(buf) > 0 {
			s := string(buf)
			words = append(words, s)
			buf = make([]rune, 0)
		}
		if r >= 1040 && r <= 1071 {
			r += 32
		}
		if r >= 1072 && r <= 1103  {
			buf = append(buf, r)
		}
	}

	if len(buf) > 0  {
		s := string(buf)
		words = append(words, s)
	}

	if t.ModeWord && len(words) > 0 {
		for i, _:= range words {
			words[i] = MorfModeWord(words[i])
		}
	}

	return words
}


var stemming = map[string]string {
	"ского": "ский", //москов-ский
	"скому": "ский",
	"ском": "ский",
	"тору": "тор",
	"тора": "тор",
	"тором": "тор",
	"торы": "тор",
	"торе": "тор",
	"дара":"дар",
	"дары":"дар",
	"дару":"дар",
	"даров":"дар",
	"даром":"дар",
	"дарам":"дар",
	"дарами":"дар",
	"дарах":"дар",
	"даре":"дар",
	"чику":"чик",
	"чика":"чик",
	"чиками":"чик",
	"чиков":"чик",
	"чиком":"чик",
	"зику":"зик",
	"зика":"зик",
	"зиками":"зик",
	"зиков":"зик",
	"зиком":"зик",

	"щика":"щик",
	"щиком":"щик",
	"щику":"щик",
	"щикам":"щик",
	"щиков":"щик",
	"утина": "утин",
}

//
func MorfModeWord(word string) string {

	abc := "ауоыиэяюёе"
	l := len(word)
	for n := 2; l > n && n < 14; n += 2 {
		s := word[l - n:]
		if v, ok := stemming[s]; ok && strings.ContainsAny(word[:l-n], abc) {
			return  word[:l-n] + v
		}
	}

	_, n, _ := morph.Parse(word)

	if len(n) > 0 {
		return n[0]
	}

	return word
}


// Tokenise returns a slice of all the tokens contained in string
// text.  If StopWords is not nil then any tokens from text present in
// StopWords will be removed from the slice.
func (t *VsTokeniser) Tokenise(text string) []string {
	words := t.tokenise(text)

	// filter out stop words
	if t.RemoveStopwords && globalStopWords != nil {
		b := words[:0]
		for _, w := range words {
			if !globalStopWords[w] {
				b = append(b, w)
			}
		}
		return b
	}

	return words
}

//
func VsNewTokeniser(removeStopwords bool) Tokeniser {
	return &VsTokeniser{
		ModeWord: true,
		RemoveStopwords: removeStopwords,
	}
}

//
func VsNewCountVectoriser(removeStopwords bool) *CountVectoriser {
	return &CountVectoriser{
		Vocabulary: make(map[string]int),
		Tokeniser:  VsNewTokeniser(removeStopwords),
	}
}
