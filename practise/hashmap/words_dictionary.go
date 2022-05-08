package hashmap

//hashmap.AddWords("english", "hello", "spanish", "ola")
//hashmap.AddWords("marathi", "namaskar", "english", "hello")
//hashmap.AddWords("english", "hi", "spanish", "oli")
//hashmap.GetWord("marathi", "ola")

var mapLanguage = make(map[string]map[string]string)

//Time O(N)
//Space O(words^2)

func AddWords(language1, word1, language2, word2 string) {
	_, isWord1Present := mapLanguage[word1]
	_, isWord2Present := mapLanguage[word2]

	//if both of the words are present
	if isWord1Present && isWord2Present {
		return
	}

	//if both of the words are not present
	if !isWord1Present && !isWord2Present {
		//Add word1 translation
		translationMapWord1 := make(map[string]string)
		translationMapWord1[language2] = word2
		mapLanguage[word1] = translationMapWord1
		//Add word2 translation
		translationMapWord2 := make(map[string]string)
		translationMapWord2[language1] = word1
		mapLanguage[word2] = translationMapWord2
	} else if isWord1Present && !isWord2Present {
		//if word1 is present
		word1Translations := mapLanguage[word1]
		word2Translations := make(map[string]string)
		for language, translation := range word1Translations {
			word2Translations[language] = translation
			dependentTranslations := mapLanguage[translation]
			dependentTranslations[language2] = word2
			mapLanguage[translation] = dependentTranslations
		}
		word2Translations[language1] = word1
		word1Translations[language2] = word2
		mapLanguage[word1] = word1Translations
		mapLanguage[word2] = word2Translations
	} else {
		//if word2 is present
		word2Translations := mapLanguage[word2]
		word1Translations := make(map[string]string)
		for language, translation := range word2Translations {
			word1Translations[language] = translation
			dependentTranslations := mapLanguage[translation]
			dependentTranslations[language1] = word1
			mapLanguage[translation] = dependentTranslations
		}
		word2Translations[language2] = word2
		word1Translations[language1] = word1
		mapLanguage[word2] = word1Translations
		mapLanguage[word1] = word2Translations
	}
}

//Time O(1)

func GetWord(language, word string) string {
	if _, isWordPresent := mapLanguage[word]; isWordPresent {
		translationsMap := mapLanguage[word]
		if _, isTranslationPresent := translationsMap[language]; isTranslationPresent {
			return translationsMap[language]
		}
	}
	return "Translation not available"
}
