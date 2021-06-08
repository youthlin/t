package plurals

func i(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
func _if(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
}

// commons holds some commonly used expression from gnu site
// https://www.gnu.org/software/gettext/manual/html_node/Plural-forms.html#index-plural_002c-in-a-PO-file-header
var commons = map[string]func(n int64) int64{
	// Asian family: Japanese, Vietnamese, Korean
	// Tai-Kadai family: Thai
	"0": func(n int64) int64 { return 0 },
	// Germanic family: English, German, Dutch, Swedish, Danish, Norwegian, Faroese
	// Romanic family: Spanish, Portuguese, Italian
	// Latin/Greek family: Greek
	// Slavic family: Bulgarian
	// Finno-Ugric family: Finnish, Estonian
	// Semitic family: Hebrew
	// Austronesian family: Bahasa Indonesian
	// Artificial: Esperanto
	"n!=1": func(n int64) int64 { return i(n != 1) },
	// Romanic family: Brazilian Portuguese, French
	"n>1": func(n int64) int64 { return i(n > 1) },
	// Baltic family: Latvian
	"n%10==1&&n%100!=11?0:n!=0?1:2": func(n int64) int64 {
		return _if(n%10 == 1 && n%100 != 11, 0, _if(n != 0, 1, 2))
	},
	// Celtic: Gaeilge (Irish)
	"n==1?0:n==2?1:2": func(n int64) int64 { return _if(n == 1, 0, _if(n == 2, 1, 2)) },
	// Romanic family: Romanian
	"n==1?0:(n==0||(n%100>0&&n%100<20))?1:2": func(n int64) int64 {
		return _if(n == 1, 0, _if(n == 0 || (n%100 > 0 && n%100 < 20), 1, 2))
	},
	// Baltic family: Lithuanian
	"n%10==1&&n%100!=11?0:n%10>=2&&(n%100<10||n%100>=20)?1:2": func(n int64) int64 {
		return _if(n%10 == 1 && n%100 != 11, 0, _if(n%10 >= 2 && (n%100 < 10 || n%100 >= 20), 1, 2))
	},
	// Slavic family: Russian, Ukrainian, Belarusian, Serbian, Croatian
	"n%10==1&&n%100!=11?0:n%10>=2&&n%10<=4&&(n%100<10||n%100>=20)?1:2": func(n int64) int64 {
		return _if(n%10 == 1 && n%100 != 11, 0, _if(n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20), 1, 2))
	},
	// Slavic family: Czech, Slovak
	"(n==1)?0:(n>=2&&n<=4)?1:2": func(n int64) int64 {
		return _if(n == 1, 0, _if(n >= 2 && n <= 4, 1, 2))
	},
	// Slavic family: Polish
	"n==1?0:n%10>=2&&n%10<=4&&(n%100<10||n%100>=20)?1:2": func(n int64) int64 {
		return _if(n == 1, 0, _if(n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20), 1, 2))
	},
	// Slavic family: Slovenian
	"n%100==1?0:n%100==2?1:n%100==3||n%100==4?2:3": func(n int64) int64 {
		return _if(n%100 == 1, 0, _if(n%100 == 2, 1, _if(n%100 == 3 || n%100 == 4, 2, 3)))
	},
	// Afroasiatic family: Arabic
	"n==0?0:n==1?1:n==2?2:n%100>=3&&n%100<=10?3:n%100>=11?4:5": func(n int64) int64 {
		return _if(n == 0, 0,
			_if(n == 1, 1,
				_if(n == 2, 2,
					_if(n%100 >= 3 && n%100 <= 10, 3,
						_if(n%100 >= 11, 4, 5)))))
	},
}
