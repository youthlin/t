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

--------------------
左递归消除
A : A a1
  | A ai
  | b1
  | bi
消除方法，变成右递归
A  : b1 A'
   | bi A'
A' : a1 A'
   | ai A'
   | 空
--------------------

start : exp;
exp : primary expMore
    | ('+' | '-' | '++' | '--') exp expMore
    | ('~' | '!') exp expMore
    ;
expMore : ('++' | '--') expMore
        | ('*' | '/' | '%') exp expMore
        | ('+' | '-') exp expMore
        | ('>>' | '<<') exp expMore
        | ('>' | '<') exp expMore
        | ('>=' | '<=') exp expMore
        | ('==' | '!=') exp expMore
        | ('&' | '^' | '|' | '&&' | '||') exp expMore
        | '?' exp ':' exp expMore
        | 空
        ;
primary : '(' exp ')'
        | 'n'
        | INT
        ;
INT : [0-9]+;

--------------------
Predict(A->a) 产生式的预测符集
= First(a)                    --- if 空不属于 First(a)
= First(a)-空 并上 Follow(A)  --- else
--------------------
First(a) 子串 a 的首字符集
= {t | a =>* tb , t 属于终极符}
--------------------
Follow(A) 非终极符的后继字符集
= {t | S =>* aAb, t 属于 First(b), t属于终极符}
--------------------

递归下降流程
1. 为每条产生式计算预测符集;
判断是否满足：对同一个非终极符的任意两条产生式，其预测符集无交集
2. 为每个非终极符编写分析子程序
