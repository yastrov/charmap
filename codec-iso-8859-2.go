
package charmap

func init() {

	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // 	NULL
		'\x01':	'\u0001',	 // 	START OF HEADING
		'\x02':	'\u0002',	 // 	START OF TEXT
		'\x03':	'\u0003',	 // 	END OF TEXT
		'\x04':	'\u0004',	 // 	END OF TRANSMISSION
		'\x05':	'\u0005',	 // 	ENQUIRY
		'\x06':	'\u0006',	 // 	ACKNOWLEDGE
		'\x07':	'\u0007',	 // 	BELL
		'\x08':	'\u0008',	 // 	BACKSPACE
		'\x09':	'\u0009',	 // 	HORIZONTAL TABULATION
		'\x0A':	'\u000A',	 // 	LINE FEED
		'\x0B':	'\u000B',	 // 	VERTICAL TABULATION
		'\x0C':	'\u000C',	 // 	FORM FEED
		'\x0D':	'\u000D',	 // 	CARRIAGE RETURN
		'\x0E':	'\u000E',	 // 	SHIFT OUT
		'\x0F':	'\u000F',	 // 	SHIFT IN
		'\x10':	'\u0010',	 // 	DATA LINK ESCAPE
		'\x11':	'\u0011',	 // 	DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // 	DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // 	DEVICE CONTROL THREE
		'\x14':	'\u0014',	 // 	DEVICE CONTROL FOUR
		'\x15':	'\u0015',	 // 	NEGATIVE ACKNOWLEDGE
		'\x16':	'\u0016',	 // 	SYNCHRONOUS IDLE
		'\x17':	'\u0017',	 // 	END OF TRANSMISSION BLOCK
		'\x18':	'\u0018',	 // 	CANCEL
		'\x19':	'\u0019',	 // 	END OF MEDIUM
		'\x1A':	'\u001A',	 // 	SUBSTITUTE
		'\x1B':	'\u001B',	 // 	ESCAPE
		'\x1C':	'\u001C',	 // 	FILE SEPARATOR
		'\x1D':	'\u001D',	 // 	GROUP SEPARATOR
		'\x1E':	'\u001E',	 // 	RECORD SEPARATOR
		'\x1F':	'\u001F',	 // 	UNIT SEPARATOR
		'\x20':	'\u0020',	 // 	SPACE
		'\x21':	'\u0021',	 // 	EXCLAMATION MARK
		'\x22':	'\u0022',	 // 	QUOTATION MARK
		'\x23':	'\u0023',	 // 	NUMBER SIGN
		'\x24':	'\u0024',	 // 	DOLLAR SIGN
		'\x25':	'\u0025',	 // 	PERCENT SIGN
		'\x26':	'\u0026',	 // 	AMPERSAND
		'\x27':	'\u0027',	 // 	APOSTROPHE
		'\x28':	'\u0028',	 // 	LEFT PARENTHESIS
		'\x29':	'\u0029',	 // 	RIGHT PARENTHESIS
		'\x2A':	'\u002A',	 // 	ASTERISK
		'\x2B':	'\u002B',	 // 	PLUS SIGN
		'\x2C':	'\u002C',	 // 	COMMA
		'\x2D':	'\u002D',	 // 	HYPHEN-MINUS
		'\x2E':	'\u002E',	 // 	FULL STOP
		'\x2F':	'\u002F',	 // 	SOLIDUS
		'\x30':	'\u0030',	 // 	DIGIT ZERO
		'\x31':	'\u0031',	 // 	DIGIT ONE
		'\x32':	'\u0032',	 // 	DIGIT TWO
		'\x33':	'\u0033',	 // 	DIGIT THREE
		'\x34':	'\u0034',	 // 	DIGIT FOUR
		'\x35':	'\u0035',	 // 	DIGIT FIVE
		'\x36':	'\u0036',	 // 	DIGIT SIX
		'\x37':	'\u0037',	 // 	DIGIT SEVEN
		'\x38':	'\u0038',	 // 	DIGIT EIGHT
		'\x39':	'\u0039',	 // 	DIGIT NINE
		'\x3A':	'\u003A',	 // 	COLON
		'\x3B':	'\u003B',	 // 	SEMICOLON
		'\x3C':	'\u003C',	 // 	LESS-THAN SIGN
		'\x3D':	'\u003D',	 // 	EQUALS SIGN
		'\x3E':	'\u003E',	 // 	GREATER-THAN SIGN
		'\x3F':	'\u003F',	 // 	QUESTION MARK
		'\x40':	'\u0040',	 // 	COMMERCIAL AT
		'\x41':	'\u0041',	 // 	LATIN CAPITAL LETTER A
		'\x42':	'\u0042',	 // 	LATIN CAPITAL LETTER B
		'\x43':	'\u0043',	 // 	LATIN CAPITAL LETTER C
		'\x44':	'\u0044',	 // 	LATIN CAPITAL LETTER D
		'\x45':	'\u0045',	 // 	LATIN CAPITAL LETTER E
		'\x46':	'\u0046',	 // 	LATIN CAPITAL LETTER F
		'\x47':	'\u0047',	 // 	LATIN CAPITAL LETTER G
		'\x48':	'\u0048',	 // 	LATIN CAPITAL LETTER H
		'\x49':	'\u0049',	 // 	LATIN CAPITAL LETTER I
		'\x4A':	'\u004A',	 // 	LATIN CAPITAL LETTER J
		'\x4B':	'\u004B',	 // 	LATIN CAPITAL LETTER K
		'\x4C':	'\u004C',	 // 	LATIN CAPITAL LETTER L
		'\x4D':	'\u004D',	 // 	LATIN CAPITAL LETTER M
		'\x4E':	'\u004E',	 // 	LATIN CAPITAL LETTER N
		'\x4F':	'\u004F',	 // 	LATIN CAPITAL LETTER O
		'\x50':	'\u0050',	 // 	LATIN CAPITAL LETTER P
		'\x51':	'\u0051',	 // 	LATIN CAPITAL LETTER Q
		'\x52':	'\u0052',	 // 	LATIN CAPITAL LETTER R
		'\x53':	'\u0053',	 // 	LATIN CAPITAL LETTER S
		'\x54':	'\u0054',	 // 	LATIN CAPITAL LETTER T
		'\x55':	'\u0055',	 // 	LATIN CAPITAL LETTER U
		'\x56':	'\u0056',	 // 	LATIN CAPITAL LETTER V
		'\x57':	'\u0057',	 // 	LATIN CAPITAL LETTER W
		'\x58':	'\u0058',	 // 	LATIN CAPITAL LETTER X
		'\x59':	'\u0059',	 // 	LATIN CAPITAL LETTER Y
		'\x5A':	'\u005A',	 // 	LATIN CAPITAL LETTER Z
		'\x5B':	'\u005B',	 // 	LEFT SQUARE BRACKET
		'\x5C':	'\u005C',	 // 	REVERSE SOLIDUS
		'\x5D':	'\u005D',	 // 	RIGHT SQUARE BRACKET
		'\x5E':	'\u005E',	 // 	CIRCUMFLEX ACCENT
		'\x5F':	'\u005F',	 // 	LOW LINE
		'\x60':	'\u0060',	 // 	GRAVE ACCENT
		'\x61':	'\u0061',	 // 	LATIN SMALL LETTER A
		'\x62':	'\u0062',	 // 	LATIN SMALL LETTER B
		'\x63':	'\u0063',	 // 	LATIN SMALL LETTER C
		'\x64':	'\u0064',	 // 	LATIN SMALL LETTER D
		'\x65':	'\u0065',	 // 	LATIN SMALL LETTER E
		'\x66':	'\u0066',	 // 	LATIN SMALL LETTER F
		'\x67':	'\u0067',	 // 	LATIN SMALL LETTER G
		'\x68':	'\u0068',	 // 	LATIN SMALL LETTER H
		'\x69':	'\u0069',	 // 	LATIN SMALL LETTER I
		'\x6A':	'\u006A',	 // 	LATIN SMALL LETTER J
		'\x6B':	'\u006B',	 // 	LATIN SMALL LETTER K
		'\x6C':	'\u006C',	 // 	LATIN SMALL LETTER L
		'\x6D':	'\u006D',	 // 	LATIN SMALL LETTER M
		'\x6E':	'\u006E',	 // 	LATIN SMALL LETTER N
		'\x6F':	'\u006F',	 // 	LATIN SMALL LETTER O
		'\x70':	'\u0070',	 // 	LATIN SMALL LETTER P
		'\x71':	'\u0071',	 // 	LATIN SMALL LETTER Q
		'\x72':	'\u0072',	 // 	LATIN SMALL LETTER R
		'\x73':	'\u0073',	 // 	LATIN SMALL LETTER S
		'\x74':	'\u0074',	 // 	LATIN SMALL LETTER T
		'\x75':	'\u0075',	 // 	LATIN SMALL LETTER U
		'\x76':	'\u0076',	 // 	LATIN SMALL LETTER V
		'\x77':	'\u0077',	 // 	LATIN SMALL LETTER W
		'\x78':	'\u0078',	 // 	LATIN SMALL LETTER X
		'\x79':	'\u0079',	 // 	LATIN SMALL LETTER Y
		'\x7A':	'\u007A',	 // 	LATIN SMALL LETTER Z
		'\x7B':	'\u007B',	 // 	LEFT CURLY BRACKET
		'\x7C':	'\u007C',	 // 	VERTICAL LINE
		'\x7D':	'\u007D',	 // 	RIGHT CURLY BRACKET
		'\x7E':	'\u007E',	 // 	TILDE
		'\x7F':	'\u007F',	 // 	DELETE
		'\x80':	'\u0080',	 // 	<control>
		'\x81':	'\u0081',	 // 	<control>
		'\x82':	'\u0082',	 // 	<control>
		'\x83':	'\u0083',	 // 	<control>
		'\x84':	'\u0084',	 // 	<control>
		'\x85':	'\u0085',	 // 	<control>
		'\x86':	'\u0086',	 // 	<control>
		'\x87':	'\u0087',	 // 	<control>
		'\x88':	'\u0088',	 // 	<control>
		'\x89':	'\u0089',	 // 	<control>
		'\x8A':	'\u008A',	 // 	<control>
		'\x8B':	'\u008B',	 // 	<control>
		'\x8C':	'\u008C',	 // 	<control>
		'\x8D':	'\u008D',	 // 	<control>
		'\x8E':	'\u008E',	 // 	<control>
		'\x8F':	'\u008F',	 // 	<control>
		'\x90':	'\u0090',	 // 	<control>
		'\x91':	'\u0091',	 // 	<control>
		'\x92':	'\u0092',	 // 	<control>
		'\x93':	'\u0093',	 // 	<control>
		'\x94':	'\u0094',	 // 	<control>
		'\x95':	'\u0095',	 // 	<control>
		'\x96':	'\u0096',	 // 	<control>
		'\x97':	'\u0097',	 // 	<control>
		'\x98':	'\u0098',	 // 	<control>
		'\x99':	'\u0099',	 // 	<control>
		'\x9A':	'\u009A',	 // 	<control>
		'\x9B':	'\u009B',	 // 	<control>
		'\x9C':	'\u009C',	 // 	<control>
		'\x9D':	'\u009D',	 // 	<control>
		'\x9E':	'\u009E',	 // 	<control>
		'\x9F':	'\u009F',	 // 	<control>
		'\xA0':	'\u00A0',	 // 	NO-BREAK SPACE
		'\xA1':	'\u0104',	 // 	LATIN CAPITAL LETTER A WITH OGONEK
		'\xA2':	'\u02D8',	 // 	BREVE
		'\xA3':	'\u0141',	 // 	LATIN CAPITAL LETTER L WITH STROKE
		'\xA4':	'\u00A4',	 // 	CURRENCY SIGN
		'\xA5':	'\u013D',	 // 	LATIN CAPITAL LETTER L WITH CARON
		'\xA6':	'\u015A',	 // 	LATIN CAPITAL LETTER S WITH ACUTE
		'\xA7':	'\u00A7',	 // 	SECTION SIGN
		'\xA8':	'\u00A8',	 // 	DIAERESIS
		'\xA9':	'\u0160',	 // 	LATIN CAPITAL LETTER S WITH CARON
		'\xAA':	'\u015E',	 // 	LATIN CAPITAL LETTER S WITH CEDILLA
		'\xAB':	'\u0164',	 // 	LATIN CAPITAL LETTER T WITH CARON
		'\xAC':	'\u0179',	 // 	LATIN CAPITAL LETTER Z WITH ACUTE
		'\xAD':	'\u00AD',	 // 	SOFT HYPHEN
		'\xAE':	'\u017D',	 // 	LATIN CAPITAL LETTER Z WITH CARON
		'\xAF':	'\u017B',	 // 	LATIN CAPITAL LETTER Z WITH DOT ABOVE
		'\xB0':	'\u00B0',	 // 	DEGREE SIGN
		'\xB1':	'\u0105',	 // 	LATIN SMALL LETTER A WITH OGONEK
		'\xB2':	'\u02DB',	 // 	OGONEK
		'\xB3':	'\u0142',	 // 	LATIN SMALL LETTER L WITH STROKE
		'\xB4':	'\u00B4',	 // 	ACUTE ACCENT
		'\xB5':	'\u013E',	 // 	LATIN SMALL LETTER L WITH CARON
		'\xB6':	'\u015B',	 // 	LATIN SMALL LETTER S WITH ACUTE
		'\xB7':	'\u02C7',	 // 	CARON
		'\xB8':	'\u00B8',	 // 	CEDILLA
		'\xB9':	'\u0161',	 // 	LATIN SMALL LETTER S WITH CARON
		'\xBA':	'\u015F',	 // 	LATIN SMALL LETTER S WITH CEDILLA
		'\xBB':	'\u0165',	 // 	LATIN SMALL LETTER T WITH CARON
		'\xBC':	'\u017A',	 // 	LATIN SMALL LETTER Z WITH ACUTE
		'\xBD':	'\u02DD',	 // 	DOUBLE ACUTE ACCENT
		'\xBE':	'\u017E',	 // 	LATIN SMALL LETTER Z WITH CARON
		'\xBF':	'\u017C',	 // 	LATIN SMALL LETTER Z WITH DOT ABOVE
		'\xC0':	'\u0154',	 // 	LATIN CAPITAL LETTER R WITH ACUTE
		'\xC1':	'\u00C1',	 // 	LATIN CAPITAL LETTER A WITH ACUTE
		'\xC2':	'\u00C2',	 // 	LATIN CAPITAL LETTER A WITH CIRCUMFLEX
		'\xC3':	'\u0102',	 // 	LATIN CAPITAL LETTER A WITH BREVE
		'\xC4':	'\u00C4',	 // 	LATIN CAPITAL LETTER A WITH DIAERESIS
		'\xC5':	'\u0139',	 // 	LATIN CAPITAL LETTER L WITH ACUTE
		'\xC6':	'\u0106',	 // 	LATIN CAPITAL LETTER C WITH ACUTE
		'\xC7':	'\u00C7',	 // 	LATIN CAPITAL LETTER C WITH CEDILLA
		'\xC8':	'\u010C',	 // 	LATIN CAPITAL LETTER C WITH CARON
		'\xC9':	'\u00C9',	 // 	LATIN CAPITAL LETTER E WITH ACUTE
		'\xCA':	'\u0118',	 // 	LATIN CAPITAL LETTER E WITH OGONEK
		'\xCB':	'\u00CB',	 // 	LATIN CAPITAL LETTER E WITH DIAERESIS
		'\xCC':	'\u011A',	 // 	LATIN CAPITAL LETTER E WITH CARON
		'\xCD':	'\u00CD',	 // 	LATIN CAPITAL LETTER I WITH ACUTE
		'\xCE':	'\u00CE',	 // 	LATIN CAPITAL LETTER I WITH CIRCUMFLEX
		'\xCF':	'\u010E',	 // 	LATIN CAPITAL LETTER D WITH CARON
		'\xD0':	'\u0110',	 // 	LATIN CAPITAL LETTER D WITH STROKE
		'\xD1':	'\u0143',	 // 	LATIN CAPITAL LETTER N WITH ACUTE
		'\xD2':	'\u0147',	 // 	LATIN CAPITAL LETTER N WITH CARON
		'\xD3':	'\u00D3',	 // 	LATIN CAPITAL LETTER O WITH ACUTE
		'\xD4':	'\u00D4',	 // 	LATIN CAPITAL LETTER O WITH CIRCUMFLEX
		'\xD5':	'\u0150',	 // 	LATIN CAPITAL LETTER O WITH DOUBLE ACUTE
		'\xD6':	'\u00D6',	 // 	LATIN CAPITAL LETTER O WITH DIAERESIS
		'\xD7':	'\u00D7',	 // 	MULTIPLICATION SIGN
		'\xD8':	'\u0158',	 // 	LATIN CAPITAL LETTER R WITH CARON
		'\xD9':	'\u016E',	 // 	LATIN CAPITAL LETTER U WITH RING ABOVE
		'\xDA':	'\u00DA',	 // 	LATIN CAPITAL LETTER U WITH ACUTE
		'\xDB':	'\u0170',	 // 	LATIN CAPITAL LETTER U WITH DOUBLE ACUTE
		'\xDC':	'\u00DC',	 // 	LATIN CAPITAL LETTER U WITH DIAERESIS
		'\xDD':	'\u00DD',	 // 	LATIN CAPITAL LETTER Y WITH ACUTE
		'\xDE':	'\u0162',	 // 	LATIN CAPITAL LETTER T WITH CEDILLA
		'\xDF':	'\u00DF',	 // 	LATIN SMALL LETTER SHARP S
		'\xE0':	'\u0155',	 // 	LATIN SMALL LETTER R WITH ACUTE
		'\xE1':	'\u00E1',	 // 	LATIN SMALL LETTER A WITH ACUTE
		'\xE2':	'\u00E2',	 // 	LATIN SMALL LETTER A WITH CIRCUMFLEX
		'\xE3':	'\u0103',	 // 	LATIN SMALL LETTER A WITH BREVE
		'\xE4':	'\u00E4',	 // 	LATIN SMALL LETTER A WITH DIAERESIS
		'\xE5':	'\u013A',	 // 	LATIN SMALL LETTER L WITH ACUTE
		'\xE6':	'\u0107',	 // 	LATIN SMALL LETTER C WITH ACUTE
		'\xE7':	'\u00E7',	 // 	LATIN SMALL LETTER C WITH CEDILLA
		'\xE8':	'\u010D',	 // 	LATIN SMALL LETTER C WITH CARON
		'\xE9':	'\u00E9',	 // 	LATIN SMALL LETTER E WITH ACUTE
		'\xEA':	'\u0119',	 // 	LATIN SMALL LETTER E WITH OGONEK
		'\xEB':	'\u00EB',	 // 	LATIN SMALL LETTER E WITH DIAERESIS
		'\xEC':	'\u011B',	 // 	LATIN SMALL LETTER E WITH CARON
		'\xED':	'\u00ED',	 // 	LATIN SMALL LETTER I WITH ACUTE
		'\xEE':	'\u00EE',	 // 	LATIN SMALL LETTER I WITH CIRCUMFLEX
		'\xEF':	'\u010F',	 // 	LATIN SMALL LETTER D WITH CARON
		'\xF0':	'\u0111',	 // 	LATIN SMALL LETTER D WITH STROKE
		'\xF1':	'\u0144',	 // 	LATIN SMALL LETTER N WITH ACUTE
		'\xF2':	'\u0148',	 // 	LATIN SMALL LETTER N WITH CARON
		'\xF3':	'\u00F3',	 // 	LATIN SMALL LETTER O WITH ACUTE
		'\xF4':	'\u00F4',	 // 	LATIN SMALL LETTER O WITH CIRCUMFLEX
		'\xF5':	'\u0151',	 // 	LATIN SMALL LETTER O WITH DOUBLE ACUTE
		'\xF6':	'\u00F6',	 // 	LATIN SMALL LETTER O WITH DIAERESIS
		'\xF7':	'\u00F7',	 // 	DIVISION SIGN
		'\xF8':	'\u0159',	 // 	LATIN SMALL LETTER R WITH CARON
		'\xF9':	'\u016F',	 // 	LATIN SMALL LETTER U WITH RING ABOVE
		'\xFA':	'\u00FA',	 // 	LATIN SMALL LETTER U WITH ACUTE
		'\xFB':	'\u0171',	 // 	LATIN SMALL LETTER U WITH DOUBLE ACUTE
		'\xFC':	'\u00FC',	 // 	LATIN SMALL LETTER U WITH DIAERESIS
		'\xFD':	'\u00FD',	 // 	LATIN SMALL LETTER Y WITH ACUTE
		'\xFE':	'\u0163',	 // 	LATIN SMALL LETTER T WITH CEDILLA
		'\xFF':	'\u02D9',	 // 	DOT ABOVE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	newCodec := &codecMap8Bit{EncodeMap: charmapEncode, DecodeMap: charmapDecode}

	register(newCodec, "ISO-8859-2", "8859-2", "ISO8859-2")

}
