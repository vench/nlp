package nlp

import (
	"regexp"
	"strings"

	"github.com/james-bowman/sparse"
	"github.com/spaolacci/murmur3"
	"gonum.org/v1/gonum/mat"
	"io"
	"encoding/json"
	"errors"
)

// Vectoriser provides a common interface for vectorisers that take a variable
// set of string arguments and produce a numerical matrix of features.
type Vectoriser interface {
	Fit(...string) Vectoriser
	Transform(...string) (mat.Matrix, error)
	FitTransform(...string) (mat.Matrix, error)
}

// Transformer provides a common interface for transformer steps.
type Transformer interface {
	Fit(mat.Matrix) Transformer
	Transform(mat mat.Matrix) (mat.Matrix, error)
	FitTransform(mat mat.Matrix) (mat.Matrix, error)
}

var (
	stopWords = []string{"a", "about", "above", "above", "across", "after", "afterwards", "again", "against", "all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "amoungst", "amount", "an", "and", "another", "any", "anyhow", "anyone", "anything", "anyway", "anywhere", "are", "around", "as", "at", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "behind", "being", "below", "beside", "besides", "between", "beyond", "bill", "both", "bottom", "but", "by", "call", "can", "cannot", "cant", "co", "con", "could", "couldnt", "cry", "de", "describe", "detail", "do", "done", "down", "due", "during", "each", "eg", "eight", "either", "eleven", "else", "elsewhere", "empty", "enough", "etc", "even", "ever", "every", "everyone", "everything", "everywhere", "except", "few", "fifteen", "fify", "fill", "find", "fire", "first", "five", "for", "former", "formerly", "forty", "found", "four", "from", "front", "full", "further", "get", "give", "go", "had", "has", "hasnt", "have", "he", "hence", "her", "here", "hereafter", "hereby", "herein", "hereupon", "hers", "herself", "him", "himself", "his", "how", "however", "hundred", "ie", "if", "in", "inc", "indeed", "interest", "into", "is", "it", "its", "itself", "keep", "last", "latter", "latterly", "least", "less", "ltd", "made", "many", "may", "me", "meanwhile", "might", "mill", "mine", "more", "moreover", "most", "mostly", "move", "much", "must", "my", "myself", "name", "namely", "neither", "never", "nevertheless", "next", "nine", "no", "nobody", "none", "noone", "nor", "not", "nothing", "now", "nowhere", "of", "off", "often", "on", "once", "one", "only", "onto", "or", "other", "others", "otherwise", "our", "ours", "ourselves", "out", "over", "own", "part", "per", "perhaps", "please", "put", "rather", "re", "same", "see", "seem", "seemed", "seeming", "seems", "serious", "several", "she", "should", "show", "side", "since", "sincere", "six", "sixty", "so", "some", "somehow", "someone", "something", "sometime", "sometimes", "somewhere", "still", "such", "system", "take", "ten", "than", "that", "the", "their", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "therefore", "therein", "thereupon", "these", "they", "thickv", "thin", "third", "this", "those", "though", "three", "through", "throughout", "thru", "thus", "to", "together", "too", "top", "toward", "towards", "twelve", "twenty", "two", "un", "under", "until", "up", "upon", "us", "very", "via", "was", "we", "well", "were", "what", "whatever", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "whereupon", "wherever", "whether", "which", "while", "whither", "who", "whoever", "whole", "whom", "whose", "why", "will", "with", "within", "without", "would", "yet", "you", "your", "yours", "yourself", "yourselves"}
)

// Tokeniser interface for tokenisers allowing substitution of different
// tokenisation strategies e.g. Regexp and also supporting different
// different token types n-grams and languages.
type Tokeniser interface {
	// ForEachIn iterates over each token within text and invokes function
	// f with the token as parameter
	ForEachIn(text string, f func(token string))

	// Tokenise returns a slice of all the tokens contained in string
	// text
	Tokenise(text string) []string
}

// RegExpTokeniser implements Tokeniser interface using a basic RegExp
// pattern for unary-gram word tokeniser supporting optional stop word
// removal
type RegExpTokeniser struct {
	RegExp    *regexp.Regexp
	StopWords map[string]bool
}

// NewTokeniser returns a new, default Tokeniser implementation.  If
// removeStopwords is true then stop words will be removed from tokens
func NewTokeniser(removeStopwords bool) Tokeniser {
	var stop map[string]bool

	if removeStopwords {
		stop = make(map[string]bool)
		for _, word := range stopWords {
			stop[word] = true
		}
	}
	return &RegExpTokeniser{
		RegExp:    regexp.MustCompile("\\w+"),
		StopWords: stop,
	}
}

// ForEachIn iterates over each token within text and invokes function
// f with the token as parameter.  If StopWords is not nil then any
// tokens from text present in StopWords will be ignored.
func (t *RegExpTokeniser) ForEachIn(text string, f func(token string)) {
	tokens := t.tokenise(text)
	for _, token := range tokens {
		if t.StopWords != nil {
			if t.StopWords[token] {
				continue
			}
		}
		f(token)
	}
}

// Tokenise returns a slice of all the tokens contained in string
// text.  If StopWords is not nil then any tokens from text present in
// StopWords will be removed from the slice.
func (t *RegExpTokeniser) Tokenise(text string) []string {
	words := t.tokenise(text)

	// filter out stop words
	if t.StopWords != nil {
		b := words[:0]
		for _, w := range words {
			if !t.StopWords[w] {
				b = append(b, w)
			}
		}
		return b
	}

	return words
}

// tokenise returns a slice of all the tokens contained in string
// text.
func (t *RegExpTokeniser) tokenise(text string) []string {
	// convert content to lower case
	c := strings.ToLower(text)

	// match whole words, removing any punctuation/whitespace
	words := t.RegExp.FindAllString(c, -1)

	return words
}

// CountVectoriser can be used to encode one or more text documents into a term document
// matrix where each column represents a document within the corpus and each row represents
// a term present in the training data set.  Each element represents the frequency the
// corresponding term appears in the corresponding document e.g. tf(t, d) = 5 would mean
// that term t (perhaps the word "dog") appears 5 times in the document d.
type CountVectoriser struct {
	// Vocabulary is a map of words to indices that point to the row number representing
	// that word in the term document matrix output from the Transform() and FitTransform()
	// methods.  The Vocabulary map is populated by the Fit() or FitTransform() methods
	// based upon the words occurring in the datasets supplied to those methods.  Within
	// Transform(), any words found in the test data set that were not present in the
	// training data set supplied to Fit() will not have an entry in the Vocabulary
	// and will be ignored.
	Vocabulary map[string]int

	// Tokeniser is used to tokenise input text into features.
	Tokeniser Tokeniser
}

// NewCountVectoriser creates a new CountVectoriser.  If removeStopwords is true then
// english stop words will be removed.
func NewCountVectoriser(removeStopwords bool) *CountVectoriser {
	return &CountVectoriser{
		Vocabulary: make(map[string]int),
		Tokeniser:  NewTokeniser(removeStopwords),
	}
}

// Fit processes the supplied training data (a variable number of strings representing
// documents).  Each word appearing inside the training data will be added to the
// Vocabulary
func (v *CountVectoriser) Fit(train ...string) Vectoriser {
	i := 0
	for _, doc := range train {
		v.Tokeniser.ForEachIn(doc, func(word string) {
			_, exists := v.Vocabulary[word]
			if !exists {
				v.Vocabulary[word] = i
				i++
			}
		})
	}

	return v
}

// Transform transforms the supplied documents into a term document matrix where each
// column is a feature vector representing one of the supplied documents.  Each element
// represents the frequency with which the associated term for that row occurred within
// that document.  The returned matrix is a sparse matrix type.
func (v *CountVectoriser) Transform(docs ...string) (mat.Matrix, error) {
	mat := sparse.NewDOK(len(v.Vocabulary), len(docs))
	isEmpty := true

	for d, doc := range docs {
		v.Tokeniser.ForEachIn(doc, func(word string) {
			i, exists := v.Vocabulary[word]

			if exists {
				mat.Set(i, d, mat.At(i, d)+1)
				isEmpty = false
			}
		})
	}

	if isEmpty {
		return nil, errors.New("Text not exists in Vocabulary")
	}

	return mat.ToCSR(), nil
}

// FitTransform is exactly equivalent to calling Fit() followed by Transform() on the
// same matrix.  This is a convenience where separate training data is not being
// used to fit the model i.e. the model is fitted on the fly to the test data.
// The returned matrix is a sparse matrix type.
func (v *CountVectoriser) FitTransform(docs ...string) (mat.Matrix, error) {
	return v.Fit(docs...).Transform(docs...)
}

// HashingVectoriser can be used to encode one or more text documents into a term document
// matrix where each column represents a document within the corpus and each row represents
// a term.  Each element represents the frequency the corresponding term appears in the
// corresponding document e.g. tf(t, d) = 5 would mean that term t (perhaps the word "dog")
// appears 5 times in the document d.
type HashingVectoriser struct {
	NumFeatures int
	Tokeniser   Tokeniser
}

// NewHashingVectoriser creates a new HashingVectoriser.  If removeStopwords is true then
// english stop words will be removed.  numFeatures specifies the number of features
// that should be present in produced vectors.  Each word in a document is hashed and
// the mod of the hash and numFeatures gives the row in the matrix corresponding to that
// word.
func NewHashingVectoriser(removeStopwords bool, numFeatures int) *HashingVectoriser {
	return &HashingVectoriser{
		NumFeatures: numFeatures,
		Tokeniser:   NewTokeniser(removeStopwords),
	}
}

// Fit does nothing for a HashingVectoriser.  As the HashingVectoriser vectorises features
// based on their hash, it does require a pre-determined vocabulary to map features to their
// correct row in the vector.  It is effectively stateless and does not require fitting to
// training data.  The method is included for compatibility with other vectorisers.
func (v *HashingVectoriser) Fit(train ...string) Vectoriser {
	// The hashing vectoriser is stateless and does not require pre-training so this
	// method does nothing.
	return v
}

// Transform transforms the supplied documents into a term document matrix where each
// column is a feature vector representing one of the supplied documents.  Each element
// represents the frequency with which the associated term for that row occurred within
// that document.  The returned matrix is a sparse matrix type.
func (v *HashingVectoriser) Transform(docs ...string) (mat.Matrix, error) {
	mat := sparse.NewDOK(v.NumFeatures, len(docs))

	for d, doc := range docs {
		v.Tokeniser.ForEachIn(doc, func(word string) {
			h := murmur3.Sum32([]byte(word))
			i := int(h) % v.NumFeatures

			mat.Set(i, d, mat.At(i, d)+1)
		})
	}
	return mat.ToCSR(), nil
}

// FitTransform for a HashingVectoriser is exactly equivalent to calling
// Transform() with the same matrix.  For most vectorisers, Fit() must be called
// prior to Transform() and so this method is a convenience where separate
// training data is not used to fit the model.  For a HashingVectoriser, fitting is
// not required and so this method is exactly equivalent to Transform().  As with
// Fit(), this method is included with the HashingVectoriser for compatibility
// with other vectorisers.  The returned matrix is a sparse matrix type.
func (v *HashingVectoriser) FitTransform(docs ...string) (mat.Matrix, error) {
	return v.Transform(docs...)
}

// Pipeline is a mechanism for composing processing pipelines out of vectorisers
// transformation steps.  For example to compose a classic LSA/LSI pipeline
// (vectorisation -> TFIDF transformation -> Truncated SVD) one could use a
// Pipeline as follows:
// 	lsaPipeline := NewPipeline(NewCountVectoriser(false), NewTfidfTransformer(), NewTruncatedSVD(100))
//
type Pipeline struct {
	Vectoriser   Vectoriser
	Transformers []Transformer
}

// NewPipeline constructs a new processing pipline with the supplied Vectoriser
// and one or more transformers
func NewPipeline(vectoriser Vectoriser, transformers ...Transformer) *Pipeline {
	pipeline := Pipeline{
		Vectoriser:   vectoriser,
		Transformers: transformers,
	}

	return &pipeline
}

// Fit fits the model(s) to the supplied training data
func (p *Pipeline) Fit(docs ...string) Vectoriser {
	if _, err := p.FitTransform(docs...); err != nil {
		panic("nlp: Failed to Fit pipeline because " + err.Error())
	}

	return p
}

// Transform transforms the supplied documents into a matrix representation
// of numerical feature vectors using a model(s) previously fitted to supplied
// training data.
func (p *Pipeline) Transform(docs ...string) (mat.Matrix, error) {
	matrix, err := p.Vectoriser.Transform(docs...)
	if err != nil {
		return matrix, err
	}
	for _, t := range p.Transformers {
		matrix, err = t.Transform(matrix)
		if err != nil {
			return matrix, err
		}
	}
	return matrix, nil
}

// FitTransform transforms the supplied documents into a matrix representation
// of numerical feature vectors fitting the model to the supplied data in the
// process.
func (p *Pipeline) FitTransform(docs ...string) (mat.Matrix, error) {
	matrix, err := p.Vectoriser.FitTransform(docs...)
	if err != nil {
		return matrix, err
	}
	for _, t := range p.Transformers {
		matrix, err = t.FitTransform(matrix)
		if err != nil {
			return matrix, err
		}
	}
	return matrix, nil
}



//
func (c *CountVectoriser) Save(w io.Writer) error {
	buf, err := json.Marshal(c.Vocabulary)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	return err
}

//
func (c *CountVectoriser) Load(r io.Reader) error {
  	dec := json.NewDecoder(r)
	return dec.Decode( &c.Vocabulary)
}