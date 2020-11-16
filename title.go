type Result struct{ Type string }

func main() {
    types := map[string]string{
        "FindAllString":      "FindAllString",
        "FindString":         "FindString",
        "FindStringSubmatch": "FindStringSubmatch",
    }
    res := &Result{Type: "FindAllString"}

    templateString := `
    <select name="type">
        {{ range $key,$value := .Types }}
            {{ if eq $key $.Res.Type }}
                <option value="{{$key}}" selected>{{$value}}</option>
            {{ else }}
                <option value="{{$key}}">{{$value}}</option>
            {{ end }}
        {{ end }}
    </select>`
    t, err := template.New("index").Parse(templateString)
    if err != nil {
        panic(err)
    }
    var b bytes.Buffer
    writer := bufio.NewWriter(&b)
    err = t.Execute(writer, struct {
        Types map[string]string
        Res   *Result
    }{types, res})
    if err != nil {
        panic(err)
    }
    writer.Flush()
    log.Println(b.String())
}