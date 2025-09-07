grammar KQuery;

// ---------- Parser ----------
query
  : FROM resource (WHERE expr)? SELECT fieldList (LIMIT INT)? EOF
  ;

resource : IDENT ;
fieldList: field (',' field)* ;
field    : IDENT ('.' IDENT)* ;

// Placeholder — we'll grow this in next steps
expr     : IDENT ;

// ---------- Lexer (keywords before IDENT so they win) ----------
FROM   : 'from' ;
WHERE  : 'where' ;
SELECT : 'select' ;
LIMIT  : 'limit' ;

IDENT  : [a-zA-Z_][a-zA-Z0-9_]* ;
INT    : [0-9]+ ;

WS     : [ \t\r\n]+ -> skip ;
COMMENT: '//' ~[\r\n]* -> skip ;
