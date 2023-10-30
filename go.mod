module filippo.io/mkcert

go 1.18

require (
	golang.org/x/net v0.0.0-20220421235706-1d1ef9303861
	howett.net/plist v1.0.0
	sm2 v0.0.0-00010101000000-000000000000
	software.sslmate.com/src/go-pkcs12 v0.2.0
)

require (
	github.com/GmSSL/GmSSL-Go v1.3.1 // indirect
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace sm2 => ./sm2
