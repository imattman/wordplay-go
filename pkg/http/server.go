package http

import (
	_ "net/http"
)

var notes = `
Valid Word
==========
/.../word/fizzlegrub

response:  200 | 404
{
  "query" : "fizzlegrub",
  "in_lexicon" : true|false,
  "score" : 22
}


Matches for Rack
================
/.../matches/a+b+c+d+e+i   (accept spaces)

response:  200 | 404

{
  "characters" : ['a', 'b', 'c', 'd', 'e', 'i'],
	"matches" : [
	  { "word": "~~~~~~~", "score": 22 },
	  { "word": "~~~~~", "score": 12 },
	  { "word": "~~~", "score": 2 }
	]
}


Full Lexicon
============
/.../lexicon

response: 200

["aaa", "bbb", "ccc" ]
`
