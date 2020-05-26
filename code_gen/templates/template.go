package templates

import (
	"errors"
	"io/ioutil"

	"github.com/lpxxn/yggdrasill/utils"
)

type ITemplate interface {
	SetPath(path string) error
	GetTemplate() string
}

type template struct {
	template ITemplate
}

func NewTemplate() *template {
	return &template{template: &defaultTemplate{}}
}

func (t *template) SetPath(path string) (err error) {
	if path == "" {
		t.template = &defaultTemplate{}
		return nil
	}
	t.template, err = newCustomerTemplate(path)
	return err
}

func (t *template) GetTemplate() string {
	return t.template.GetTemplate()
}
func (t *template) TemplateHeader() string {
	return templateHeader
}

type defaultTemplate struct {
}

func (t *defaultTemplate) SetPath(path string) error {
	return nil
}

func (t *defaultTemplate) GetTemplate() string {
	return tableModelTemplate
}

type customerTemplate struct {
	body []byte
}

func newCustomerTemplate(path string) (*customerTemplate, error) {
	newTemp := &customerTemplate{}
	if err := newTemp.SetPath(path); err != nil {
		return nil, err
	}
	return newTemp, nil
}

func (t *customerTemplate) SetPath(path string) error {
	exist, err := utils.FileExists(path)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("file not exist")
	}
	t.body, err = ioutil.ReadFile(path)
	return err
}

func (t *customerTemplate) GetTemplate() string {
	return string(t.body)
}
