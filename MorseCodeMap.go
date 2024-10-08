package convert

// AsciiToMorseCode - A map used by ToMorseCode() and FromMorseCode() with all supported characters.
var AsciiToMorseCode = map[string]string{
	"A":  ".-",
	"B":  "-...",
	"C":  "-.-.",
	"D":  "-..",
	"E":  ".",
	"F":  "..-.",
	"G":  "--.",
	"H":  "....",
	"I":  "..",
	"J":  ".---",
	"K":  "-.-",
	"L":  ".-..",
	"M":  "--",
	"N":  "-.",
	"O":  "---",
	"P":  ".--.",
	"Q":  "--.-",
	"R":  ".-.",
	"S":  "...",
	"T":  "-",
	"U":  "..-",
	"V":  "...-",
	"W":  ".--",
	"X":  "-..-",
	"Y":  "-.--",
	"Z":  "--..",
	"0":  "-----",
	"1":  ".----",
	"2":  "..---",
	"3":  "...--",
	"4":  "....-",
	"5":  ".....",
	"6":  "-....",
	"7":  "--...",
	"8":  "---..",
	"9":  "----.",
	".":  ".-.-.-",
	",":  "--..--",
	"?":  "..--..",
	"'":  ".----.",
	"!":  "-.-.--",
	"/":  "-..-.",
	"(":  "-.--.",
	")":  "-.--.-",
	"&":  ".-...",
	":":  "---...",
	";":  "-.-.-.",
	"=":  "-...-",
	"+":  ".-.-.",
	"-":  "-....-",
	"_":  "..--.-",
	"\"": ".-..-.",
	"$":  "...-..-",
	"@":  ".--.-.",
	" ":  "/",
}
var morseCodeToASCII = map[string]string{
	".-":      "A",
	"-...":    "B",
	"-.-.":    "C",
	"-..":     "D",
	".":       "E",
	"..-.":    "F",
	"--.":     "G",
	"....":    "H",
	"..":      "I",
	".---":    "J",
	"-.-":     "K",
	".-..":    "L",
	"--":      "M",
	"-.":      "N",
	"---":     "O",
	".--.":    "P",
	"--.-":    "Q",
	".-.":     "R",
	"...":     "S",
	"-":       "T",
	"..-":     "U",
	"...-":    "V",
	".--":     "W",
	"-..-":    "X",
	"-.--":    "Y",
	"--..":    "Z",
	"-----":   "0",
	".----":   "1",
	"..---":   "2",
	"...--":   "3",
	"....-":   "4",
	".....":   "5",
	"-....":   "6",
	"--...":   "7",
	"---..":   "8",
	"----.":   "9",
	".-.-.-":  ".",
	"--..--":  ",",
	"..--..":  "?",
	".----.":  "'",
	"-.-.--":  "!",
	"-..-.":   "/",
	"-.--.":   "(",
	"-.--.-":  ")",
	".-...":   "&",
	"---...":  ":",
	"-.-.-.":  ";",
	"-...-":   "=",
	".-.-.":   "+",
	"-....-":  "-",
	"..--.-":  "_",
	".-..-.":  "\"",
	"...-..-": "$",
	".--.-.":  "@",
	"/":       " ",
}
