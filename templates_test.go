package mandrill

import (
	"testing"
)

func TestTemplatesLifecycle(t *testing.T) {
	// delete template from previous tests
	m := NewMandrill(TestAPIKey)
	name := "53102-test-template"
	m.Templates().Delete(name)

	// add template
	tpl := &Template{
		Name:      name,
		FromEmail: TestFromEmail,
		FromName:  "Mr Test",
		Subject:   "Test Subject",
		Code:      `<div mc:edit="name"></div><p mc:edit="address"></p>`,
		Text:      "Test Content",
		Publish:   false,
		Labels:    []string{"test-label"},
	}
	if resp, err := m.Templates().Add(tpl); err != nil {
		t.Error(err)
		return
	} else if resp.Name != name {
		t.Errorf("response template name mistmatch with request: Response: %+v", resp)
		return
	}

	// get info
	if resp, err := m.Templates().Info(name); err != nil {
		t.Error(err)
		return
	} else if resp.Name != name {
		t.Errorf("response template name mistmatch with request: Response: %+v", resp)
		return
	}

	// update
	tpl.Text = "updated"
	if resp, err := m.Templates().Update(tpl); err != nil {
		t.Error(err)
		return
	} else if resp.Text != "updated" {
		t.Errorf("response template text did not update. Response: %+v", resp)
		return
	}

	// publish
	if resp, err := m.Templates().Publish(name); err != nil {
		t.Error(err)
		return
	} else if resp.PublishName != name {
		t.Errorf("response published template name mismatch. Response: %+v", resp)
		return
	}

	// render
	req := &TemplatesRenderRequest{
		TemplateName: name,
		TemplateContent: []TemplateMergeVar{
			{"name", "Greetings *|FNAME|*"},
			{"address", "Mail to *|ADDRESS|*"},
		},
		MergeVars: []TemplateMergeVar{
			{"fname", "Timothy"},
			{"lname", "QA Tester"},
			{"address", "Cul-De-Sac"},
		},
	}
	if resp, err := m.Templates().Render(req); err != nil {
		t.Error(err)
		return
	} else if resp != `<div>Greetings Timothy</div><p>Mail to Cul-De-Sac</p>` {
		exp := `<div>Greetings Timothy</div><p>Mail to Cul-De-Sac</p>`
		t.Errorf("merged content mismatch\nExpected \"%s\"\nReceived: %s", exp, resp)
		return
	}

	// get time-series
	if _, err := m.Templates().TimeSeries(name); err != nil {
		t.Error(err)
		return
	}

	// delete
	if resp, err := m.Templates().Delete(name); err != nil {
		t.Error(err)
		return
	} else if resp.Name != name {
		t.Errorf("response deleted template name mismatch. Response: %+v", resp)
		return
	}
}
