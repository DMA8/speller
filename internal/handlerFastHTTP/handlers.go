package handlerFastHTTP

import (
	"encoding/json"
	"fmt"
	"speller/internal/storage"
	"strings"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi!")
}

type AddStruct struct {
	SpellName	string `json:"spellName"`
	MisSpells	[]string `json:"misSpells"`

}

func SpellInfo(ctx *fasthttp.RequestCtx) {

}

type Hand struct {
	st *storage.SpellStorage
}

func (h *Hand) Read(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	content, err := h.st.ReadSPell(string(name))
	if err != nil {
		ctx.Error(err.Error(), 404)
	} else {
		fmt.Fprint(ctx, strings.Join(content, " "))
	}
}

func convertToCSV(a AddStruct) string{
	wrongWords := strings.Join(a.MisSpells, "|")
	return fmt.Sprintf("%s;%s;", a.SpellName, wrongWords)
}

func (h *Hand) Add(ctx *fasthttp.RequestCtx) {
	var misSpells AddStruct
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(err.Error(), 404)
	}
	fmt.Println(misSpells)
	err = h.st.AddSpell(convertToCSV(misSpells))
	if err != nil {
		ctx.Error(err.Error(), 500)
	} else {
		fmt.Fprintf(ctx, "%s added", misSpells.SpellName)
	}
}

func (h *Hand) Create(ctx *fasthttp.RequestCtx) {
	//fmt.Println(string(ctx.Request.Body()))
	var misSpells AddStruct
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(err.Error(), 404)
	}
	fmt.Println(misSpells)
	err = h.st.CreateSpell(convertToCSV(misSpells))
	if err != nil {
		ctx.Error(err.Error(), 500)
	} else {
		fmt.Fprintf(ctx, "%s added", misSpells.SpellName)
	}
}

func (h *Hand) FullDelete(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	err := h.st.DeleteSpell(string(name))
	if err != nil {
		fmt.Fprint(ctx, err)
	} else {
		fmt.Fprint(ctx, "ok")
	}
}

func (h *Hand) Delete(ctx *fasthttp.RequestCtx) {
	var misSpells AddStruct
	err := json.Unmarshal(ctx.Request.Body(), &misSpells)
	if err != nil {
		ctx.Error(err.Error(), 404)
	} else {
		err = h.st.DeleteParticularSpellings(convertToCSV(misSpells))
	}
	if err == nil {
		ans, err2 := h.st.ReadSPell(convertToCSV(misSpells))
		if err2 != nil {
			fmt.Fprintf(ctx, "%s",err2)
		} else {
			fmt.Fprint(ctx, ans)
		}
	} else {
		fmt.Fprint(ctx, err)
	}
}

func ConfiguredRouter(a *storage.SpellStorage) *router.Router {
	r := router.New()
	ab := Hand{a}
	//Ro := abc{r, a}
	r.GET("/read", ab.Read)
	r.POST("/add", ab.Add)
	r.POST("/create", ab.Create)
	r.DELETE("/fullDelete", ab.FullDelete)
	r.DELETE("/delete", ab.Delete)
	return r
}
