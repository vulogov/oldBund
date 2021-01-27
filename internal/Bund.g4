//
//
//

grammar Bund;

expressions
  : term*
  ;

term
  : NUMBER
  ;

NUMBER
  : INTEGER
  ;

INTEGER
  : DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  ;

SKIP_
  : ( SPACES | COMMENT | LINE_JOINING ) -> skip
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

//
// SKIP fragments
//

fragment LINE_JOINING
  : '\\' SPACES? ( '\r'? '\n' | '\r' | '\f' )
  ;
  
fragment COMMENT
  : '#' ~[\r\n\f]*
  ;

fragment SPACES
  : [ \t]+
  ;
