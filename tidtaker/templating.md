# Go html/template – konsis referanse

Prosjektet bruker Gos innebygde [`html/template`](https://pkg.go.dev/html/template@go1.26.3)-pakke (Go 1.26.3, siste stabile per juni 2026). Den er identisk med `text/template`, men med automatisk HTML-escaping for sikkerhet.

## Variabler og data

```html
{{.Felt}}                        <!-- felt på gjeldende objekt (.) -->
{{.Timing.GetString "description"}} <!-- metodekall med argument -->
$var := .Verdi                   <!-- lokal variabel (kun i range/with/if) -->
```

## Betingelser

```html
{{if .Timing.GetBool "isActive"}}
  Aktiv
{{else if eq .Status "paused"}}
  Pauset
{{else}}
  Inaktiv
{{end}}
```

## Løkker (range)

```html
{{range .Items}}
  {{.Navn}}            <!-- . er nå hvert element -->
  {{$.GlobalFelt}}     <!-- $ er alltid rot-objektet -->
{{else}}
  Ingen elementer
{{end}}
```

## Med-blokk (with)

```html
{{with .Timing.GetString "description"}}
  <p>{{.}}</p>         <!-- . er verdien, og blokken hoppes over om verdien er tom -->
{{end}}
```

## Definere og bruke deler (partials)

```html
<!-- Definer -->
{{define "timing_item"}}
  <div>...</div>
{{end}}

<!-- Bruk (fra en annen template) -->
{{template "timing_item" .}}
{{template "timing_item" (dict "Timing" .Timing "Extra" 42)}}
```

## Tilgjengelige egendefinerte funksjoner

| Funksjon | Beskrivelse |
|---|---|
| `formatDuration start stop` | Viser varighet som `1t 23m` |
| `formatTime dt` | Viser klokkeslett som `14:05` |
| `formatDate dt` | Viser dato som `man 2. jan` |
| `localTime dt` | Returnerer datetime-streng for `<input type="datetime-local">` |
| `add a b` | Addisjon |
| `sub a b` | Subtraksjon |
| `seq n` | Lager tallsekvens `[0, 1, …, n-1]` |
| `dict "nøkkel" verdi …` | Lager map for å sende flere verdier til `{{template}}` |
| `contains slice item` | Sjekker om slice inneholder element |
| `eq a b` | Sammenligner to strenger |

## Pipeline og nøsting

```html
{{funcA (funcB .Verdi) .AnnetFelt}}
```

## Kommentarer

```html
{{/* dette er en kommentar */}}
```

## Hvor kommer dataene fra? (PocketBase-eksempel)

Templatene mottar data fra Go-handlere. Et record fra PocketBase sendes inn slik:

```go
// handlers_timer.go
return renderPartial(e, "timing_item", "timing_item", map[string]any{
    "Timing": record,  // *core.Record fra PocketBase
})
```

`record` er et PocketBase-databaseobjekt med metoder for å lese felt:

| Metode | Felttype |
|---|---|
| `GetBool "isActive"` | boolean |
| `GetString "description"` | tekst |
| `GetDateTime "startTime"` | dato/tid |
| `GetStringSlice "tags"` | liste av strenger |

I `timings_list.html` sendes hvert record videre til `timing_item` via `dict`:

```html
{{range .Timings}}
  {{template "timing_item" (dict "Timing" .)}}
{{end}}
```

Og i `timing_item.html` leses feltene via `.Timing.GetBool "isActive"` osv.
