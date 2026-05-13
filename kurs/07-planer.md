# Planer
Utforskningen fra forrige steg er et fint utgangspunkt for å lage en plan. En plan er et skriftlig selvstendig dokument som beskriver hva som skal lages, og som kan agenten kan jobbe med uten at du trenger å gi input. Her sparer en seg mye tid, for en har beskrevet et større stykke arbeid som agenten kan jobbe uavhengig med.

Hvis du ikke gjorde det i forrige steg, lagre oppsummeringen nå:

>  Lagre oppsummeringen til eksport.md.

## Oppgave: Vurder om det er nok informasjon
Skumles eksport.md. Hvis du selv skulle implementert denne funksjonaliteten, ville du hatt nok informasjon til å lage den 100% uten å bruke andre kilder?

Hvis du ikke holdt på lenge og var nøye med å gi den elementer fra timeføringsprogrammet, tipper jeg den ikke inneholder nok informasjon. Spesielt slike detaljer som hvordan en skal klikke mellom uker og hvilke felt som må fylles ut.

Hvis svaret ditt er "ja", da har du en plan 🙌 Hvis ikke, gjør oppgavene under.

## Oppgave: La agenten selv finne informasjon
TODO: MCP er skrudd av hos DA. Finne en annen vei?

Agenten kan hente nettsider selv, men da har den ikke dine cookies og innlogginger. For å la agenten se timeføringsprogrammet og klikke i det, skal vi bruke en [MCP-server for Playwright](https://playwright.dev/docs/getting-started-mcp).

ADVARSEL: Vi gir agenten tilgang til timeføringene, så informasjonen på nettsiden blir sendt og behandlet i AI-modellen. Dersom du ikke ønsker dette, må du selv gjøre analysen manuelt. I så tilfelle, kan du fortsatt installere MCP-server og prøve deg på noen andre nettsider.

Opprett [mcp.json](../.vscode/mcp.json) med dette innholdet:
```json
{
  "servers": {
    "playwright": {
      "command": "npx",
      "args": ["-y", "@microsoft/mcp-server-playwright"]
    }
  }
}
```


## Oppgave: Gjør dialogen fra utforskningen til en plan

Lagre resultatet i git og push det.

Neste steg er [08-gode-afk-resultat.md](08-gode-afk-resultat.md).
