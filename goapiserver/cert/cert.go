// Go Api server
// @jeffotoni
// 2019-01-04
//
package cert

// openssl genrsa -out private.rsa 1024
// openssl rsa -in private.rsa -pubout > public.rsa.pub

const (
	RSA_PRIVATE = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQC6CYlGSwprBFLsYqknJ93fzliMlT6GjqdFmX1CzdD4PLPEHEOd
s/54KilfltXyLMRYrYM4HeF/B3/7rnwPgPbd4ORK8ioMIlaTk9hej8HUi+kpZUhw
vgKHrRHN4r9cYqovoiUw+4PccbHeOTyOTnQT7fu3wDAM1zKCh58Nw+u3LwIDAQAB
AoGAS/ci5+IygdqFJeNvoP6FeiMfQ2CZ5IYRxbjUTlgOsZ7P5Q+JiLb7/QNyW4cT
G82t62wGvf5tmtpsJ1BrdkU6CFNN2Qo9GMez4lERukti1+8HrwWucjo3aa9X60qD
Uala7foabCfOx3p4ZrJ798tIYzmoDbSfviTNSxHYRyqIE2ECQQDr3+E6GNK1GMz2
cbIqKsgaxfx2IgGnZp6qwHgoOvaytau6b9l6uGrviVEpiFIVG5qvp49JsRIa8LER
qiFHrifRAkEAyekUQ4CX527egvW0AqpBH74bycLXbf1hd+2I3XnEfbezMywG1Fie
N+DZ3n+5U+2f0siG5L3hEGzxlSYOMsKu/wJAdJi6sQlMPxD/YGNbetSjDfkIjyzI
PIPRsv5pZxmekUCUnxhjHPLEiZwLbshgKub2VBY0Em1hUcfg/6ZlxRUlwQJARQ+D
7tK9Ilu5n/GKcJ7rR4Au4QPTy7su62ZDuf08SAPdN1OHwnnNJC+0VXY6XYqZb+9G
tFZ99LBOsUUi9hnA4QJAXh85k4kNHODq7tSWW9vKRFONxKNFA4wwxISTJ6ouvqu7
GYIFhFEunzQnRII4ChG0oMFsQFKVhcFBU2aS0zjbkQ==
-----END RSA PRIVATE KEY-----
`

	RSA_PUBLIC = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC6CYlGSwprBFLsYqknJ93fzliM
lT6GjqdFmX1CzdD4PLPEHEOds/54KilfltXyLMRYrYM4HeF/B3/7rnwPgPbd4ORK
8ioMIlaTk9hej8HUi+kpZUhwvgKHrRHN4r9cYqovoiUw+4PccbHeOTyOTnQT7fu3
wDAM1zKCh58Nw+u3LwIDAQAB
-----END PUBLIC KEY-----
`
)
