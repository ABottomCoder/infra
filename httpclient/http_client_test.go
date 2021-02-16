package httpclient

import (
	"github.com/ABottomCoder/infra/lb"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/go-eureka-client/eureka"
	"github.com/tietang/props/ini"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttpClient_Do(t *testing.T) {
	//创建一个eureka client
	conf := ini.NewIniFileConfigSource("ec_test.ini")
	client := eureka.NewClient(conf)
	client.Start()
	client.Applications, _ = client.GetApplications()

	//创建一个apps实例
	apps := &lb.Apps{Client: client}

	c := NewHttpClient(apps, &Option{
		Timeout: defaultHttpTimeout,
	})
	Convey("http客户端", t, func() {
		for i := 0; i < 10; i++ {

			req, err := c.NewRequest(http.MethodGet,
				"http://red_envelope/",
				nil, nil)
			So(err, ShouldBeNil)
			So(req, ShouldNotBeNil)
			res, err := c.Do(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			//So(res.StatusCode, ShouldEqual, http.StatusOK)

			defer res.Body.Close()
			d, err := ioutil.ReadAll(res.Body)
			So(err, ShouldBeNil)
			So(d, ShouldNotBeNil)

		}
	})
}
