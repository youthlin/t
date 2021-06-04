grammar plural;

// 关于复数表达式的官方描述：
// https://www.gnu.org/software/gettext/manual/html_node/Plural-forms.html#index-specifying-plural-form-in-a-PO-file
// plural 表达式是 C 语法的表达式，但不允许负数出现，而且只能出现整数，变量只允许 n。可以有空白，但不能反斜杠换行。

// 生成 Go 代码的命令 （在本目录下执行）： antlr4 -Dlanguage=Go plural.g4 -o parser

start: exp;
exp:
	primary
	| exp postfix = ('++' | '--')
	| prefix = ('+' | '-' | '++' | '--') exp
	| prefix = ('~' | '!') exp
	| exp bop = ('*' | '/' | '%') exp
	| exp bop = ('+' | '-') exp
	| exp bop = ('>>' | '<<') exp
	| exp bop = ('>' | '<') exp
	| exp bop = ('>=' | '<=') exp
	| exp bop = ('==' | '!=') exp
	| exp bop = '&' exp
	| exp bop = '^' exp
	| exp bop = '|' exp
	| exp bop = '&&' exp
	| exp bop = '||' exp
	| <assoc = right> exp bop = '?' exp ':' exp;
primary: '(' exp ')' | 'n' | INT;
INT: [0-9]+;
WS: [ \t]+ -> skip;
