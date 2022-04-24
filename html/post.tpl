
{{ define "blog"  }}


<h3>
{{ range $i := .Post}}
{{$i.Id}}<br/>
{{$i.Title}}<br/>
{{$i.Body}}<br/>
{{$i.Date}}<br/>
{{end}}
</h3>

<!--

<ul>
    {{range index .post.lst}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
-->

{{ end }}
