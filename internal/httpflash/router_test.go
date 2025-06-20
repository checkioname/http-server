package httpflash

import (
	"flash/internal/core"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockPlugin struct {
	name       string
	matchPaths []string
	called     bool
}

func (m *mockPlugin) Init(config map[string]interface{}) error {
	return nil
}

func (m *mockPlugin) Name() string {
	return m.name
}

func (m *mockPlugin) Match(r *http.Request) bool {
	for _, p := range m.matchPaths {
		if strings.HasPrefix(r.URL.Path, p) {
			return true
		}
	}
	return false
}

func (m *mockPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.called = true
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("called " + m.name))
}

func TestRouter_ServeHTTP_MatchingPlugin(t *testing.T) {
	plugin2 := &mockPlugin{name: "plugin2", matchPaths: []string{"/test2"}}
	plugin1 := &mockPlugin{name: "plugin1", matchPaths: []string{"/test1"}}

	router := NewRouter([]core.PluginHTTP{plugin1, plugin2})

	// Testa caminho que deve casar com plugin1
	req := httptest.NewRequest("GET", "/test1/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", resp.StatusCode)
	}
	if !strings.Contains(body, "plugin1") {
		t.Errorf("Resposta correta esperada para plugin1, got: %s", body)
	}
	if !plugin1.called {
		t.Errorf("plugin1 não foi chamado")
	}
	if plugin2.called {
		t.Errorf("plugin2 não deveria ter sido chamado")
	}
}

func TestRouter_ServeHTTP_NoPluginMatched(t *testing.T) {
	plugin1 := &mockPlugin{name: "plugin1", matchPaths: []string{"/test1"}}
	router := NewRouter([]core.PluginHTTP{plugin1})

	req := httptest.NewRequest("GET", "/unmatchedpath", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Esperava status 404 NotFound para rota não casada, recebeu %d", resp.StatusCode)
	}
}

func TestRouter_ServeHTTP_PluginOrder(t *testing.T) {
	plugin1 := &mockPlugin{name: "plugin1", matchPaths: []string{"/"}}
	plugin2 := &mockPlugin{name: "plugin2", matchPaths: []string{"/test2"}}

	router := NewRouter([]core.PluginHTTP{plugin1, plugin2})

	req := httptest.NewRequest("GET", "/test2/something", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	body := w.Body.String()

	if !plugin1.called {
		t.Errorf("Plugin1 deveria ter sido chamado mas não foi")
	}
	if plugin2.called {
		t.Errorf("Plugin2 não deveria ter sido chamado")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", resp.StatusCode)
	}
	if !strings.Contains(body, "plugin1") {
		t.Errorf("Resposta deveria ser do plugin1, mas é: %s", body)
	}
}
