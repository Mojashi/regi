package regi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
	Body   *bytes.Buffer
	Status int
}

type Log struct {
	ID           int64     `db:"id" json:"id"`
	Path         string    `db:"path" json:"path"`
	Request      string    `db:"request" json:"request"`
	Status       int64     `db:"status" json:"status"`
	Body         string    `db:"body" json:"body"`
	BodyGolden   string    `db:"body_golden" json:"body_golden"`
	StatusGolden string    `db:"status_golden" json:"status_golden"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type RegressionTestConfig struct {
	GoldenDst     string
	CurrentDst    string
	Skipper       middleware.Skipper
	WebUIPort     string
	StaticFilePos string
}

var enabled int64 = 0
var db *sqlx.DB

func RegressionTestWithConfig(config RegressionTestConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = defaultSkipper
	}
	if config.StaticFilePos == "" {
		config.StaticFilePos = "/var/tmp/regifront"
	}

	var err error
	db, err = sqlx.Connect("sqlite3", "/var/tmp/regi.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := setupDB(); err != nil {
		log.Fatal(err)
	}
	serveWebUI(config)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			en := atomic.LoadInt64(&enabled)

			if en == 0 || config.Skipper(c) {
				return next(c)
			}
			req2 := cloneRequest(config.GoldenDst, c.Request())
			resChan := make(chan Response)
			go checker(req2, resChan, config)

			resWatcher := MultiWriteResponse{
				res:  c.Response().Writer,
				buff: new(bytes.Buffer),
			}
			c.Response().Writer = resWatcher
			c.Response().After(func() {
				resChan <- Response{Body: resWatcher.buff, Status: c.Response().Status}
			})
			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}

func defaultSkipper(c echo.Context) bool {
	return c.Request().Method != "GET"
}

func cloneRequest(dst string, req *http.Request) *http.Request {
	var b bytes.Buffer
	b.ReadFrom(req.Body)
	req.Body = ioutil.NopCloser(&b)

	uri, _ := url.ParseRequestURI(req.RequestURI)
	req2, err := http.NewRequest(req.Method, "http://"+dst+uri.Path, ioutil.NopCloser(bytes.NewReader(b.Bytes())))
	if err != nil {
		log.Print(err)
		return nil
	}

	return req2
}

func setupDB() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS difflog (
		id INTEGER primary key autoincrement, 
		path char(30) not null,
		request text not null,
		status int not null,
		body text not null,
		body_golden text not null,
		status_golden int not null,
		created_at TIMESTAMP NOT NULL DEFAULT (DATETIME('now', 'localtime'))
	)`)
	if err != nil {
		return err
	}
	return nil
}

func getResponse(dst string, req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func checker(req2 *http.Request, resChan chan Response, config RegressionTestConfig) bool {
	res2, err := getResponse(config.GoldenDst, req2)
	if err != nil {
		log.Print(err)
	}
	res := <-resChan

	diff := false
	// if res.Status != res2.Status {
	// 	diff = true
	// 	log.Printf("[Different Status] Golden:%s Current:%s", res2.Status, res.Status)
	// }
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return diff
	}
	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Print(err)
		return diff
	}
	// log.Printf("Golden:%s\n--------\nCurrent:%s", string(body), string(body2))
	if !bytes.Equal(body, body2) {
		diff = true
		reqbytes, err := httputil.DumpRequest(req2, true)
		if err != nil {
			log.Print(err)
		}
		_, err = db.Exec("INSERT INTO difflog (path,request,status,body,body_golden,status_golden) values(?,?,?,?,?,?)",
			req2.URL.Path, string(reqbytes), res.Status, body, body2, res2.Status)
		if err != nil {
			log.Print(err)
		}
		// log.Printf("[Different Body] Golden:%s\n--------\nCurrent:%s", string(body), string(body2))
	}
	return diff
}

type MultiWriteResponse struct {
	res  http.ResponseWriter
	buff *bytes.Buffer
}

func (r MultiWriteResponse) Header() http.Header {
	return r.res.Header()
}
func (r MultiWriteResponse) WriteHeader(statusCode int) {
	r.res.WriteHeader(statusCode)
}
func (r MultiWriteResponse) Write(body []byte) (int, error) {
	_, err := r.buff.Write(body)
	if err != nil {
		log.Print(err, string(body))
	}
	return r.res.Write(body)
}

//全然クローンしてないけどまあ
func cloneResponse(res *http.Response, body MultiWriteResponse) *http.Response {
	ret := http.Response{}
	ret = *res
	ret.Body = ioutil.NopCloser(body.buff)
	return &ret
}
