package t

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func TestTags(t *testing.T) {
	type args struct {
		locales []string
	}
	tests := []struct {
		name     string
		args     args
		wantTags []language.Tag
	}{
		{"tags", args{[]string{"zh", "zh"}}, []language.Tag{language.Make("zh"), language.Make("zh")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTags := Tags(tt.args.locales); !reflect.DeepEqual(gotTags, tt.wantTags) {
				t.Errorf("Tags() = %v, want %v", gotTags, tt.wantTags)
			}
		})
	}
}
