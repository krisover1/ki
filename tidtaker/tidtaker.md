User: lag et sammendrag av kurset

GitHub Copilot: Read [](file:///workspaces/ki/kurs/01-start.md)

Read [](file:///workspaces/ki/kurs/02-chatting.md)

Read [](file:///workspaces/ki/kurs/03-kontekst.md)

Read [](file:///workspaces/ki/kurs/04-agenter.md)

Read [](file:///workspaces/ki/kurs/05-feils%C3%B8king.md)

Read [](file:///workspaces/ki/kurs/06-utforskning.md)

Read [](file:///workspaces/ki/kurs/07-planer.md)

Read [](file:///workspaces/ki/kurs/08-veien-videre.md)

Her er et sammendrag av kurset:

---

## Sammendrag: KI-kurs med GitHub Copilot

Kurset er et praktisk introduksjonskurs i bruk av KI-agenter i programvareutvikling, med eksempelapplikasjonen **tidtaker** (en enkel timeregistreringsapp i Go) som arbeidsgrunnlag.

### Modul 1 – Start
Sette opp GitHub Codespaces og verifisere at Copilot-chatten fungerer. Første oppgave er å lage et sammendrag av kurset og committe det.

### Modul 2 – Chat
Hvordan bruke chat til å stille spørsmål om kodebasen, gjøre små kontekstavhengige endringer og utforske løsninger. Oppgaven er å bli kjent med tidtaker-kodebasen og teknologiene (PocketBase, htmx, Tailwind).

### Modul 3 – Kontekst
Hvordan gi KI-modellen bedre kontekst: legge ved filer, markere bestemte kodelinjer og hente informasjon fra internett. God kontekst gir bedre svar.

### Modul 4 – Agenter
Agentmodus lar KI-en utføre flere steg automatisk – lage tester, fikse kode, oppdatere dokumentasjon og gjøre sikkerhetsvurderinger i én instruksjon. `AGENTS.md` brukes til å lagre faste instrukser som alltid følger med, slik at man slipper å gjenta seg.

### Modul 5 – Feilsøking
Bruke agenten til å debugge ved å beskrive feilen med naturlig språk eller lime inn feilmeldinger/stack traces.

### Modul 6 – Utforskning
Bruke skills (gjenbrukbare instruksjonsfiler) og agentens brede kunnskap til å utforske løsninger man ikke ville tenkt på selv. Eksempel: lage nettleser-plugin for automatisk timeføring.

### Modul 7 – Planer
Lage skriftlige planer som agenten kan jobbe med selvstendig (AFK). Bruk av Playwright MCP-server og OpenCode CLI for å la agenten styre nettleseren og samle nødvendig informasjon til planen.

### Modul 8 – Veien videre
Optimalisere AFK-resultater: gi agenten mulighet til å verifisere sitt eget arbeid, kjøre uten avbrudd, og bruke AGENTS.md for å styre hva agenten alltid skal gjøre. Temaer som kontekstproblemer, kostnad/ytelse og lokale modeller nevnes som videre emner.

User: I mappen tidtaker ligger det et program. Forklar meg hva det gjør, hvilken teknologi det bruker og hvordan jeg kommer i gang.

GitHub Copilot: Read [](file:///workspaces/ki/tidtaker/main.go)

Read [](file:///workspaces/ki/tidtaker/README.md#1-1), lines 1 to 100

**Tidtaker** er en tidsregistreringsapplikasjon for å holde oversikt over arbeidstimer.

### Hva den gjør
- Start/stopp tidtaking med ett klikk
- Legg til beskrivelse og tagger på registreringene
- Rediger start-/stopptid i etterkant
- Søk og filtrer på tagger og beskrivelse
- Uendelig scroll gjennom historikken

### Teknologi
| Lag | Teknologi |
|-----|-----------|
| Backend | **PocketBase** (Go) – database + API i én binær |
| Frontend | **HTMX** – dynamiske HTML-sider uten å skrive JavaScript |
| Styling | **Tailwind CSS** (via CDN) |
| Testing | **Playwright** (ende-til-ende-tester) |

### Kom i gang

```bash
cd tidtaker
./scripts/dev.sh
```

Åpne [http://localhost:8090](http://localhost:8090), registrer en bruker og begynn å registrere tid.

For å kjøre testene:
```bash
npm install
npx playwright test
```

User: Kan du forklare meg mer om teknologiene pocketbase, htmx og tailwind?

GitHub Copilot: ### PocketBase
Et open source backend-rammeverk skrevet i Go som pakker database, autentisering og API i én enkelt kjørbar fil. Du slipper å sette opp separate tjenester for database (bruker SQLite) og backend – alt kjører som én prosess. Har også et innebygd admin-panel på `/api/collections` og `/_/`. Populært for hobbyprosjekter og intern-verktøy fordi det er enkelt å deploye (bare kopier én binær).

### HTMX
Et JavaScript-bibliotek der du styrer dynamisk oppførsel direkte i HTML-attributter, uten å skrive JavaScript selv. Eksempel fra tidtaker:

```html
<button hx-post="/api/timings/start" hx-target="#timer-controls" hx-swap="outerHTML">
  Start
</button>
```

Her sender HTMX en POST-request når du klikker, og bytter ut `#timer-controls` med HTML-en serveren svarer med. Serveren returnerer HTML-fragmenter i stedet for JSON, så frontenden trenger ingen React/Vue/Svelte.

Passer bra når:
- Serveren din allerede rendrer HTML (Go-templates, som her)
- Du vil ha interaktivitet uten å lære et frontend-rammeverk
- Du vil unngå å vedlikeholde et separat API + SPA

### Tailwind CSS
Et utility-first CSS-rammeverk der du styler ting med ferdiglagde klasser direkte i HTML:

```html
<div class="flex items-center gap-4 rounded-lg bg-white p-4 shadow">
```

I stedet for å skrive CSS-filer skriver du klasser som `text-sm`, `font-bold`, `mt-4`, `hover:bg-blue-500` osv. direkte på elementene. Resultatet er at du sjelden forlater HTML-filen, og du slipper å finne på klassenavn.

---

**Kombinasjonen** er populær for serverrendret web: PocketBase håndterer data og auth, Go-templates lager HTML, HTMX gjør siden interaktiv, og Tailwind styler alt – alt uten et eget frontend-byggesteg.