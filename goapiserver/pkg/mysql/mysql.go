// Go Api server
// @jeffotoni
// 2019-01-04

package mysql

///
import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once

type MysqlStruct struct {
	Mydb *sql.DB
}

// cache sync.Map
type cache struct {
	mm sync.Map
	sync.Mutex
}

var (
	err  error
	MyDb MysqlStruct
)

var (
	pool = &cache{}
)

// put sync.Map
func (c *cache) put(key, value interface{}) {

	c.Lock()
	defer c.Unlock()
	c.mm.Store(key, value)
}

// get sync.Map
func (c *cache) get(key interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	v, _ := c.mm.Load(key)
	return v

}

// setLoad... fn func() interface{}
func (c *cache) loadStore(key interface{}, fc func() interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	if v, ok := c.mm.Load(key); ok {
		return v
	}

	val := fc()
	c.mm.Store(key, val)
	return val
}

// conectando de forma segura usando goroutine
func Connect() interface{} {

	// ensure uniqueness
	// when using goroutine,
	// avoid running
	once.Do(func() {

		// username:password@tcp(127.0.0.1:3306)/test
		DBINFO := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

		// log.Println(DBINFO)
		// func para ser executada
		// dentro do loadStore
		// quando duas ou mais
		// goroutine chegarem
		// neste mesmo momento
		// de fazer um Store
		fn := func() interface{} {

			sql.Open(DB_SORCE, "")
			MyDb.Mydb, err = sql.Open(DB_SORCE, DBINFO)

			if err != nil {
				errordb := fmt.Sprintf("Unable to connection to database: %v\n", err)
				log.Println("error:: ", errordb)
				defer MyDb.Mydb.Close()
				return nil
			}

			if ok2 := MyDb.Mydb.Ping(); ok2 != nil {

				log.Println("connect error...: ", ok2)
				defer MyDb.Mydb.Close()
				return nil
			}
			//log.Println("connect return sucess:: client [" + DB_NAME + "]")
			return MyDb.Mydb
		}

		// recebendo conexao

		// armazenando cache loadStore
		sqlDb := pool.loadStore(DB_NAME, fn)

		if sqlDb != nil {

			return sqlDb.(*sql.DB)

		} else {

			return nil
		}
	})

	// return only if you have in memory
	if dbMy := pool.get(bancoDb); dbMy != nil {

		// return objeto conexao
		return dbMy.(*sql.DB)
	}
}
