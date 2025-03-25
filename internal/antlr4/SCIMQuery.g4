grammar SCIMQuery;

root
  : query EOF
  ;

query
  : NOT? SP? LPAREN SP? query SP? RPAREN  #parenExp
  | query SP LOGICAL_OPERATOR SP query    #logicalExp
  | attrPath SP 'pr'                      #presentExp
  | (attrPath | functionCall) SP op=(EQ|NE|GT|LT|GE|LE|CO|SW|EW|IN) SP value #compareExp
  ;

NOT : 'not' | 'NOT' ;
LOGICAL_OPERATOR : 'and' | 'or' ;

BOOLEAN : 'true' | 'false' ;
NULL : 'null' | 'nil' ;
IN : 'IN' | 'in';
EQ : 'eq' | 'EQ';
NE : 'ne' | 'NE';
GT : 'gt' | 'GT';
LT : 'lt' | 'LT';
GE : 'ge' | 'GE';
LE : 'le' | 'LE';
CO : 'co' | 'CO';
SW : 'sw' | 'SW';
EW : 'ew' | 'EW';

attrPath
   : ATTRNAME subAttr?
   | functionCall
   ;

typeAnnotation
  : '[f64]' | '[i64]' | '[ui64]' | '[i]' | '[ui]' | '[i32]' | '[ui32]' | '[d]' | '[s]' | '[f32]'
  ;

functionCall
   : ATTRNAME LPAREN argList? RPAREN
   ;

argList
   : value (COMMA value)*
   ;

subAttr
   : '.' attrPath
   ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : [0-9]
   ;

fragment ALPHA
   : [A-Za-z]
   ;

typedValue
   : typeAnnotation? STRING           #typedString
   | typeAnnotation? DOUBLE           #typedDouble
   | typeAnnotation? '-'? INT EXP?    #typedInteger
   ;

value
   : typedValue        #typedVal
   | BOOLEAN           #boolean
   | NULL              #null
   | listInts          #listOfInts
   | listDoubles       #listOfDoubles
   | listStrings       #listOfStrings
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

listStrings
   : '[' subListOfStrings
   ;

subListOfStrings
   : STRING COMMA subListOfStrings
   | STRING ']'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9]+ EXP?
   ;

listDoubles
   : '[' subListOfDoubles
   ;

subListOfDoubles
   : DOUBLE COMMA subListOfDoubles
   | DOUBLE ']'
   ;

listInts
   : '[' subListOfInts
   ;

subListOfInts
   : INT COMMA subListOfInts
   | INT ']'
   ;

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

LPAREN : '(' ;
RPAREN : ')' ;

// EXP we escape '-' with '\-' in a bracket expression
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

COMMA
   : ',' ' '*;

SP
   : ' ' NEWLINE*
   ;
