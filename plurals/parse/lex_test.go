package parse

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantTokens []*Token
		wantErr    bool
	}{
		{name: "empty", args: args{""}, wantTokens: nil, wantErr: false},
		{name: "id-n", args: args{"n"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "id-nn", args: args{"nn"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "nn", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: "id-space", args: args{"a b"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "a", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeID, Value: "b", Start: Pos{1, 3}, End: Pos{1, 4}},
		}, wantErr: false},
		{name: "num-0", args: args{"0"}, wantTokens: []*Token{
			{Type: TokenTypeInt, Value: "0", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "num-10", args: args{"10"}, wantTokens: []*Token{
			{Type: TokenTypeInt, Value: "10", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: "num-space", args: args{"0 1"}, wantTokens: []*Token{
			{Type: TokenTypeInt, Value: "0", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 3}, End: Pos{1, 4}},
		}, wantErr: false},
		{name: "+", args: args{"+"}, wantTokens: []*Token{
			{Type: TokenTypePlus, Value: "+", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "++", args: args{"++"}, wantTokens: []*Token{
			{Type: TokenTypeIncr, Value: "++", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: "-", args: args{"-"}, wantTokens: []*Token{
			{Type: TokenTypeMinus, Value: "-", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "--", args: args{"--"}, wantTokens: []*Token{
			{Type: TokenTypeDecr, Value: "--", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: ">", args: args{">"}, wantTokens: []*Token{
			{Type: TokenTypeGt, Value: ">", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: ">>", args: args{">>"}, wantTokens: []*Token{
			{Type: TokenTypeShiftR, Value: ">>", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: ">=", args: args{">="}, wantTokens: []*Token{
			{Type: TokenTypeGe, Value: ">=", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},

		{name: "<", args: args{"<"}, wantTokens: []*Token{
			{Type: TokenTypeLt, Value: "<", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "<<", args: args{"<<"}, wantTokens: []*Token{
			{Type: TokenTypeShiftL, Value: "<<", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: "<=", args: args{"<="}, wantTokens: []*Token{
			{Type: TokenTypeLe, Value: "<=", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},

		{name: "!", args: args{"!"}, wantTokens: []*Token{
			{Type: TokenTypeNot, Value: "!", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "!=", args: args{"!="}, wantTokens: []*Token{
			{Type: TokenTypeNe, Value: "!=", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},

		{name: "&", args: args{"&"}, wantTokens: []*Token{
			{Type: TokenTypeBitAnd, Value: "&", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "&&", args: args{"&&"}, wantTokens: []*Token{
			{Type: TokenTypeAnd, Value: "&&", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},

		{name: "|", args: args{"|"}, wantTokens: []*Token{
			{Type: TokenTypeBitOr, Value: "|", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "||", args: args{"||"}, wantTokens: []*Token{
			{Type: TokenTypeOr, Value: "||", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},

		{name: "==", args: args{"=="}, wantTokens: []*Token{
			{Type: TokenTypeEq, Value: "==", Start: Pos{1, 1}, End: Pos{1, 3}},
		}, wantErr: false},
		{name: "=", args: args{"="}, wantTokens: nil, wantErr: true},

		{name: "(", args: args{"("}, wantTokens: []*Token{
			{Type: TokenTypeParenL, Value: "(", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: ")", args: args{")"}, wantTokens: []*Token{
			{Type: TokenTypeParenR, Value: ")", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "*", args: args{"*"}, wantTokens: []*Token{
			{Type: TokenTypeTimes, Value: "*", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "/", args: args{"/"}, wantTokens: []*Token{
			{Type: TokenTypeOver, Value: "/", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "%", args: args{"%"}, wantTokens: []*Token{
			{Type: TokenTypeMod, Value: "%", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "~", args: args{"~"}, wantTokens: []*Token{
			{Type: TokenTypeBitNot, Value: "~", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "^", args: args{"^"}, wantTokens: []*Token{
			{Type: TokenTypeBitXor, Value: "^", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "?", args: args{"?"}, wantTokens: []*Token{
			{Type: TokenTypeQst, Value: "?", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: ":", args: args{":"}, wantTokens: []*Token{
			{Type: TokenTypeCol, Value: ":", Start: Pos{1, 1}, End: Pos{1, 2}},
		}, wantErr: false},
		{name: "err", args: args{"\\"}, wantTokens: nil, wantErr: true},

		{name: "n!=1", args: args{"n!=1"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeNe, Value: "!=", Start: Pos{1, 2}, End: Pos{1, 4}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 4}, End: Pos{1, 5}},
		}, wantErr: false},
		{name: "n>1", args: args{"n>1"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeGt, Value: ">", Start: Pos{1, 2}, End: Pos{1, 3}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 3}, End: Pos{1, 4}},
		}, wantErr: false},
		{name: "n%10==1&&n%100!=11?0:n!=0?1: 2", args: args{"n%10==1&&n%100!=11?0:n!=0?1: 2"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeMod, Value: "%", Start: Pos{1, 2}, End: Pos{1, 3}},
			{Type: TokenTypeInt, Value: "10", Start: Pos{1, 3}, End: Pos{1, 5}},
			{Type: TokenTypeEq, Value: "==", Start: Pos{1, 5}, End: Pos{1, 7}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 7}, End: Pos{1, 8}},
			{Type: TokenTypeAnd, Value: "&&", Start: Pos{1, 8}, End: Pos{1, 10}},
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 10}, End: Pos{1, 11}},
			{Type: TokenTypeMod, Value: "%", Start: Pos{1, 11}, End: Pos{1, 12}},
			{Type: TokenTypeInt, Value: "100", Start: Pos{1, 12}, End: Pos{1, 15}},
			{Type: TokenTypeNe, Value: "!=", Start: Pos{1, 15}, End: Pos{1, 17}},
			{Type: TokenTypeInt, Value: "11", Start: Pos{1, 17}, End: Pos{1, 19}},
			{Type: TokenTypeQst, Value: "?", Start: Pos{1, 19}, End: Pos{1, 20}},
			{Type: TokenTypeInt, Value: "0", Start: Pos{1, 20}, End: Pos{1, 21}},
			{Type: TokenTypeCol, Value: ":", Start: Pos{1, 21}, End: Pos{1, 22}},
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 22}, End: Pos{1, 23}},
			{Type: TokenTypeNe, Value: "!=", Start: Pos{1, 23}, End: Pos{1, 25}},
			{Type: TokenTypeInt, Value: "0", Start: Pos{1, 25}, End: Pos{1, 26}},
			{Type: TokenTypeQst, Value: "?", Start: Pos{1, 26}, End: Pos{1, 27}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 27}, End: Pos{1, 28}},
			{Type: TokenTypeCol, Value: ":", Start: Pos{1, 28}, End: Pos{1, 29}},
			{Type: TokenTypeInt, Value: "2", Start: Pos{1, 30}, End: Pos{1, 31}},
		}, wantErr: false},
		{name: "n==1?0:n==2?1:2", args: args{"n==1?0:n==2?1:2"}, wantTokens: []*Token{
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 1}, End: Pos{1, 2}},
			{Type: TokenTypeEq, Value: "==", Start: Pos{1, 2}, End: Pos{1, 4}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 4}, End: Pos{1, 5}},
			{Type: TokenTypeQst, Value: "?", Start: Pos{1, 5}, End: Pos{1, 6}},
			{Type: TokenTypeInt, Value: "0", Start: Pos{1, 6}, End: Pos{1, 7}},
			{Type: TokenTypeCol, Value: ":", Start: Pos{1, 7}, End: Pos{1, 8}},
			{Type: TokenTypeID, Value: "n", Start: Pos{1, 8}, End: Pos{1, 9}},
			{Type: TokenTypeEq, Value: "==", Start: Pos{1, 9}, End: Pos{1, 11}},
			{Type: TokenTypeInt, Value: "2", Start: Pos{1, 11}, End: Pos{1, 12}},
			{Type: TokenTypeQst, Value: "?", Start: Pos{1, 12}, End: Pos{1, 13}},
			{Type: TokenTypeInt, Value: "1", Start: Pos{1, 13}, End: Pos{1, 14}},
			{Type: TokenTypeCol, Value: ":", Start: Pos{1, 14}, End: Pos{1, 15}},
			{Type: TokenTypeInt, Value: "2", Start: Pos{1, 15}, End: Pos{1, 16}},
		}, wantErr: false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTokens, err := Lex(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTokens, tt.wantTokens) {
				t.Errorf("Lex() = %v, want %v", toJSON(gotTokens), toJSON(tt.wantTokens))
			}
		})
	}
}

func toJSON(a any) string {
	b, _ := json.Marshal(a)
	return string(b)
}
