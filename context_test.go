package t

import (
	"context"
	"reflect"
	"testing"
)

func TestWithContext(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *Translations
	}{
		{name: "init", args: args{ctx: ctx}, want: NewTranslations()},
		{name: "locale", args: args{ctx: SetCtxLocale(ctx, "zh_CN")}, want: NewTranslations().L("zh_CN")},
		{name: "domain", args: args{ctx: SetCtxDomain(ctx, "my-domain")}, want: NewTranslations().D("my-domain")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
