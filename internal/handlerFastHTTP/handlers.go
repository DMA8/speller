package handlerFastHTTP

import (
	"encoding/json"
	"fmt"
	"log"
	"spellCheck/internal/storage"
	"sync"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type IStorage interface {
	CreateSpell(*storage.Spelling) error
	ReadSpell(string)(*storage.Spelling, error)
	AddSpell(*storage.Spelling) error
	DeleteSpell(string) error
	DeleteParticularSpellings(*storage.Spelling) error
}

type Handler struct {
	st IStorage
	mu sync.Mutex
}

//ConfigureRouter 
func ConfiguredRouter(spellStorage *storage.SpellStorage) *router.Router {
	r := router.New()
	h := Handler{spellStorage, sync.Mutex{}}

	r.GET("/read", h.Read)
	r.POST("/add", h.Add)
	r.POST("/create", h.Create)
	r.DELETE("/fullDelete", h.FullDelete)
	r.DELETE("/delete", h.Delete)
	return r
}

//Just for fun
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi!")
}

//Read: returns all misspels for required word; If there is no such word - returns error
func (h *Handler) Read(ctx *fasthttp.RequestCtx) {
	log.Println("Read query")
	name := ctx.QueryArgs().Peek("name")
	h.mu.Lock()
	content, err := h.st.ReadSpell(string(name))
	h.mu.Unlock()
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
		//ctx.Write(makeJSONErrorResponse(err.Error()))
	} else {
		ctx.Write(makeJSONResultResponse(*content))
	}
}

//Add handler adds accepts json like "{"spellName":"word", "misSpells": ["wordd", "worrd"]}"
// and if "spellName" is in storage, adds misSpells to regarding spellName in storage
func (h *Handler) Add(ctx *fasthttp.RequestCtx) {
	var misSpells storage.Spelling
	log.Println("Add query")
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
	}
	h.mu.Lock()
	err = h.st.AddSpell(&misSpells)
	h.mu.Unlock()
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 500)
	} else {
		ctx.Write(makeJSONAddmit(fmt.Sprintf("added to %s", misSpells.SpellName)))
	}
}

//Create handler creates new pair of spellName and misSpells and adds it in storage
//if spellName was already in storage - returns error
func (h *Handler) Create(ctx *fasthttp.RequestCtx) {
	var misSpells storage.Spelling
	log.Println("Create query")
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
	}
	h.mu.Lock()
	err = h.st.CreateSpell(&misSpells)
	h.mu.Unlock()
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 500)
	} else {
		ctx.Write(makeJSONAddmit(fmt.Sprintf("%s created", misSpells.SpellName)))
	}
}

//FullDelete deletes everything about given spellName
//spellName expected as a parameter in URI (hostname:port/fullDelete?name=spellName)
func (h *Handler) FullDelete(ctx *fasthttp.RequestCtx) {
	log.Println("FullDelete query")
	name := ctx.QueryArgs().Peek("name")
	h.mu.Lock()
	err := h.st.DeleteSpell(string(name))
	h.mu.Unlock()
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
	} else {
		ctx.Write(makeJSONAddmit(fmt.Sprintf("%s deleted", name)))
	}
}

//Delete handler deletes given misSpellings for given spellName
// expects json like  "{"spellName":"word", "misSpells": ["wordd", "worrd"]}" 
//and removes "wordd" and "worrd" from storage for spellName
func (h *Handler) Delete(ctx *fasthttp.RequestCtx) {
	var misSpells storage.Spelling
	log.Println("Delete query")
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
	} else {
		h.mu.Lock()
		err = h.st.DeleteParticularSpellings(&misSpells)
		h.mu.Unlock()
	}
	if err == nil {
		ans, err2 := h.st.ReadSpell(misSpells.SpellName)
		if err2 != nil {
			ctx.Error(string(makeJSONErrorResponse(err.Error())), 404)
		} else {
			ctx.Write(makeJSONResultResponse(*ans))
		}
	} else {
		ctx.Error(string(makeJSONErrorResponse(err.Error())), 500)
	}
}
