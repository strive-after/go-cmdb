package module

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bndr/gojenkins"
	"net/http"
)

type JenkinsInfo struct {
	user string
	password string
	host string
	token string
}

func NewJenkins() JenkinsInfo{
	return JenkinsInfo{
		user:     beego.AppConfig.String("jenkins_user"),
		password: beego.AppConfig.String("jenkins_password"),
		host:     beego.AppConfig.String("jenkins_host"),
		token:    beego.AppConfig.String("jenkins_token"),
	}
}


var (
	jenkinsclient *gojenkins.Jenkins
	request *gojenkins.Requester
)

func init() {
	jenkins := NewJenkins()
	jenkinsclient ,err = gojenkins.CreateJenkins(nil,jenkins.host,jenkins.user,jenkins.token).Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	request = &gojenkins.Requester{
		Base:      jenkins.host,
		BasicAuth: &gojenkins.BasicAuth{
			Username: jenkins.user,
			Password: jenkins.password,
		},
		Client:    nil,
		CACert:    nil,
		SslVerify: false,
	}
	response ,err := request.GetXML(jenkins.host,&http.Response{}, map[string]string{"job":"test"})
	if err != nil {
		beego.Error(err)
		return
	}
	beego.Info(response)
}


