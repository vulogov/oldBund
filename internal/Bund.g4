//
//
//

grammar Bund;

expressions
  : (term | block | directive | cmd | lambda | DUPLICATE | DROP | TOBEGIN | TOEND )*
  ;

term
  : INTEGER
  | OCT_INTEGER
  | HEX_INTEGER
  | BIN_INTEGER
  | FLOAT_NUMBER
  | TRUE
  | FALSE
  | STRING
  | NAME
  | pair
  ;

block
  : '|' (values+=term)* '|'
  ;

pair
  : '(' head=term tail=term ')'
  ;

directive
  : (op+=DIR)+ name=NAME
  ;

cmd
  : (command+=CMD)+
  ;

lambda
  : '[' name=NAME ']' (body+=term)* '.'
  ;

INTEGER
  : (SIGN)? DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

OCT_INTEGER
  : '0' [oO] [0-7]+
  ;

HEX_INTEGER
  : '0' [xX] [0-9a-fA-F]+
  ;

BIN_INTEGER
  : '0' [bB] [01]+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

STRING
  : SHORT_STRING
  | LONG_STRING
  ;

TRUE
  : 'true'
  | 'TRUE'
  | '#T'
  ;

FALSE
  : 'false'
  | 'FALSE'
  | '#F'
  ;

NAME
  : ID_START ID_CONTINUE*
  ;



TOBEGIN: ':' ;
TOEND:   ';' ;
SLASH:   '/' ;

DIR
  : ('+'|'-'|'*'|'`'|'~')
  ;

CMD
  : ('@'|'$'|'&'|'='|','|'%'|'<'|'>')
  ;

DUPLICATE
  : '^'
  ;

DROP
  : '_'
  ;

SKIP_
  : ( SPACES | COMMENT | LINE_JOINING | CCOMMENT | BLOCK_COMMENT ) -> skip
  ;


//
// Fragments generic term fragments
//
fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment RN
  : '\r'? '\n'
  ;

fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment LONG_STRING
  : '\'\'\'' LONG_STRING_ITEM*? '\'\'\''
  | '"""' LONG_STRING_ITEM*? '"""'
  ;

fragment LONG_STRING_ITEM
  : ~'\\'
  | '\\' (RN | .)
  ;

fragment ID_START
 : ([A-Z]|[a-z]|SLASH)
 | [a-z]
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | SLASH
 ;

//
// SKIP fragments
//

fragment LINE_JOINING
  : '\\' SPACES? ( '\r'? '\n' | '\r' | '\f' )
  ;

fragment COMMENT
  : '##' ~[\r\n\f]*
  ;

fragment BLOCK_COMMENT
  :   '/*' .*? '*/'
  ;

fragment CCOMMENT
  :   '//' ~[\r\n]*
  ;

fragment SPACES
  : [ \t]+
  ;
